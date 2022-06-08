package com.hyperion.datalake;

import org.json.simple.parser.ParseException;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.io.IOException;
import java.security.NoSuchAlgorithmException;
import java.util.ArrayList;
import java.util.List;
import java.util.Optional;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.web.bind.annotation.RequestMapping;

@CrossOrigin(origins = "http://localhost:8081")
@RestController
@RequestMapping("/lake")
public class BlockchainController {
    @Autowired
    BlockchainRepository blockchainRepository;

    @Autowired
    HashRepository hashRepository;

    BotnetFuncs botnetFuncs = new BotnetFuncs();
    SqlInter sqlInter = new SqlInter();

    private final Logger logger = LoggerFactory.getLogger(this.getClass());

    @PostMapping("/crypto")
    public ResponseEntity<Blockchain> handlePost(@RequestBody Blockchain blockchain) throws Exception {

        logger.debug("Post mapping triggered");
//        sqlInter.main();

        try {
            switch (blockchain.getVerb()) {
                case "CRT": {
                    logger.debug("Attempting CRT");
                    logger.info("Attempting CRT");

//                    Blockchain response = botnetFuncs.createAccount(blockchainRepository, hashRepository, blockchain);

                Blockchain response = blockchainRepository.save(blockchain);
                    return new ResponseEntity<>(response, HttpStatus.CREATED);
                }
                case "TRAN": {
                    logger.debug("Attempting TRAN");
                    logger.info("Attempting TRAN");
                    Blockchain response = botnetFuncs.transaction(blockchainRepository, hashRepository, blockchain);

                    return new ResponseEntity<>(response, HttpStatus.OK);
                }
                case "ADD":
//                todo placeholder
                    break;
                case "QUERY": {
                    logger.debug("Attempting QUERY");
                    logger.info("Attempting QUERY");
                    Blockchain response = botnetFuncs.findAccount(blockchainRepository, blockchain.getAccount());

                    return new ResponseEntity<>(response, HttpStatus.OK);
                }
                case "DLT": {
                    logger.debug("Attempting DLT");
                    logger.info("Attempting DLT");
                    Blockchain response = botnetFuncs.deleteAccount(blockchainRepository, blockchain.getAccount());

                    return new ResponseEntity<>(response, HttpStatus.OK);
                }
            }
            Blockchain respMsg = null;
            respMsg.setMessage("Internal Failure");

            return new ResponseEntity<>(respMsg, HttpStatus.INTERNAL_SERVER_ERROR);

        } catch (Exception e) {
            logger.error("handlePost triggered exception: " + e);
            return new ResponseEntity<>(null, HttpStatus.INTERNAL_SERVER_ERROR);
        }
    }

    @PutMapping("/crypto")
    public ResponseEntity<Blockchain> handlePut(@RequestBody Blockchain blockchain) throws IOException, ParseException, NoSuchAlgorithmException {
        String account = blockchain.getAccount();

        if (blockchain.getVerb().equals("TRAN")) {
            logger.debug("Attempting TRAN");
            logger.info("Attempting TRAN");
            Blockchain response = botnetFuncs.transaction(blockchainRepository, hashRepository, blockchain);

            return new ResponseEntity<>(blockchainRepository.save(response), HttpStatus.OK);

//            List<Blockchain> BlockchainData = blockchainRepository.findByAccountContaining(account);
//
//            if (!BlockchainData.isEmpty()) {
//                Blockchain _Blockchain = BlockchainData.get(0);
//                _Blockchain.setAmount(Blockchain.getAmount());
//                _Blockchain.setAccount(Blockchain.getAccount());
//
//                return new ResponseEntity<>(blockchainRepository.save(_Blockchain), HttpStatus.OK);
//            }

//            if (BlockchainData.isPresent()) {
//                Blockchain _Blockchain = BlockchainData.get();
//                _Blockchain.setAmount(Blockchain.getAmount());
////                _Blockchain.setTitle(Blockchain.getTitle());
////                _Blockchain.setDescription(Blockchain.getDescription());
////                _Blockchain.setPublished(Blockchain.isPublished());
//                return new ResponseEntity<>(blockchainRepository.save(_Blockchain), HttpStatus.OK);
//            } else {
//                return new ResponseEntity<>(HttpStatus.NOT_FOUND);
//            }
        } else {
            Blockchain respMsg = null;
            respMsg.setMessage("Internal Failure");
            logger.error("TRAN caused error");
            return new ResponseEntity<>(respMsg, HttpStatus.INTERNAL_SERVER_ERROR);
        }

//        try {
//
//            if (Blockchain.getVerb().equals("TRAN")) {
//                Blockchain _Blockchain = blockchainRepository.save(new Blockchain(Blockchain.getAccount(), Blockchain.getAmount()));
//                return new ResponseEntity<>(_Blockchain, HttpStatus.CREATED);
//            }
//            Blockchain respMsg = null;
//            respMsg.setMessage("Internal Failure");
//
//            return new ResponseEntity<>(respMsg, HttpStatus.INTERNAL_SERVER_ERROR);
//
//        } catch (Exception e) {
//            return new ResponseEntity<>(null, HttpStatus.INTERNAL_SERVER_ERROR);
//        }
//        Blockchain respMsg = null;
//        respMsg.setMessage("Internal Failure");

//        return new ResponseEntity<>(respMsg, HttpStatus.INTERNAL_SERVER_ERROR);
    }

