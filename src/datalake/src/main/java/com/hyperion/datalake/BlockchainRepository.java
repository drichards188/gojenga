package com.hyperion.datalake;

import java.util.List;

import org.springframework.data.mongodb.repository.MongoRepository;

public interface BlockchainRepository extends MongoRepository<Blockchain, String> {
//    List<Tutorial> findByTitleContaining(String account);
//    List<Tutorial> findByPublished(boolean published);
    List<Blockchain> findByAccountContaining(String Account);
    List<Blockchain> findByAccount2Containing(String Account);
    Long deleteByAccount(String Account);
//    List<Tutorial> findByAccount(String Account);
}