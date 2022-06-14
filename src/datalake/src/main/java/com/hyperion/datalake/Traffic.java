package com.hyperion.datalake;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;

@JsonIgnoreProperties(ignoreUnknown = true)
public class Traffic {
    public String Name;
    public String Amount;
    public String Account;
    public String SourceAccount;
    public String DestinationAccount;
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

    public String getSource() {
        return SourceAccount;
    }

    public void setSource(String source) {
        SourceAccount = source;
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

    public String getDestinationAccount() {
        return DestinationAccount;
    }

    public void setDestinationAccount(String destinationAccount) {
        DestinationAccount = destinationAccount;
    }

    public void clear() {
        this.Name = "";
        this.Amount = "";
        this.SourceAccount = "";
        this.DestinationAccount = "";
        this.Verb = "";
        this.Role = "";
        this.Port = "";
        this.Payload = "";
    }
}
