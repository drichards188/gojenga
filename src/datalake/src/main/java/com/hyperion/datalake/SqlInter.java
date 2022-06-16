package com.hyperion.datalake;

import java.sql.*;
import java.text.SimpleDateFormat;
import java.util.ArrayList;
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
                    sqlQuery(stmt, "ledger", traffic);
                    break;
                }
                case "RECENT": {
                    sqlGetMostRecent(stmt, "ledger", traffic);

                    break;
                }
                case "ALL": {
                    sqlQueryAll(stmt, traffic);
                    break;
                }
                case "CRT":
                    sqlInsertLedger(stmt, "ledger", traffic);
                    sqlInsertUser(stmt, "user", traffic);
                    break;
                case "UPDATE": {
                    sqlUpdate(stmt, "ledger", traffic);
                    break;
                }
                case "HASH": {
                    sqlInsertHash(stmt, "hashHistory", traffic);
                    break;
                }
                case "DELETE": {
                    sqlDelete(stmt, "ledger", traffic);

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
    private Traffic sqlInsertLedger(Statement stmt, String tableName, Traffic traffic) throws SQLException {
        logger.debug("running ledger insert");
        String query = "INSERT INTO " + tableName + " (account, amount) VALUES ('" + traffic.user.getAccount() + "', " + traffic.user.getAmount() + ");";
        int rs = stmt.executeUpdate(query);

        traffic.setMessage("insert successful");

        return traffic;
    }

    private Traffic sqlInsertUser(Statement stmt, String tableName, Traffic traffic) throws SQLException {
        logger.debug("running user insert");
        String query = "INSERT INTO " + tableName + " (account, password) VALUES ('" + traffic.user.getAccount() + "', '" + traffic.user.getPassword() + "');";
        int rs = stmt.executeUpdate(query);

        traffic.setMessage("insert successful");

        return traffic;
    }

    private Traffic sqlInsertHash(Statement stmt, String tableName, Traffic traffic) throws SQLException {
        logger.debug("running user insert");
        String query = "INSERT INTO " + tableName + " (timestamp, previousHash, hash, ledger) VALUES ('" + traffic.hash.getTimestamp() + "', '" + traffic.hash.getPreviousHash() + "', '" + traffic.hash.getHash() + "', '" + traffic.hash.getLedger() + "');";
        int rs = stmt.executeUpdate(query);

        traffic.setMessage("insert successful");

        return traffic;
    }

    private Traffic sqlQuery(Statement stmt, String tableName, Traffic traffic) throws SQLException {
        String QUERY = "SELECT _id, account, amount FROM " + tableName + "  WHERE account= '" + traffic.user.getAccount() + "'";

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

    private ArrayList<User> sqlQueryAll(Statement stmt, Traffic traffic) throws SQLException {
//        String QUERY = "SELECT _id, account, amount FROM ledger;";
//        ResultSet rs = stmt.executeQuery(QUERY);
//
//        ArrayList<User> userArrayList =
//
//        while (rs.next()) {
//            // Retrieve by column name
//            System.out.print("ID: " + rs.getInt("_id"));
//            traffic.setId(rs.getInt("_id"));
//            System.out.print(", Account: " + rs.getString("account"));
//            traffic.user.setAccount(rs.getString("account"));
//            System.out.print(", Amount: " + rs.getString("amount"));
//            traffic.user.setAmount(rs.getString("amount"));
//        }

        ArrayList<User> userArrayList = new ArrayList<User>();
        return userArrayList;
    }

    private Traffic sqlUpdate(Statement stmt, String tableName, Traffic traffic) throws SQLException {
        logger.debug("running insert");
        String query = "UPDATE " + tableName + "  SET amount=" + traffic.user.getAmount() + " WHERE account='" + traffic.user.getAccount() + "';";
        int rs = stmt.executeUpdate(query);
        traffic.setMessage("insert successful");

        return traffic;
    }

    private Traffic sqlDelete(Statement stmt, String tableName, Traffic traffic) throws SQLException {
        logger.debug("running delete");
        String query = "DELETE FROM " + tableName + " WHERE account='" + traffic.user.getAccount() + "';";
        int rs = stmt.executeUpdate(query);
        traffic.setMessage("delete successful");

        return traffic;
    }

    private Traffic sqlGetMostRecent(Statement stmt, String tableName, Traffic traffic) throws SQLException {
        Timestamp timestamp = new Timestamp(System.currentTimeMillis());
        String myTime = new SimpleDateFormat("yyyy-MM-dd HH:mm:ss").format(timestamp);

        String QUERY = "SELECT previousHash, hash, iteration from hashHistory ORDER BY iteration DESC LIMIT 1;";
        ResultSet rs = stmt.executeQuery(QUERY);

        if (rs.next()) {
//            while (rs.next()) {
                // Retrieve by column name

                System.out.print(", PreviousHash: " + rs.getString("previousHash"));
                traffic.hash.setPreviousHash(rs.getString("previousHash"));
                System.out.print(", Hash: " + rs.getString("hash"));
                traffic.hash.setHash(rs.getString("hash"));
                System.out.print(", Iteration: " + rs.getString("iteration"));
                traffic.hash.setIteration(rs.getInt("iteration"));
//            }
        } else {
            logger.info("sqlGetMostRecent result is empty");
        }

        return traffic;
    }
}

