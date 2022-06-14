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

private final Logger logger = LoggerFactory.getLogger(this.getClass());


    public SqlInter() throws Exception {

    }

    public void sqlHandler(String verb, Traffic traffic) {
        String url = "jdbc:mysql://localhost:3306/crypto?useJDBCCompliantTimezoneShift=true&useLegacyDatetimeCode=false&serverTimezone=UTC";
        String username = "david";
        String password = "password";

        System.out.println("Connecting database...");

        try (Connection connection = DriverManager.getConnection(url, username, password)) {

            Statement stmt = connection.createStatement();
            System.out.println("Database connected!");

            this.sqlCommands(stmt, verb, traffic);

        } catch (SQLException e) {
            throw new IllegalStateException("Cannot connect the database!", e);
        }
    }

    private void sqlInsert(Statement stmt, String verb, Traffic traffic) throws SQLException {
        String tableName = "ledger";
        String accountName = traffic.getAccount();

        String query = "INSERT INTO " + tableName + " (account, amount) VALUES ('" + accountName + "', "+ traffic.getAmount() +");";
        int rs = stmt.executeUpdate(query);
        logger.debug("running insert");
    }

    public void sqlCommands(Statement stmt, String verb, Traffic traffic) throws SQLException {
        String tableName = "ledger";
        String accountName = traffic.getAccount();

        switch (verb) {
            case "SELECT": {
                String QUERY = "SELECT _id, account, amount FROM ledger;";
                ResultSet rs = stmt.executeQuery(QUERY);

                while (rs.next()) {
                    // Retrieve by column name
                    System.out.print("ID: " + rs.getInt("_id"));
                    System.out.print(", Account: " + rs.getString("account"));
                    System.out.print(", Amount: " + rs.getString("amount"));
                }
                break;
            }
            case "INSERT":
                sqlInsert(stmt, verb, traffic);
                break;
            case "UPDATE": {
                String query = "UPDATE " + tableName + " SET amount=" + traffic.getAmount() + " WHERE account='" + accountName + "';";
                int rs = stmt.executeUpdate(query);
                logger.debug("running insert");
                break;
            }
            case "DELETE": {
                String query = "DELETE FROM " + tableName + " WHERE account='" + accountName + "';";
                int rs = stmt.executeUpdate(query);
                logger.debug("running insert");
                break;
            }
            case "CREATE":
                try {
                    String query = "CREATE TABLE " + tableName + " ("
                            + "idNo INT(64) NOT NULL AUTO_INCREMENT,"
                            + "initials VARCHAR(2),"
                            + "agentDate DATE,"
                            + "agentCount INT(64), "
                            + "PRIMARY KEY(idNo))";
                    int rs = stmt.executeUpdate(query);
                    logger.debug("running insert");
                } catch (Exception e) {
                    logger.error("handlePost triggered exception: " + e);

                }
                break;
        }
    }
}

