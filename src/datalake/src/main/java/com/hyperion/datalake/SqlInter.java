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

    public Traffic sqlHandler(String verb, Traffic traffic) {
        String url = "jdbc:mysql://localhost:3306/crypto?useJDBCCompliantTimezoneShift=true&useLegacyDatetimeCode=false&serverTimezone=UTC";
        String username = "david";
        String password = "password";

        System.out.println("Connecting database...");

        try (Connection connection = DriverManager.getConnection(url, username, password)) {

            Statement stmt = connection.createStatement();
            System.out.println("Database connected!");

            String tableName = "ledger";
            String accountName = traffic.user.getAccount();

            switch (verb) {
                case "QUERY": {
                    sqlQuery(stmt, traffic);
                    break;
                }
                case "ALL": {
                    sqlQueryAll(stmt, traffic);
                    break;
                }
                case "CRT":
                    sqlInsert(stmt, verb, traffic);
                    break;
                case "UPDATE": {
                    String query = "UPDATE " + tableName + " SET amount=" + traffic.user.getAmount() + " WHERE account='" + accountName + "';";
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

            return traffic;

        } catch (SQLException e) {
            throw new IllegalStateException("Cannot connect the database!", e);
        }
    }

    private Traffic sqlInsert(Statement stmt, String verb, Traffic traffic) throws SQLException {
        String tableName = "ledger";
        String accountName = traffic.user.getAccount();

        String query = "INSERT INTO " + tableName + " (account, amount) VALUES ('" + accountName + "', " + traffic.user.getAmount() + ");";
        int rs = stmt.executeUpdate(query);
        logger.debug("running insert");
        return traffic;
    }

    private Traffic sqlQuery(Statement stmt, Traffic traffic) throws SQLException {
        String QUERY = "SELECT _id, account, amount FROM ledger WHERE account= '" + traffic.user.getAccount() + "'";
        ResultSet rs = stmt.executeQuery(QUERY);

        while (rs.next()) {
            // Retrieve by column name
            System.out.print("ID: " + rs.getInt("_id"));
            traffic.setId(rs.getInt("_id"));
            System.out.print(", Account: " + rs.getString("account"));
            traffic.user.setAccount(rs.getString("account"));
            System.out.print(", Amount: " + rs.getString("amount"));
            traffic.user.setAmount(rs.getString("amount"));
        }

        return traffic;
    }

    private Traffic sqlQueryAll(Statement stmt, Traffic traffic) throws SQLException {
        String QUERY = "SELECT _id, account, amount FROM ledger;";
        ResultSet rs = stmt.executeQuery(QUERY);

        while (rs.next()) {
            // Retrieve by column name
            System.out.print("ID: " + rs.getInt("_id"));
            traffic.setId(rs.getInt("_id"));
            System.out.print(", Account: " + rs.getString("account"));
            traffic.user.setAccount(rs.getString("account"));
            System.out.print(", Amount: " + rs.getString("amount"));
            traffic.user.setAmount(rs.getString("amount"));
        }

        return traffic;
    }
}

