package com.hyperion.datalake;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.ArrayList;
import java.util.List;
import java.util.Optional;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.web.bind.annotation.RequestMapping;

@CrossOrigin(origins = "http://localhost:8081")
@RestController
@RequestMapping("/lake")
public class BankingController {
    @Autowired
    BankingRepository bankingRepository;
    @Autowired
    HashRepository hashRepository;
    BankingFuncs bankingFuncs = new BankingFuncs();

    private final Logger logger = LoggerFactory.getLogger(this.getClass());

    @PostMapping("/crypto")
    public ResponseEntity<Traffic> handlePost(@RequestBody Traffic traffic) throws Exception {
        logger.debug("Post mapping triggered");

        try {
            switch (traffic.getVerb()) {
                case "CRT": {
                    logger.debug("Attempting CRT");
                    logger.info("Attempting CRT");

                    //createAccount also calls the hashing functions
                    Traffic response = bankingFuncs.createAccount(traffic);
                    return new ResponseEntity<>(response, HttpStatus.CREATED);
                }
                case "ADD":
//                todo placeholder
                    break;
                case "QUERY": {
                    logger.debug("Attempting QUERY");
                    logger.info("Attempting QUERY");
                    
                    Traffic response = bankingFuncs.findAccount(traffic);
                    return new ResponseEntity<>(response, HttpStatus.OK);
                }
                case "DLT": {
                    logger.debug("Attempting DLT");
                    logger.info("Attempting DLT");
                    
//                    Traffic response = bankingFuncs.deleteAccount(blockchainRepository, traffic.getAccount());

                    Traffic response = new Traffic();
                    return new ResponseEntity<>(response, HttpStatus.OK);
                }
            }
            Traffic respMsg = null;
            respMsg.setMessage("Internal Failure");

            return new ResponseEntity<>(respMsg, HttpStatus.INTERNAL_SERVER_ERROR);

        } catch (Exception e) {
            logger.error("handlePost triggered exception: " + e);
            return new ResponseEntity<>(null, HttpStatus.INTERNAL_SERVER_ERROR);
        }
    }

    @PutMapping("/crypto")
    public ResponseEntity<Traffic> handlePut(@RequestBody Traffic traffic) throws Exception {
        String sourceAccount = traffic.getSourceAccount();

        if (traffic.getVerb().equals("TRAN")) {
            logger.debug("Attempting TRAN");
            logger.info("Attempting TRAN");
            Traffic response = bankingFuncs.transaction(traffic);

            return new ResponseEntity<>(response, HttpStatus.OK);
        } else {
            Traffic respMsg = null;
            respMsg.setMessage("Internal Failure");
            logger.error("TRAN caused error");
            return new ResponseEntity<>(respMsg, HttpStatus.INTERNAL_SERVER_ERROR);
        }
    }

    @GetMapping("/crypto")
    public ResponseEntity<List<Traffic>> getAllBlockchains(@RequestParam(required = false) String account) {
        try {

            List<Traffic> blockchains = new ArrayList<Traffic>();
            Traffic traffic = new Traffic();

//            if (account == null)
//                blockchainRepository.findAll().forEach(blockchains::add);
//            else
//                blockchain = bankingFuncs.findAccount(blockchainRepository, account);
//            blockchains.add(blockchain);
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
        Optional<Blockchain> BlockchainData = bankingRepository.findById(id);

        if (BlockchainData.isPresent()) {
            return new ResponseEntity<>(BlockchainData.get(), HttpStatus.OK);
        } else {
            return new ResponseEntity<>(HttpStatus.NOT_FOUND);
        }
    }

    @DeleteMapping("/crypto/{id}")
    public ResponseEntity<HttpStatus> deleteBlockchain(@PathVariable("id") String id) {
        try {
            bankingRepository.deleteById(id);
            return new ResponseEntity<>(HttpStatus.NO_CONTENT);
        } catch (Exception e) {
            return new ResponseEntity<>(HttpStatus.INTERNAL_SERVER_ERROR);
        }
    }

    @DeleteMapping("/crypto")
    public ResponseEntity<HttpStatus> deleteAllBlockchains() {
        try {
            bankingRepository.deleteAll();
            return new ResponseEntity<>(HttpStatus.NO_CONTENT);
        } catch (Exception e) {
            return new ResponseEntity<>(HttpStatus.INTERNAL_SERVER_ERROR);
        }
    }
}