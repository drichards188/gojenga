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
        sqlInter.main();
//        logger.info("Info: Log4j2 Usage");
//        logger.debug("Debug: Program has finished successfully");
//        logger.error("Error: Program has errors");

        try {
            if (blockchain.getVerb().equals("CRT")) {
                logger.debug("Attempting CRT");
                logger.info("Attempting CRT");

                Blockchain response = botnetFuncs.createAccount(blockchainRepository, hashRepository, blockchain);

//                Tutorial _tutorial = tutorialRepository.save(new Tutorial(tutorial.getAccount(), tutorial.getAmount()));
                return new ResponseEntity<>(response, HttpStatus.CREATED);
            } else if (blockchain.getVerb().equals("TRAN")) {
                logger.debug("Attempting TRAN");
                logger.info("Attempting TRAN");
                Blockchain response = botnetFuncs.transaction(blockchainRepository, hashRepository, blockchain);

                return new ResponseEntity<>(response, HttpStatus.OK);
            } else if (blockchain.getVerb().equals("ADD")) {
//                todo placeholder
            }else if (blockchain.getVerb().equals("QUERY")) {
                logger.debug("Attempting QUERY");
                logger.info("Attempting QUERY");
                Blockchain response = botnetFuncs.findAccount(blockchainRepository, blockchain.getAccount());

                return new ResponseEntity<>(response, HttpStatus.OK);
            } else if (blockchain.getVerb().equals("DLT")) {
                logger.debug("Attempting DLT");
                logger.info("Attempting DLT");
                Blockchain response = botnetFuncs.deleteAccount(blockchainRepository, blockchain.getAccount());

                return new ResponseEntity<>(response, HttpStatus.OK);
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

//            List<Tutorial> tutorialData = tutorialRepository.findByAccountContaining(account);
//
//            if (!tutorialData.isEmpty()) {
//                Tutorial _tutorial = tutorialData.get(0);
//                _tutorial.setAmount(tutorial.getAmount());
//                _tutorial.setAccount(tutorial.getAccount());
//
//                return new ResponseEntity<>(tutorialRepository.save(_tutorial), HttpStatus.OK);
//            }

//            if (tutorialData.isPresent()) {
//                Tutorial _tutorial = tutorialData.get();
//                _tutorial.setAmount(tutorial.getAmount());
////                _tutorial.setTitle(tutorial.getTitle());
////                _tutorial.setDescription(tutorial.getDescription());
////                _tutorial.setPublished(tutorial.isPublished());
//                return new ResponseEntity<>(tutorialRepository.save(_tutorial), HttpStatus.OK);
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
//            if (tutorial.getVerb().equals("TRAN")) {
//                Tutorial _tutorial = tutorialRepository.save(new Tutorial(tutorial.getAccount(), tutorial.getAmount()));
//                return new ResponseEntity<>(_tutorial, HttpStatus.CREATED);
//            }
//            Tutorial respMsg = null;
//            respMsg.setMessage("Internal Failure");
//
//            return new ResponseEntity<>(respMsg, HttpStatus.INTERNAL_SERVER_ERROR);
//
//        } catch (Exception e) {
//            return new ResponseEntity<>(null, HttpStatus.INTERNAL_SERVER_ERROR);
//        }
//        Tutorial respMsg = null;
//        respMsg.setMessage("Internal Failure");

//        return new ResponseEntity<>(respMsg, HttpStatus.INTERNAL_SERVER_ERROR);
    }

    @GetMapping("/crypto")
    public ResponseEntity<List<Blockchain>> getAllTutorials(@RequestParam(required = false) String account) {
        try {

            List<Blockchain> blockchains = new ArrayList<Blockchain>();
            Blockchain blockchain = new Blockchain();

            if (account == null)
                blockchainRepository.findAll().forEach(blockchains::add);
            else
                blockchain = botnetFuncs.findAccount(blockchainRepository, account);
            blockchains.add(blockchain);
//                tutorialRepository.findByAccountContaining(account).forEach(tutorials::add);

//            if (tutorials.isEmpty()) {
//                return new ResponseEntity<>(HttpStatus.NO_CONTENT);
//            }

            return new ResponseEntity<>(blockchains, HttpStatus.OK);
        } catch (Exception e) {
            logger.error("Get mapping triggered a error: " + e);
            return new ResponseEntity<>(null, HttpStatus.INTERNAL_SERVER_ERROR);
        }
    }

    @GetMapping("/crypto/{id}")
    public ResponseEntity<Blockchain> getTutorialById(@PathVariable("id") String id) {
        Optional<Blockchain> tutorialData = blockchainRepository.findById(id);

        if (tutorialData.isPresent()) {
            return new ResponseEntity<>(tutorialData.get(), HttpStatus.OK);
        } else {
            return new ResponseEntity<>(HttpStatus.NOT_FOUND);
        }
    }

//    @GetMapping("/crypto/published")
//    public ResponseEntity<List<Tutorial>> findByPublished() {
//        try {
//            List<Tutorial> tutorials = tutorialRepository.findByPublished(true);
//
//            if (tutorials.isEmpty()) {
//                return new ResponseEntity<>(HttpStatus.NO_CONTENT);
//            }
//            return new ResponseEntity<>(tutorials, HttpStatus.OK);
//        } catch (Exception e) {
//            return new ResponseEntity<>(HttpStatus.INTERNAL_SERVER_ERROR);
//        }
//    }

//    @PutMapping("/crypto/{id}")
//    public ResponseEntity<Tutorial> updateTutorial(@PathVariable("id") String id, @RequestBody Tutorial tutorial) {
//        Optional<Tutorial> tutorialData = tutorialRepository.findById(id);
//
//        if (tutorialData.isPresent()) {
//            Tutorial _tutorial = tutorialData.get();
//            _tutorial.setTitle(tutorial.getTitle());
//            _tutorial.setDescription(tutorial.getDescription());
//            _tutorial.setPublished(tutorial.isPublished());
//            return new ResponseEntity<>(tutorialRepository.save(_tutorial), HttpStatus.OK);
//        } else {
//            return new ResponseEntity<>(HttpStatus.NOT_FOUND);
//        }
//    }

    @DeleteMapping("/crypto/{id}")
    public ResponseEntity<HttpStatus> deleteTutorial(@PathVariable("id") String id) {
        try {
            blockchainRepository.deleteById(id);
            return new ResponseEntity<>(HttpStatus.NO_CONTENT);
        } catch (Exception e) {
            return new ResponseEntity<>(HttpStatus.INTERNAL_SERVER_ERROR);
        }
    }

    @DeleteMapping("/crypto")
    public ResponseEntity<HttpStatus> deleteAllTutorials() {
        try {
            blockchainRepository.deleteAll();
            return new ResponseEntity<>(HttpStatus.NO_CONTENT);
        } catch (Exception e) {
            return new ResponseEntity<>(HttpStatus.INTERNAL_SERVER_ERROR);
        }
    }
}