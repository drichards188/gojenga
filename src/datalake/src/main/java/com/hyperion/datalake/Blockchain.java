package com.hyperion.datalake;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import org.springframework.data.mongodb.core.mapping.Document;

@Document(collection = "ledger")

@JsonIgnoreProperties(ignoreUnknown = true)
public class Blockchain {
    public String amount;
    public String sourceAccount;
    public String destinationAccount;
    public String verb;
    public String role;
    public String port;
    public String payload;
    public String message;

    public Blockchain() {

    }

public Blockchain(String sourceAccount, String amount) {
        this.sourceAccount = sourceAccount;
        this.amount = amount;
        }


    public String getsourceAccount() {
        return sourceAccount;
    }

    public void setsourceAccount(String sourceAccount) {
        this.sourceAccount = sourceAccount;
    }

    public String getdestinationAccount() {
        return destinationAccount;
    }

    public void setdestinationAccount(String destinationAccount) {
        this.destinationAccount = destinationAccount;
    }

    public String getAmount() {
        return amount;
    }

    public void setAmount(String amount) {
        this.amount = amount;
    }

    public String getVerb() {
        return verb;
    }

    public String getMessage() {
        return message;
    }

    public void setMessage(String message) {
        this.message = message;
    }

    public void setVerb(String verb) {
        this.verb = verb;
    }

    public void clear() {
        this.amount = "";
        this.sourceAccount = "";
        this.destinationAccount = "";
        this.verb = "";
        this.role = "";
        this.port = "";
        this.payload = "";
    }

    public String toHashString() {
        return "sourceAccount"+sourceAccount+"amount"+amount;
    }
}