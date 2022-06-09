package com.hyperion.datalake;

import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;

import static org.junit.jupiter.api.Assertions.*;

@SpringBootTest
class BlockchainRepositoryTest {
    @Autowired
    private BlockchainRepository blockchainRepository;

    @Test
    void testRepositorySave() {
        Blockchain blockchain = new Blockchain();
        blockchain.setsourceAccount("test1");
        blockchain.setAmount("200");
        blockchain.setVerb("CRT");

        assertEquals(blockchain, blockchainRepository.save(blockchain), "BlockchainRepository should work");
    }

    @Test
    void findByAccountContaining() {
    }

    @Test
    void findByAccount2Containing() {
    }

    @Test
    void deleteByAccount() {
    }
}