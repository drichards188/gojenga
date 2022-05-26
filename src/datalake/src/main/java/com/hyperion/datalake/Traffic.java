package com.hyperion.datalake;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;

@JsonIgnoreProperties(ignoreUnknown = true)
public class Traffic {
    public String Name;
    public String Amount;
    public String Account;
    public String Account2;
    public String Verb;
    public String Role;
    public String Port;
    public String Payload;

    public String getAccount() {
        return Account;
    }

    public void setAccount(String account) {
        Account = account;
    }

    public String getAmount() {
        return Amount;
    }

    public void setAmount(String amount) {
        Amount = amount;
    }

    public String getVerb() {
        return Verb;
    }

    public void setVerb(String verb) {
        Verb = verb;
    }

    public void clear() {
        this.Name = "";
        this.Amount = "";
        this.Account = "";
        this.Account2 = "";
        this.Verb = "";
        this.Role = "";
        this.Port = "";
        this.Payload = "";
    }
}
