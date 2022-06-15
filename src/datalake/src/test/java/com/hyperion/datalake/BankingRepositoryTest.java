package com.hyperion.datalake;

import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;

import static org.junit.jupiter.api.Assertions.*;

@SpringBootTest
class BankingRepositoryTest {
    @Autowired
    private BankingRepository bankingRepository;

    @Test
    void testRepositorySave() {
        Blockchain blockchain = new Blockchain();
        blockchain.setSourceAccount("test1");
        blockchain.setAmount("200");
        blockchain.setVerb("CRT");

        assertEquals(blockchain, bankingRepository.save(blockchain), "BlockchainRepository should work");
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