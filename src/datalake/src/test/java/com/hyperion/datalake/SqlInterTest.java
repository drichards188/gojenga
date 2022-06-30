package com.hyperion.datalake;

import org.junit.jupiter.api.AfterEach;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SqlInterTest {
    private SqlInter sqlInter;

    @BeforeEach
    void setUp() {
        sqlInter = new SqlInter();
    }

    @AfterEach
    void tearDown() {
    }

    @Test
    void sqlInsertLedger() {
        Traffic traffic = new Traffic();
        traffic.setRole("TEST");
        traffic.setVerb("CRT");
        traffic.user.setAmount("200");
        traffic.user.setAccount("david");
        traffic.user.setPassword("mypassword");

        Traffic trafficResponse = sqlInter.sqlHandler(BankingFuncs.Crud.CREATE, BankingFuncs.Datatypes.USER, traffic);
        assertEquals("insert successful", traffic.getMessage());
    }
}