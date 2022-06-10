package com.hyperion.datalake;

import org.json.simple.parser.ParseException;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.io.IOException;
import java.security.NoSuchAlgorithmException;
import java.sql.Timestamp;
import java.text.SimpleDateFormat;
import java.util.ArrayList;
import java.util.List;

public class BankingFuncs {
    private final Logger logger = LoggerFactory.getLogger(this.getClass());

    public Blockchain createAccount(BlockchainRepository blockchainRepository, HashRepository hashRepository, Blockchain blockchain) throws NoSuchAlgorithmException {
        Blockchain _blockchain = new Blockchain();
        logger.debug("Attempting createAccount");
        logger.info("Attempting createAccount");
        try {
            _blockchain = blockchainRepository.save(new Blockchain(blockchain.getAccount(), blockchain.getAmount()));
            hashLedger(blockchainRepository, hashRepository, blockchain);
            return _blockchain;
        } catch (Exception e) {
            logger.error("createAccount threw exception");
            _blockchain.setMessage("createAccount failed");
            return _blockchain;
        }
    }

    public Blockchain transaction(BlockchainRepository blockchainRepository, HashRepository hashRepository, Blockchain blockchain) throws IOException, ParseException, NoSuchAlgorithmException {
        logger.debug("Attempting transaction");
        logger.info("Attempting transaction");

        Blockchain sourceAccount = findAccount(blockchainRepository, blockchain.getSourceAccount());
        Blockchain destinationAccount = findAccount(blockchainRepository, blockchain.getDestinationAccount());
        Blockchain _blockchain = new Blockchain();

        Integer amount1 = Integer.parseInt(sourceAccount.getAmount()) - Integer.parseInt(blockchain.getAmount());
        Integer amount2 = Integer.parseInt(destinationAccount.getAmount()) + Integer.parseInt(blockchain.getAmount());

        sourceAccount.setAmount(amount1.toString());
        destinationAccount.setAmount(amount2.toString());

        blockchainRepository.deleteByAccount(blockchain.getSourceAccount());
        blockchainRepository.deleteByAccount(blockchain.getDestinationAccount());

        _blockchain = blockchainRepository.save(sourceAccount);
        _blockchain = blockchainRepository.save(destinationAccount);
        String hashResponse = hashLedger(blockchainRepository, hashRepository, blockchain);
        logger.info("tran hash response is --> " + hashResponse);

        return _blockchain = _blockchain;
    }

    public Blockchain findAccount(BlockchainRepository blockchainRepository, String Account) {
        Blockchain _blockchain = new Blockchain();
        logger.debug("Attempting findAccount");
        logger.info("Attempting findAccount");
        String amount = "";
        try {
            List<Blockchain> blockchainData = blockchainRepository.findByAccount(Account);

            if (!blockchainData.isEmpty()) {
                _blockchain = blockchainData.get(0);
                Account = _blockchain.getSourceAccount();
                amount = _blockchain.getAmount();

                return _blockchain;
            }

            else {
                _blockchain.setMessage("No Results Found");
                return _blockchain;
            }

        } catch (Exception e) {
            logger.error("findAccount threw exception");
            _blockchain.setMessage("createAccount failed");
            return _blockchain;
        }
    }

    public Blockchain deleteAccount(BlockchainRepository blockchainRepository, String Account) {
        Blockchain _blockchain = new Blockchain();
        logger.debug("Attempting deleteAccount");
        logger.info("Attempting deleteAccount");
        try {
           Long tutorialData = blockchainRepository.deleteBySourceAccount(Account);

            _blockchain.setMessage("Account Delete Success");
                return _blockchain;

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
            _blockchain.setMessage("createAccount failed");
            return _blockchain;
        }
    }



    public String hashLedger(BlockchainRepository blockchainRepository, HashRepository hashRepository, Blockchain blockchain) throws NoSuchAlgorithmException, ParseException {
        logger.debug("Attempting hashLedger");
        logger.info("Attempting hashLedger");
        String results = getAllTutorials(blockchainRepository);
        Hash hashStruc = new Hash();
        String prevHash = "";
//        MessageDigest digest = MessageDigest.getInstance("SHA-1");
//        byte[] hash = digest.digest(results.getBytes(StandardCharsets.UTF_8));
        String hash = org.apache.commons.codec.digest.DigestUtils.sha1Hex( results );

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

    public String getAllTutorials(BlockchainRepository blockchainRepository) {
            String result = "";

            List<Blockchain> blockchains = new ArrayList<Blockchain>();
            Blockchain blockchain = new Blockchain();

            blockchainRepository.findAll().forEach(blockchains::add);

            Integer i = 0;
            while (i < blockchains.size()) {
                Blockchain currentTut = blockchains.get(i);
                result = result + currentTut.toHashString();
                i+= 1;
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

    public String deleteAccount(BlockchainRepository blockchainRepository, Blockchain blockchain, String account) throws NoSuchAlgorithmException, ParseException {
        blockchainRepository.deleteBySourceAccount(account);

        return "account deleted";
    }
}
