package com.hyperion.datalake;

import org.json.simple.parser.ParseException;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.nio.charset.StandardCharsets;
import java.security.MessageDigest;

import java.security.NoSuchAlgorithmException;
import java.sql.Timestamp;
import java.text.SimpleDateFormat;
import java.util.ArrayList;
import java.util.List;

public class BankingFuncs {
    private final Logger logger = LoggerFactory.getLogger(this.getClass());

    enum Crud {
        CREATE,
        READ,
        UPDATE,
        DELETE
    }

    enum Datatypes {
        LEDGER,
        USER,
        HASH,
        OPLOG
    }

    public Traffic createAccount(Traffic traffic, Boolean test) throws NoSuchAlgorithmException {
        logger.debug("Attempting createAccount");
        logger.info("Attempting createAccount");

        if (test.equals(true)) {
            traffic.setRole("TEST");
        } else {
            traffic.setRole("PROD");
        }

        try {
            SqlInter sqlInter = new SqlInter();
            Traffic trafficResponse = sqlInter.sqlHandler(Crud.CREATE, Datatypes.USER, traffic);
            trafficResponse = sqlInter.sqlHandler(Crud.CREATE, Datatypes.LEDGER, traffic);
            Traffic oplogResponse = opLog(traffic, test);
            Traffic hashResponse = hashLedger(traffic, test);

            return trafficResponse;

        } catch (Exception e) {
            logger.error("createAccount threw exception");
            traffic.setMessage("createAccount failed");
            return traffic;
        }
    }

    public Traffic transaction(Traffic traffic, Boolean test) throws NoSuchAlgorithmException {
        logger.debug("Attempting transaction");
        logger.info("Attempting transaction");
        SqlInter sqlInter = new SqlInter();

        Traffic trafficMedium = new Traffic();

        if (test.equals(true)) {
            traffic.setRole("TEST");

            Traffic sourceAccount = sqlInter.sqlHandler(Crud.UPDATE, Datatypes.LEDGER, traffic);

            Traffic hashResponse = hashLedger(traffic, test);
            Traffic oplogResponse = opLog(traffic, test);

        } else {
            trafficMedium.setRole("PROD");

            trafficMedium.setVerb(traffic.getVerb());
            trafficMedium.user.setAccount(traffic.getSourceAccount());
            Traffic sourceAccount = findAccount(trafficMedium, test);

            trafficMedium.user.setAccount(traffic.getDestinationAccount());
            Traffic destinationAccount = findAccount(trafficMedium, test);

            Integer amount1 = null;
            Integer amount2 = null;

            String cleanAmount = sourceAccount.user.amount.split("\\.", 2)[0];
            String cleanAmount2 = destinationAccount.user.amount.split("\\.", 2)[0];

            sourceAccount.user.setAmount(cleanAmount);
            destinationAccount.user.setAmount(cleanAmount2);

            amount1 = Integer.parseInt(sourceAccount.user.getAmount()) - Integer.parseInt(traffic.user.getAmount());
            amount2 = Integer.parseInt(destinationAccount.user.getAmount()) + Integer.parseInt(traffic.user.getAmount());

            sourceAccount.user.setAmount(amount1.toString());
            destinationAccount.user.setAmount(amount2.toString());

            trafficMedium.setVerb(traffic.getVerb());
            trafficMedium.user.setAccount(traffic.getSourceAccount());
            trafficMedium.user.setAmount(amount1.toString());

            sourceAccount = sqlInter.sqlHandler(Crud.UPDATE, Datatypes.LEDGER, trafficMedium);

            trafficMedium.user.setAccount(traffic.getDestinationAccount());
            trafficMedium.user.setAmount(amount2.toString());

            sourceAccount = sqlInter.sqlHandler(Crud.UPDATE, Datatypes.LEDGER, trafficMedium);

            Traffic hashResponse = hashLedger(traffic, test);
            Traffic oplogResponse = opLog(traffic, test);

            logger.info("tran hash response is --> " + hashResponse);
        }

        return traffic;
    }

    public Traffic findAccount(Traffic traffic, Boolean test) {
        logger.debug("Attempting findAccount");
        logger.info("Attempting findAccount");

        SqlInter sqlInter = new SqlInter();
        if (test.equals(true)) {
            traffic.setRole("TEST");
        } else {
            traffic.setRole("PROD");
        }

        try {
            traffic = sqlInter.sqlHandler(Crud.READ, Datatypes.LEDGER, traffic);
            return traffic;

        } catch (Exception e) {
            logger.error("createAccount threw exception");
            traffic.setMessage("findAccount failed");
            return traffic;
        }
    }

