package com.hyperion.datalake;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class BlockchainControllerTest {

    @Test
    void testHandlePost() throws Exception {
        BlockchainController blockchainController = new BlockchainController();

        Blockchain blockchain = new Blockchain();
        blockchain.setSourceAccount("sheldon");
        blockchain.setAmount("200");
        blockchain.setVerb("CRT");

        assertEquals(blockchain, blockchainController.handlePost(blockchain), "handlePost should work");
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