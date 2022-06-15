package com.hyperion.datalake;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class BankingControllerTest {

    @Test
    void testHandlePost() throws Exception {
        BankingController bankingController = new BankingController();

        Traffic blockchain = new Traffic();
        blockchain.user.setAccount("sheldon");
        blockchain.user.setAmount("200");
        blockchain.setVerb("CRT");

        assertEquals(blockchain, bankingController.handlePost(blockchain), "handlePost should work");
    }

    @Test
    void handlePut() {
    }

    @Test
    void getAllBlockchains() {
    }

    @Test
    void getBlockchainById() {
    }

    @Test
    void deleteBlockchain() {
    }

    @Test
    void deleteAllBlockchains() {
    }
}