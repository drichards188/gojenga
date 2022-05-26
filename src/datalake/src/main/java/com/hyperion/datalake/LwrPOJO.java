package com.hyperion.datalake;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import org.springframework.data.mongodb.core.mapping.Document;

@Document(collection = "ledger")

@JsonIgnoreProperties(ignoreUnknown = true)
public class LwrPOJO {
    public String amount;
    public String account;
    public String account2;
    public String verb;
    public String Role;
    public String Port;
    public String Payload;
    public String Message;

    public LwrPOJO() {

    }

    public LwrPOJO(String account, String amount) {
        this.account = account;
        this.amount = amount;
    }



    public String getaccount() {
        return account;
    }

    public void setaccount(String account) {
        this.account = account;
    }

    public String getaccount2() {
        return account2;
    }

    public void setaccount2(String account2) {
        this.account2 = account2;
    }

    public String getamount() {
        return amount;
    }

    public void setamount(String amount) {
        this.amount = amount;
    }

    public String getverb() {
        return verb;
    }

    public String getMessage() {
        return Message;
    }

    public void setMessage(String message) {
        this.Message = message;
    }

    public void setverb(String verb) {
        this.verb = verb;
    }

    public void clear() {
        this.amount = "";
        this.account = "";
        this.account2 = "";
        this.verb = "";
        this.Role = "";
        this.Port = "";
        this.Payload = "";
    }

    public String toHashString() {
        return "account"+account+"amount"+amount;
    }
}