    @GetMapping("/crypto")
    public ResponseEntity<List<Blockchain>> getAllBlockchains(@RequestParam(required = false) String account) {
        try {

            List<Blockchain> blockchains = new ArrayList<Blockchain>();
            Blockchain blockchain = new Blockchain();

            if (account == null)
                blockchainRepository.findAll().forEach(blockchains::add);
            else
                blockchain = botnetFuncs.findAccount(blockchainRepository, account);
            blockchains.add(blockchain);
//                blockchainRepository.findByAccountContaining(account).forEach(Blockchains::add);

//            if (Blockchains.isEmpty()) {
//                return new ResponseEntity<>(HttpStatus.NO_CONTENT);
//            }

            return new ResponseEntity<>(blockchains, HttpStatus.OK);
        } catch (Exception e) {
            logger.error("Get mapping triggered a error: " + e);
            return new ResponseEntity<>(null, HttpStatus.INTERNAL_SERVER_ERROR);
        }
    }

    @GetMapping("/crypto/{id}")
    public ResponseEntity<Blockchain> getBlockchainById(@PathVariable("id") String id) {
        Optional<Blockchain> BlockchainData = blockchainRepository.findById(id);

        if (BlockchainData.isPresent()) {
            return new ResponseEntity<>(BlockchainData.get(), HttpStatus.OK);
        } else {
            return new ResponseEntity<>(HttpStatus.NOT_FOUND);
        }
    }

//    @GetMapping("/crypto/published")
//    public ResponseEntity<List<Blockchain>> findByPublished() {
//        try {
//            List<Blockchain> Blockchains = blockchainRepository.findByPublished(true);
//
//            if (Blockchains.isEmpty()) {
//                return new ResponseEntity<>(HttpStatus.NO_CONTENT);
//            }
//            return new ResponseEntity<>(Blockchains, HttpStatus.OK);
//        } catch (Exception e) {
//            return new ResponseEntity<>(HttpStatus.INTERNAL_SERVER_ERROR);
//        }
//    }

//    @PutMapping("/crypto/{id}")
//    public ResponseEntity<Blockchain> updateBlockchain(@PathVariable("id") String id, @RequestBody Blockchain Blockchain) {
//        Optional<Blockchain> BlockchainData = blockchainRepository.findById(id);
//
//        if (BlockchainData.isPresent()) {
//            Blockchain _Blockchain = BlockchainData.get();
//            _Blockchain.setTitle(Blockchain.getTitle());
//            _Blockchain.setDescription(Blockchain.getDescription());
//            _Blockchain.setPublished(Blockchain.isPublished());
//            return new ResponseEntity<>(blockchainRepository.save(_Blockchain), HttpStatus.OK);
//        } else {
//            return new ResponseEntity<>(HttpStatus.NOT_FOUND);
//        }
//    }

    @DeleteMapping("/crypto/{id}")
    public ResponseEntity<HttpStatus> deleteBlockchain(@PathVariable("id") String id) {
        try {
            blockchainRepository.deleteById(id);
            return new ResponseEntity<>(HttpStatus.NO_CONTENT);
        } catch (Exception e) {
            return new ResponseEntity<>(HttpStatus.INTERNAL_SERVER_ERROR);
        }
    }

    @DeleteMapping("/crypto")
    public ResponseEntity<HttpStatus> deleteAllBlockchains() {
        try {
            blockchainRepository.deleteAll();
            return new ResponseEntity<>(HttpStatus.NO_CONTENT);
        } catch (Exception e) {
            return new ResponseEntity<>(HttpStatus.INTERNAL_SERVER_ERROR);
        }
    }
}