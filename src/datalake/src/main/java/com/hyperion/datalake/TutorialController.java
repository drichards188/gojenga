package com.hyperion.datalake;

import com.hyperion.datalake.Tutorial;
import com.hyperion.datalake.TutorialRepository;
import org.apache.logging.log4j.LogManager;
import org.json.simple.parser.ParseException;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.*;
import org.springframework.web.bind.annotation.*;
import org.springframework.http.HttpEntity;

import java.io.IOException;
import java.security.NoSuchAlgorithmException;
import java.util.ArrayList;
import java.util.List;
import java.util.Optional;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestMapping;

@CrossOrigin(origins = "http://localhost:8081")
@RestController
@RequestMapping("/lake")
public class TutorialController {
    @Autowired
    TutorialRepository tutorialRepository;

    @Autowired
    HashRepository hashRepository;

    BotnetFuncs botnetFuncs = new BotnetFuncs();
    SqlInter sqlInter = new SqlInter();

    private final Logger logger = LoggerFactory.getLogger(this.getClass());

    @PostMapping("/crypto")
    public ResponseEntity<Tutorial> handlePost(@RequestBody Tutorial tutorial) throws Exception {

        logger.debug("Post mapping triggered");
        sqlInter.main();
//        logger.info("Info: Log4j2 Usage");
//        logger.debug("Debug: Program has finished successfully");
//        logger.error("Error: Program has errors");

        try {
            if (tutorial.getVerb().equals("CRT")) {
                logger.debug("Attempting CRT");
                logger.info("Attempting CRT");

                Tutorial response = botnetFuncs.createAccount(tutorialRepository, hashRepository, tutorial);

//                Tutorial _tutorial = tutorialRepository.save(new Tutorial(tutorial.getAccount(), tutorial.getAmount()));
                return new ResponseEntity<>(response, HttpStatus.CREATED);
            } else if (tutorial.getVerb().equals("TRAN")) {
                logger.debug("Attempting TRAN");
                logger.info("Attempting TRAN");
                Tutorial response = botnetFuncs.transaction(tutorialRepository, hashRepository, tutorial);

                return new ResponseEntity<>(response, HttpStatus.OK);
            } else if (tutorial.getVerb().equals("ADD")) {
//                todo placeholder
            }else if (tutorial.getVerb().equals("QUERY")) {
                logger.debug("Attempting QUERY");
                logger.info("Attempting QUERY");
                Tutorial response = botnetFuncs.findAccount(tutorialRepository, tutorial.getAccount());

                return new ResponseEntity<>(response, HttpStatus.OK);
            } else if (tutorial.getVerb().equals("DLT")) {
                logger.debug("Attempting DLT");
                logger.info("Attempting DLT");
                Tutorial response = botnetFuncs.deleteAccount(tutorialRepository, tutorial.getAccount());

                return new ResponseEntity<>(response, HttpStatus.OK);
            }
            Tutorial respMsg = null;
            respMsg.setMessage("Internal Failure");

            return new ResponseEntity<>(respMsg, HttpStatus.INTERNAL_SERVER_ERROR);

        } catch (Exception e) {
            logger.error("handlePost triggered exception: " + e);
            return new ResponseEntity<>(null, HttpStatus.INTERNAL_SERVER_ERROR);
        }
    }

    @PutMapping("/crypto")
    public ResponseEntity<Tutorial> handlePut(@RequestBody Tutorial tutorial) throws IOException, ParseException, NoSuchAlgorithmException {
        String account = tutorial.getAccount();

        if (tutorial.getVerb().equals("TRAN")) {
            logger.debug("Attempting TRAN");
            logger.info("Attempting TRAN");
            Tutorial response = botnetFuncs.transaction(tutorialRepository, hashRepository, tutorial);

            return new ResponseEntity<>(tutorialRepository.save(response), HttpStatus.OK);

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
            Tutorial respMsg = null;
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
    public ResponseEntity<List<Tutorial>> getAllTutorials(@RequestParam(required = false) String account) {
        try {

            List<Tutorial> tutorials = new ArrayList<Tutorial>();
            Tutorial tutorial = new Tutorial();

            if (account == null)
                tutorialRepository.findAll().forEach(tutorials::add);
            else
                tutorial = botnetFuncs.findAccount(tutorialRepository, account);
            tutorials.add(tutorial);
//                tutorialRepository.findByAccountContaining(account).forEach(tutorials::add);

//            if (tutorials.isEmpty()) {
//                return new ResponseEntity<>(HttpStatus.NO_CONTENT);
//            }

            return new ResponseEntity<>(tutorials, HttpStatus.OK);
        } catch (Exception e) {
            logger.error("Get mapping triggered a error: " + e);
            return new ResponseEntity<>(null, HttpStatus.INTERNAL_SERVER_ERROR);
        }
    }

    @GetMapping("/crypto/{id}")
    public ResponseEntity<Tutorial> getTutorialById(@PathVariable("id") String id) {
        Optional<Tutorial> tutorialData = tutorialRepository.findById(id);

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
            tutorialRepository.deleteById(id);
            return new ResponseEntity<>(HttpStatus.NO_CONTENT);
        } catch (Exception e) {
            return new ResponseEntity<>(HttpStatus.INTERNAL_SERVER_ERROR);
        }
    }

    @DeleteMapping("/crypto")
    public ResponseEntity<HttpStatus> deleteAllTutorials() {
        try {
            tutorialRepository.deleteAll();
            return new ResponseEntity<>(HttpStatus.NO_CONTENT);
        } catch (Exception e) {
            return new ResponseEntity<>(HttpStatus.INTERNAL_SERVER_ERROR);
        }
    }
}