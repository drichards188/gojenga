package com.hyperion.datalake;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import org.springframework.data.mongodb.core.mapping.Document;

@Document(collection = "ledger")

@JsonIgnoreProperties(ignoreUnknown = true)
public class Blockchain {
    public String amount;
    public String account;
    public String account2;
    public String verb;
    public String role;
    public String port;
    public String payload;
    public String message;

    public Blockchain() {

    }

public Blockchain(String account, String amount) {
        this.account = account;
        this.amount = amount;
        }


    public String getAccount() {
        return account;
    }

    public void setAccount(String account) {
        this.account = account;
    }

    public String getAccount2() {
        return account2;
    }

    public void setAccount2(String account2) {
        this.account2 = account2;
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
        this.account = "";
        this.account2 = "";
        this.verb = "";
        this.role = "";
        this.port = "";
        this.payload = "";
    }

    public String toHashString() {
        return "account"+account+"amount"+amount;
    }
}


////package com.hyperion.datalake;
//
//import org.springframework.data.annotation.Id;
//import org.springframework.data.mongodb.core.mapping.Document;
//
//@Document(collection = "ledger")
//public class Tutorial {
//    @Id
//    private String id;
//
//    private String account;
//    private String amount;
//    private String verb;
//    private String message;
//    private String account2;
//
//    public Tutorial() {
//
//    }
//
//    public Tutorial(String account, String amount) {
//        this.account = account;
//        this.amount = amount;
//
////        this.title = title;
////        this.description = description;
////        this.published = published;
//    }
//
//    public String getmessage() {
//        return message;
//    }
//
//    public String getaccount2() {
//        return account2;
//    }
//
//    public void setaccount2(String account2) {
//        this.account2 = account2;
//    }
//
//    public void setverb(String verb) {
//        this.verb = verb;
//    }
//
//    public String getaccount() {
//        return account;
//    }
//
//    public String getamount() {
//        return amount;
//    }
//
//    public String getverb() {
//        return verb;
//    }
//
//    public String getId() {
//        return id;
//    }
//
//    public void setamount(String amount) {
//        this.amount = amount;
//    }
//
//    public void setaccount(String account) {
//        this.account = account;
//    }
//
//    public void setmessage(String message) {
//        this.message = message;
//    }
//
////    @Override
////    public String toString() {
////        return "Tutorial [id=" + id + ", title=" + title + ", desc=" + description + ", published=" + published + "]";
////    }
//
//    public String toHashString() {
//        return "account"+account+"amount"+amount;
//    }
//}