    public Traffic deleteAccount(Traffic traffic, Boolean test) {
        logger.debug("Attempting deleteAccount");
        logger.info("Attempting deleteAccount");

        SqlInter sqlInter = new SqlInter();

        if (test.equals(true)) {
            traffic.setRole("TEST");
        } else {
            traffic.setRole("PROD");
        }

        try {
            sqlInter.sqlHandler(Crud.DELETE, Datatypes.LEDGER, traffic);
            sqlInter.sqlHandler(Crud.DELETE, Datatypes.USER, traffic);
            Traffic oplogResponse = opLog(traffic, test);

            traffic.setMessage("Account Delete Success");
            return traffic;

        } catch (Exception e) {
            logger.error("deleteAccount threw an exception");
            traffic.setMessage("createAccount failed");
            return traffic;
        }
    }

    public Traffic hashLedger(Traffic traffic, Boolean test) throws NoSuchAlgorithmException {
        logger.debug("Attempting hashLedger");
        logger.info("Attempting hashLedger");

        SqlInter sqlInter = new SqlInter();
        if (test.equals(true)) {
            traffic.setRole("TEST");

        traffic.setVerb("HASH");


            try {
                Traffic hashResults = sqlInter.sqlHandler(Crud.CREATE, Datatypes.HASH, traffic);
                logger.info("hash saveResults are --> " + hashResults);
            } catch (Exception e) {
                logger.error("hashLedger threw an exception");
                System.out.println(e);
            }
        } else {
            traffic.setRole("PROD");


            traffic.setVerb("HASH");

            traffic = sqlInter.sqlHandler(Crud.READ, Datatypes.HASH, traffic);

            String results = String.valueOf(traffic);

            Integer iteration = null;
            if (traffic.hash.getIteration() != null) {
                iteration = traffic.hash.getIteration();
            } else {
                iteration = 0;
            }

            traffic.hash.setIteration(iteration);

            String prevHash = "";
            if (iteration != 0) {
                prevHash = traffic.hash.getHash();
                traffic.hash.setPreviousHash(prevHash);
            } else {
                prevHash = "00000";
                traffic.hash.setPreviousHash(prevHash);
            }

            traffic.hash.setPreviousHash(prevHash);

            MessageDigest digest = MessageDigest.getInstance("SHA-1");
//        byte[] hash = digest.digest(results.getBytes(StandardCharsets.UTF_8));
            String hash = org.apache.commons.codec.digest.DigestUtils.sha1Hex(results);

            traffic.hash.setHash(hash);

            System.out.println(hash);

            Timestamp timestamp = new Timestamp(System.currentTimeMillis());

            String myTime = new SimpleDateFormat("yyyy-MM-dd HH:mm:ss").format(timestamp);

            traffic.hash.setTimestamp(myTime);


            if (traffic.getVerb().equals("CRT")) {
                String ledgerStr = traffic.user.getAccount() + traffic.user.getAmount();

                traffic.hash.setLedger(ledgerStr);
            } else if (traffic.getVerb().equals("TRAN")) {
                String ledgerStr = traffic.user.getAccount() + traffic.getDestinationAccount() + traffic.user.getAmount();

                traffic.hash.setLedger(ledgerStr);
            } else if (traffic.getVerb().equals("HASH")) {
                String ledgerStr = traffic.user.getAccount() + traffic.user.getAmount();

                traffic.hash.setLedger(ledgerStr);
            }

            try {
                Traffic hashResults = sqlInter.sqlHandler(Crud.CREATE, Datatypes.HASH, traffic);
                logger.info("hash saveResults are --> " + hashResults);
            } catch (Exception e) {
                logger.error("hashLedger threw an exception");
                System.out.println(e);
            }
        }

        return traffic;
    }

    public Traffic opLog(Traffic traffic, Boolean test) {
        logger.debug("Attempting opLog");
        logger.info("Attempting opLog");

        SqlInter sqlInter = new SqlInter();
        if (test.equals(true)) {
            traffic.setRole("TEST");
        } else {
            traffic.setRole("PROD");
        }

        try {
            Traffic trafficResponse = sqlInter.sqlHandler(Crud.CREATE, Datatypes.OPLOG, traffic);
        } catch (Exception e) {
            throw new RuntimeException(e);
        }

        return traffic;
    }

    public String getAllTutorials(BankingRepository bankingRepository) {
        String result = "";

        List<Blockchain> blockchains = new ArrayList<Blockchain>();
        Blockchain blockchain = new Blockchain();

        bankingRepository.findAll().forEach(blockchains::add);

        Integer i = 0;
        while (i < blockchains.size()) {
            Blockchain currentTut = blockchains.get(i);
            result = result + currentTut.toHashString();
            i += 1;
        }

        return result;
    }

    public Integer countDocs(HashRepository hashRepository) {
        List<Hash> tutorials = new ArrayList<Hash>();

        tutorials.addAll(hashRepository.findAll());

        return tutorials.size();
    }

    public String findHash(HashRepository hashRepository, Integer Iteration) {
        Hash wholeHash = hashRepository.findByIteration(Iteration);
        if (wholeHash == null) {
            return "000000";
        }
        return wholeHash.getHash();
    }
}
