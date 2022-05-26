package com.hyperion.datalake;

import org.springframework.data.annotation.Id;
import org.springframework.data.mongodb.core.mapping.Document;

@Document(collection = "hashHistory")
public class Hash {
    @Id
    private String id;

    private Integer iteration;
    private String timestamp;
    private String hash;
    private String previousHash;
    private String ledger;

    private String account;
    private String amount;
    private String verb;
    private String message;
    private String account2;

    public Hash() {

    }

    public Hash(String account, String amount) {
        this.account = account;
        this.amount = amount;

//        this.title = title;
//        this.description = description;
//        this.published = published;
    }

    public void setVerb(String verb) {
        this.verb = verb;
    }

    public String getLedger() {
        return ledger;
    }

    public void setLedger(String ledger) {
        this.ledger = ledger;
    }

    public String getPreviousHash() {
        return previousHash;
    }

    public void setPreviousHash(String previousHash) {
        this.previousHash = previousHash;
    }

    public String getHash() {
        return hash;
    }

    public void setHash(String Hash) {
        this.hash = Hash;
    }

    public String getTimestamp() {
        return timestamp;
    }

    public void setTimestamp(String timestamp) {
        this.timestamp = timestamp;
    }

    public Integer getIteration() {
        return iteration;
    }

    public void setIteration(Integer iteration) {
        this.iteration = iteration;
    }

    public String getmessage() {
        return message;
    }

    public String getAccount2() {
        return account2;
    }

    public void setAccount2(String account2) {
        this.account2 = account2;
    }

    public String getAccount() {
        return account;
    }

    public String getAmount() {
        return amount;
    }

    public String getVerb() {
        return verb;
    }

    public String getId() {
        return id;
    }


    public void setAmount(String amount) {
        this.amount = amount;
    }

    public void setAccount(String account) {
        this.account = account;
    }

    public void setMessage(String message) {
        this.message = message;
    }

    public String toHashString() {
        return "account"+account+"amount"+amount;
    }

//    @Override
//    public String toString() {
//        return "Tutorial [id=" + id + ", title=" + title + ", desc=" + description + ", published=" + published + "]";
//    }
}