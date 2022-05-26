package com.hyperion.datalake;

import java.sql.*;
import java.util.Vector;

import com.fasterxml.jackson.databind.ObjectMapper;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;

import java.io.File;
import java.io.IOException;

public class SqlInter {
//TRAN
//CRT
//QUERY
//DISCO
//HASH
//PING
private final Logger logger = LoggerFactory.getLogger(this.getClass());
    public static void main() throws Exception {

        SqlInter sqlInter = new SqlInter();
        String url = "jdbc:mysql://localhost:3306/crypto";
        String username = "david";
        String password = "password";

        System.out.println("Connecting database...");

        try (Connection connection = DriverManager.getConnection(url, username, password)) {

            Statement stmt = connection.createStatement();
            System.out.println("Database connected!");

            String[] data = {"hiya"};
            sqlInter.sqlCommands(stmt, "UPDATE", data);
        } catch (SQLException e) {
            throw new IllegalStateException("Cannot connect the database!", e);
        }

//        Server.createServer();
    }

    public void sqlCommands(Statement stmt, String verb, String[] data) throws SQLException {
        String tableName = "ledger";
        String accountName = "tucker";
        verb = "SELECT";
        Integer amount = 200;
        if (verb.equals("SELECT") ) {
            String QUERY = "SELECT _id, account, amount FROM ledger;";

            ResultSet rs = stmt.executeQuery(QUERY);

            while (rs.next()) {
                // Retrieve by column name
                System.out.print("ID: " + rs.getInt("_id"));
                System.out.print(", Account: " + rs.getString("account"));
                System.out.print(", Amount: " + rs.getString("amount"));
            }
        }

        else if (verb.equals("INSERT")) {

            String query = "INSERT INTO " + tableName + " (account, amount) VALUES ('" + accountName + "', "+ amount +");";
            int rs = stmt.executeUpdate(query);
            logger.debug("running insert");
        }

        else if (verb.equals("UPDATE")) {
            String query = "UPDATE " + tableName + " SET amount="+amount+" WHERE account='"+accountName+"';";
            int rs = stmt.executeUpdate(query);
            logger.debug("running insert");
        }

        else if (verb.equals("DELETE")) {
            String query = "DELETE FROM " + tableName + " WHERE account='"+accountName+"';";
            int rs = stmt.executeUpdate(query);
            logger.debug("running insert");
        }

        else if (verb.equals("CREATE")) {
            try {
                String query = "CREATE TABLE " + tableName + " ("
                        + "idNo INT(64) NOT NULL AUTO_INCREMENT,"
                        + "initials VARCHAR(2),"
                        + "agentDate DATE,"
                        + "agentCount INT(64), "
                        + "PRIMARY KEY(idNo))";
                int rs = stmt.executeUpdate(query);
                logger.debug("running insert");
            }
            catch (Exception e) {
                logger.error("handlePost triggered exception: " + e);

            }
        }
    }
}

