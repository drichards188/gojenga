package com.hyperion.datalake;


import java.util.List;

import org.springframework.data.mongodb.repository.MongoRepository;

public interface HashRepository extends MongoRepository<Hash, String> {
    Hash findByIteration(Integer iter);
    Long deleteByIteration(Integer iter);
//    List<Tutorial> findByPublished(boolean published);
}
