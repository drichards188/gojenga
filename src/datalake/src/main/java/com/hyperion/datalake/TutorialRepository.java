package com.hyperion.datalake;

import java.util.List;

import org.springframework.data.mongodb.repository.MongoRepository;

public interface TutorialRepository extends MongoRepository<Tutorial, String> {
//    List<Tutorial> findByTitleContaining(String account);
//    List<Tutorial> findByPublished(boolean published);
    List<Tutorial> findByAccountContaining(String Account);
    List<Tutorial> findByAccount2Containing(String Account);
    Long deleteByAccount(String Account);
//    List<Tutorial> findByAccount(String Account);
}