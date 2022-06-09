package com.hyperion.datalake;

import java.util.List;

import org.springframework.data.mongodb.repository.MongoRepository;

public interface BlockchainRepository extends MongoRepository<Blockchain, String> {
//    List<Tutorial> findByTitleContaining(String account);
//    List<Tutorial> findByPublished(boolean published);
    List<Blockchain> findBySourceAccountContaining(String Account);
    List<Blockchain> findByDestinationAccountContaining(String Account);
    Long deleteBySourceAccount(String Account);
//    List<Tutorial> findByAccount(String Account);
}