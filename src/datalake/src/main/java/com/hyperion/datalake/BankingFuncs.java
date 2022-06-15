package com.hyperion.datalake;

import org.json.simple.parser.ParseException;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.security.NoSuchAlgorithmException;
import java.sql.Timestamp;
import java.text.SimpleDateFormat;
import java.util.ArrayList;
import java.util.List;

public class BankingFuncs {
    private final Logger logger = LoggerFactory.getLogger(this.getClass());

    public Traffic createAccount(Traffic traffic) throws NoSuchAlgorithmException {
        logger.debug("Attempting createAccount");
        logger.info("Attempting createAccount");
        try {
            SqlInter sqlInter = new SqlInter();
            traffic = sqlInter.sqlHandler("CRT", traffic);

            return traffic;
        } catch (Exception e) {
            logger.error("createAccount threw exception");
            traffic.setMessage("createAccount failed");
            return traffic;
        }
    }

    public Traffic transaction(Traffic traffic) throws Exception {
        logger.debug("Attempting transaction");
        logger.info("Attempting transaction");
        SqlInter sqlInter = new SqlInter();

        Traffic trafficMedium = new Traffic();
        trafficMedium.setVerb(traffic.getVerb());
        trafficMedium.user.setAccount(traffic.getSourceAccount());
        Traffic sourceAccount = findAccount(trafficMedium);

        trafficMedium.user.setAccount(traffic.getDestinationAccount());
        Traffic destinationAccount = findAccount(trafficMedium);

        Integer amount1 = null;
        Integer amount2 = null;

        String cleanAmount = sourceAccount.user.amount.split("\\.", 2)[0];
        String cleanAmount2 = destinationAccount.user.amount.split("\\.", 2)[0];

        sourceAccount.user.setAmount(cleanAmount);
        destinationAccount.user.setAmount(cleanAmount2);

        try {
            amount1 = Integer.parseInt(sourceAccount.user.getAmount()) - Integer.parseInt(traffic.user.getAmount());
            amount2 = Integer.parseInt(destinationAccount.user.getAmount()) + Integer.parseInt(traffic.user.getAmount());
        } catch (Exception e) {
            logger.debug(String.valueOf(e));
        }

        sourceAccount.user.setAmount(amount1.toString());
        destinationAccount.user.setAmount(amount2.toString());

        trafficMedium.setVerb(traffic.getVerb());
        trafficMedium.user.setAccount(traffic.getSourceAccount());
        trafficMedium.user.setAmount(amount1.toString());

        sourceAccount = sqlInter.sqlHandler("UPDATE", trafficMedium);

        trafficMedium.user.setAccount(traffic.getDestinationAccount());
        trafficMedium.user.setAmount(amount2.toString());

        sourceAccount = sqlInter.sqlHandler("UPDATE", trafficMedium);

        String hashResponse = null;
        logger.info("tran hash response is --> " + hashResponse);

        return traffic;
    }

    public Traffic findAccount(Traffic traffic) {
        logger.debug("Attempting findAccount");
        logger.info("Attempting findAccount");
        try {
            SqlInter sqlInter = new SqlInter();
            traffic = sqlInter.sqlHandler("QUERY", traffic);
            return traffic;

        } catch (Exception e) {
            logger.error("createAccount threw exception");
            traffic.setMessage("findAccount failed");
            return traffic;
        }
    }

    public Blockchain deleteAccount(BankingRepository bankingRepository, String Account) {
        Blockchain traffic = new Blockchain();
        logger.debug("Attempting deleteAccount");
        logger.info("Attempting deleteAccount");
        try {
            Long tutorialData = bankingRepository.deleteBySourceAccount(Account);

            traffic.setMessage("Account Delete Success");
            return traffic;

//            if (!tutorialData.isEmpty()) {
//                _tutorial = tutorialData.get(0);
//                Account = _tutorial.getSourceAccount();
//                amount = _tutorial.getAmount();
//
//                return _tutorial;
//            }
//
//            else {
//                _tutorial.setMessage("No Results Found");
//                return _tutorial;
//            }

        } catch (Exception e) {
            logger.error("deleteAccount threw an exception");
            traffic.setMessage("createAccount failed");
            return traffic;
        }
    }


    public String hashLedger(BankingRepository bankingRepository, HashRepository hashRepository, Blockchain blockchain) throws NoSuchAlgorithmException, ParseException {
        logger.debug("Attempting hashLedger");
        logger.info("Attempting hashLedger");
        String results = getAllTutorials(bankingRepository);
        Hash hashStruc = new Hash();
        String prevHash = "";
//        MessageDigest digest = MessageDigest.getInstance("SHA-1");
//        byte[] hash = digest.digest(results.getBytes(StandardCharsets.UTF_8));
        String hash = org.apache.commons.codec.digest.DigestUtils.sha1Hex(results);

        hashStruc.setHash(hash);

        System.out.println(hash);

        Timestamp timestamp = new Timestamp(System.currentTimeMillis());

        String myTime = new SimpleDateFormat("yyyy-MM-dd HH:mm:ss").format(timestamp);

        hashStruc.setTimestamp(myTime);

        Integer iteration = countDocs(hashRepository);

        hashStruc.setIteration(iteration);

        if (iteration != 0) {
            prevHash = findHash(hashRepository, iteration - 1);
        } else {
            prevHash = "00000";
        }

        hashStruc.setPreviousHash(prevHash);

        if (blockchain.getVerb().equals("CRT")) {
            String ledgerStr = blockchain.getAccount() + blockchain.getAmount();

            hashStruc.setLedger(ledgerStr);
        } else if (blockchain.getVerb().equals("TRAN")) {
            String ledgerStr = blockchain.getAccount() + blockchain.getDestinationAccount() + blockchain.getAmount();

            hashStruc.setLedger(ledgerStr);
        }

        try {
            Hash saveResults = hashRepository.save(hashStruc);
            logger.info("hash saveResults are --> " + saveResults);
        } catch (Exception e) {
            logger.error("hashLedger threw an exception");
            System.out.println(e);
        }

        return "Hash Complete";
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

    public String deleteAccount(BankingRepository bankingRepository, Blockchain blockchain, String account) throws NoSuchAlgorithmException, ParseException {
        bankingRepository.deleteBySourceAccount(account);

        return "account deleted";
    }
}
