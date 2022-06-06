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

public class BotnetFuncs {
//    @Autowired
//    HashRepository hashRepository;
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

//        MongoStruct query = new MongoStruct();
//        query.Account = account;
//        query.Collection = "ledger";
//
//        Document document = new Document();
//        document.append("Account", account);
//        document.append("Amount", defaultAmount);
//
//        String data = account + defaultAmount;
//
//        String hashResult = hashLedger(data);
//
//        String results = MongoInter.InsertOne(query, document);
//        String results = MongoInter.InsertOne(accountVec);
//        if (!results.equals("No Results Found")) {
//            return "Created account: " + account;
//        }

//        return "account not found";
    }

    public Blockchain findAccount(BlockchainRepository blockchainRepository, String Account) {
        Blockchain _blockchain = new Blockchain();
        logger.debug("Attempting findAccount");
        logger.info("Attempting findAccount");
        String amount = "";
        try {
            List<Blockchain> blockchainData = blockchainRepository.findByAccountContaining(Account);

            if (!blockchainData.isEmpty()) {
                _blockchain = blockchainData.get(0);
                Account = _blockchain.getAccount();
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
           Long tutorialData = blockchainRepository.deleteByAccount(Account);

            _blockchain.setMessage("Account Delete Success");
                return _blockchain;

//            if (!tutorialData.isEmpty()) {
//                _tutorial = tutorialData.get(0);
//                Account = _tutorial.getAccount();
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

    public Blockchain transaction(BlockchainRepository blockchainRepository, HashRepository hashRepository, Blockchain blockchain) throws IOException, ParseException, NoSuchAlgorithmException {
        logger.debug("Attempting transaction");
        logger.info("Attempting transaction");

        Blockchain account1 = findAccount(blockchainRepository, blockchain.getAccount());
        Blockchain account2 = findAccount(blockchainRepository, blockchain.getAccount2());
        Blockchain _blockchain = new Blockchain();

//        if (result1.equals("No results Found")) {
//            _tutorial.setMessage("No Results Found");
//            return _tutorial;
//        }
//
//        if (result2.equals("No Results Found")) {
//            _tutorial.setMessage("No Results Found");
//            return _tutorial;
//        }

        Integer amount1 = Integer.parseInt(account1.getAmount()) - Integer.parseInt(blockchain.getAmount());
        Integer amount2 = Integer.parseInt(account2.getAmount()) + Integer.parseInt(blockchain.getAmount());

        account1.setAmount(amount1.toString());
        account2.setAmount(amount2.toString());

        blockchainRepository.deleteByAccount(blockchain.getAccount());
        blockchainRepository.deleteByAccount(blockchain.getAccount2());

        _blockchain = blockchainRepository.save(account1);
        _blockchain = blockchainRepository.save(account2);
        hashLedger(blockchainRepository, hashRepository, blockchain);

        return _blockchain = _blockchain;


//        MongoStruct query = new MongoStruct();
//
//        query.Account = values.Account;
//        query.Amount = amount1.toString();
//
//        String updateMsg = MongoInter.UpdateOne(query);
//
//        if ((updateMsg.equals("Update Unsuccessful"))) {
//            return "Transaction failed";
//        }
//
//        query.clear();
//
//        query.Account = values.Account2;
//        query.Amount = amount2.toString();
//        String data = values.Account + amount1 + values.Account2 + amount2;
//
//        hashLedger(data);
//        updateMsg = MongoInter.UpdateOne(query);
//
//        if ((updateMsg.equals("Update Unsuccessful"))) {
//            return "Transaction failed";
//        }
//
//        return "Transaction Successful";
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
            String ledgerStr = blockchain.getAccount() + blockchain.getAccount2() + blockchain.getAmount();

            hashStruc.setLedger(ledgerStr);
        }

        try {
            Hash saveResults = hashRepository.save(hashStruc);
        } catch (Exception e) {
            logger.error("hashLedger threw an exception");
            System.out.println(e);
        }

        return "Hash Complete";

//        MongoStruct query = new MongoStruct();
//        MongoStruct hashQuery = new MongoStruct();
//        StringBuilder doc = new StringBuilder();
//
//
//        doc.append(results);
//        doc.append(hash);
//        doc.append(timestamp);
//
//        query.Extra = doc.toString();
//        query.Collection = "hashHistory";
//
//        Integer iterator = MongoInter.countDocs(query);
//
//        hashQuery.Field = "Iteration";
//        hashQuery.IntValue = iterator - 1;
//
//        String hashHistory = MongoInter.FindHash(hashQuery);
//
//        Document document = new Document();
//        document.append("Iteration", iterator);
//        document.append("Timestamp", myTime);
//        document.append("Hash", hash);
//        document.append("PreviousHash", hashHistory);
//        document.append("ledger", data);
//
//
//        MongoInter.InsertOne(query, document);
//
////        Long newIterator = iterator - 1;
//
////        query.clear();
////        query.Field = "Version";
////        query.Value = newIterator.toString();
//
//        return "Unsuccessful";
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

//

//
//    public static String deleteAccount(String account) throws NoSuchAlgorithmException, ParseException {
//        MongoStruct query = new MongoStruct();
//        query.Account = account;
//        query.Collection = "ledger";
//
//        Document document = new Document();
//        document.append("Account", account);
//
//        String data = account;
//
//        String hashResult = hashLedger(data);
//
//        String results = MongoInter.DeleteOne(query);
////        String results = MongoInter.InsertOne(accountVec);
//        if (!results.equals("No Results Found")) {
//            return "Created account: " + account;
//        }
//
//        return "account not found";
//    }
//

//

//
//    public static String disco() {
//        return "Unsuccessful";
//    }
//
//    public static String ping(String account) {
//
//        return "Unsuccessful";
//    }
//
//    public static String findAccount(String account) throws ParseException {
//        MongoStruct query = new MongoStruct();
//        query.Account = account;
//        query.Field = "Account";
//        query.Value = account;
//
//        String results = MongoInter.FindOne(query);
////        String results = MongoInter.InsertOne(accountVec);
//        if (!results.equals("No Results Found")) {
//
//            JSONParser parser = new JSONParser();
//            JSONObject json = (JSONObject) parser.parse(results);
//            String userAmount = (String) json.get("Amount");
//            return results;
//        }
//
//        return "account not found";
//    }
//
//    public static Traffic extractValues(String result) throws IOException, ParseException {
//        Traffic values = new Traffic();
//
//        JSONParser parser = new JSONParser();
//        JSONObject jsonObject = (JSONObject) parser.parse(result);
//        System.out.println(jsonObject);
//
//        values.Verb = (String) jsonObject.get("Verb");
//        values.Account = (String) jsonObject.get("Account");
//        values.Account2 = (String) jsonObject.get("Account2");
//        values.Amount = (String) jsonObject.get("Amount");
//        values.Payload = (String) jsonObject.get("Payload");
//
//        return values;
//    }
//
}
