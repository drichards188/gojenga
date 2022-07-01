package com.hyperion.datalake;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;

import java.security.NoSuchAlgorithmException;

import static org.junit.jupiter.api.Assertions.*;

class BankingFuncsTest {
    SqlInter sqlInter;
    BankingFuncs bankingFuncs;

    @BeforeEach
    void setUp() {
        sqlInter = new SqlInter();
        bankingFuncs = new BankingFuncs();
    }

    @Test
    void createAccount() throws NoSuchAlgorithmException {
        Traffic traffic = new Traffic();
        traffic.setVerb("CRT");
        traffic.user.setAccount("david");
        traffic.user.setAmount("200");
        traffic.user.setPassword("mypassword");

        bankingFuncs.createAccount(traffic, true);
    }

    @Test
    void transaction() throws NoSuchAlgorithmException {
        Traffic traffic = new Traffic();
        traffic.setVerb("TRAN");
        traffic.setSourceAccount("david");
        traffic.setDestinationAccount("allie");
        traffic.user.setAmount("10");

        bankingFuncs.transaction(traffic, true);
    }

    @Test
    void findAccount() {
        Traffic traffic = new Traffic();
        traffic.setVerb("QUERY");
        traffic.user.setAccount("david");

        bankingFuncs.findAccount(traffic, true);
    }

    @Test
    void deleteAccount() {
        Traffic traffic = new Traffic();
        traffic.setVerb("DLT");
        traffic.user.setAccount("david");

        bankingFuncs.deleteAccount(traffic, true);
    }

    @Test
    void hashLedger() throws NoSuchAlgorithmException {
        Traffic traffic = new Traffic();
        traffic.setRole("TEST");
        traffic.setVerb("CRT");
        traffic.user.setAmount("200");
        traffic.user.setAccount("david");
        traffic.user.setPassword("mypassword");

        Traffic trafficResponse = bankingFuncs.hashLedger(traffic, true);
//        assertEquals("insert successful", traffic.getMessage());
    }

    @Test
    void opLog() {
        Traffic traffic = new Traffic();
        traffic.setRole("TEST");
        traffic.setVerb("CRT");
        traffic.user.setAmount("200");
        traffic.user.setAccount("david");
        traffic.user.setPassword("mypassword");

        Traffic trafficResponse = bankingFuncs.opLog(traffic, true);
//        assertEquals("insert successful", traffic.getMessage());
    }
}