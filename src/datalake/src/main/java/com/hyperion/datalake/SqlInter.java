package com.hyperion.datalake;

import java.sql.*;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class SqlInter {

    private final Logger logger = LoggerFactory.getLogger(this.getClass());


    public SqlInter() {

    }

    public Traffic sqlHandler(BankingFuncs.Crud verb, BankingFuncs.Datatypes datatype, Traffic traffic) {
        String url = "jdbc:mysql://localhost:3306/crypto?useJDBCCompliantTimezoneShift=true&useLegacyDatetimeCode=false&serverTimezone=UTC";
        String username = "david";
        String password = "password";

        System.out.println("Connecting database...");

        try (Connection connection = DriverManager.getConnection(url, username, password)) {

            Statement stmt = connection.createStatement();
            System.out.println("Database connected!");

            switch (verb) {
                case CREATE -> {
                    switch (datatype) {
                        case LEDGER -> {
                            Traffic result = sqlInsertLedger(stmt, "ledger", traffic);
                            if (result.getFail().equals(true)) {
                                logger.info("case Ledger result: " + result.getFailMessage());
                            }

                            return result;
                        }
                        case USER -> {
                            Traffic result = sqlInsertUser(stmt, "user", traffic);
                            if (result.getFail().equals(true)) {
                                logger.info("case user result: " + result.getFailMessage());
                            }

                            return result;
                        }
                        case HASH -> {
                            Traffic result = sqlInsertHash(stmt, "hashHistory", traffic);
                            if (result.getFail().equals(true)) {
                                logger.info("case hash result: " + result.getFailMessage());
                            }

                            return result;
                        }
                        case OPLOG -> {
                            Traffic result = sqlInsertOplog(stmt, "opsLog", traffic);
                            if (result.getFail().equals(true)) {
                                logger.info("case oplog result: " + result.getFailMessage());
                            }

                            return result;
                        }
                    }

                }
                case READ -> {
                    switch (datatype) {
                        case LEDGER -> {
                            Traffic result = sqlQueryLedger(stmt, "ledger", traffic);
                            if (result.getFail().equals(true)) {
                                logger.info("case Ledger result: " + result.getFailMessage());
                            }

                            return result;
                        }
                        case USER -> {
                            Traffic result = sqlQueryUser(stmt, "user", traffic);
                            if (result.getFail().equals(true)) {
                                logger.info("case user result: " + result.getFailMessage());
                            }

                            return result;
                        }
                        case HASH -> {
                            Traffic result = sqlQueryHash(stmt, "hashHistory", traffic);
                            if (result.getFail().equals(true)) {
                                logger.info("case hash result: " + result.getFailMessage());
                            }

                            return result;
                        }
                        case OPLOG -> {
                            Traffic result = sqlQueryOplog(stmt, "opsLog", traffic);
                            if (result.getFail().equals(true)) {
                                logger.info("case oplog result: " + result.getFailMessage());
                            }

                            return result;
                        }
                    }
                }
                case UPDATE -> {
                    switch (datatype) {
                        case LEDGER -> {
                            Traffic result = sqlUpdateLedger(stmt, "ledger", traffic);
                            if (result.getFail().equals(true)) {
                                logger.info("case Ledger result: " + result.getFailMessage());
                            }

                            return result;
                        }
                        case USER -> {
                            Traffic result = sqlUpdateUser(stmt, "user", traffic);
                            if (result.getFail().equals(true)) {
                                logger.info("case user result: " + result.getFailMessage());
                            }

                            return result;
                        }
                        case HASH -> {
                            Traffic result = sqlUpdateHash(stmt, "hashHistory", traffic);
                            if (result.getFail().equals(true)) {
                                logger.info("case hash result: " + result.getFailMessage());
                            }

                            return result;
                        }
                        case OPLOG -> {
                            Traffic result = sqlUpdateOplog(stmt, "opsLog", traffic);
                            if (result.getFail().equals(true)) {
                                logger.info("case oplog result: " + result.getFailMessage());
                            }

                            return result;
                        }
                    }
                }
                case DELETE -> {
                    switch (datatype) {
                        case LEDGER -> {
                            Traffic result = sqlDeleteLedger(stmt, "ledger", traffic);
                            if (result.getFail().equals(true)) {
                                logger.info("case Ledger result: " + result.getFailMessage());
                            }

                            return result;
                        }
                        case USER -> {
                            Traffic result = sqlDeleteUser(stmt, "user", traffic);
                            if (result.getFail().equals(true)) {
                                logger.info("case user result: " + result.getFailMessage());
                            }

                            return result;
                        }
                        case HASH -> {
                            Traffic result = sqlDeleteHash(stmt, "hashHistory", traffic);
                            if (result.getFail().equals(true)) {
                                logger.info("case hash result: " + result.getFailMessage());
                            }

                            return result;
                        }
                        case OPLOG -> {
                            Traffic result = sqlDeleteOplog(stmt, "opsLog", traffic);
                            if (result.getFail().equals(true)) {
                                logger.info("case oplog result: " + result.getFailMessage());
                            }

                            return result;
                        }
                    }
                }
            }

            return traffic;

        } catch (SQLException e) {
            throw new IllegalStateException("Cannot connect the database!", e);
        }
    }

    //the insert methods
    private Traffic sqlInsertLedger(Statement stmt, String tableName, Traffic traffic) throws SQLException {

        int rs = 0;
        if (traffic.getRole().equals("TEST")) {
            //todo mock sql response
            rs = 1;
        } else {
            logger.debug("running ledger insert");
            String query = "INSERT INTO " + tableName + " (account, amount) VALUES ('" + traffic.user.getAccount() + "', " + traffic.user.getAmount() + ");";
            rs = stmt.executeUpdate(query);
        }

        if (rs == 0) {
            traffic.setFail(true);
            traffic.setFailMessage("sqlInsertLedger failed");
            return traffic;
        }

        traffic.setMessage("insert successful");

        return traffic;
    }

    private Traffic sqlInsertUser(Statement stmt, String tableName, Traffic traffic) throws SQLException {
        logger.debug("running user insert");
        String query = "INSERT INTO " + tableName + " (account, password) VALUES ('" + traffic.user.getAccount() + "', '" + traffic.user.getPassword() + "');";
        int rs = stmt.executeUpdate(query);

        if (rs == 0) {
            traffic.setFail(true);
            traffic.setFailMessage("sqlInsertUser failed");
            return traffic;
        }

        traffic.setMessage("insert successful");

        return traffic;
    }

    private Traffic sqlInsertHash(Statement stmt, String tableName, Traffic traffic) throws SQLException {
        logger.debug("running user insert");
        String query = "INSERT INTO " + tableName + " (timestamp, previousHash, hash, ledger) VALUES ('" + traffic.hash.getTimestamp() + "', '" + traffic.hash.getPreviousHash() + "', '" + traffic.hash.getHash() + "', '" + traffic.hash.getLedger() + "');";
        int rs = stmt.executeUpdate(query);

        if (rs == 0) {
            traffic.setFail(true);
            traffic.setFailMessage("sqlInsertHash failed");
            return traffic;
        }

        traffic.setMessage("insert successful");

        return traffic;
    }

    private Traffic sqlInsertOplog(Statement stmt, String tableName, Traffic traffic) throws SQLException {
        logger.debug("running user insert");

        if (traffic.getSourceAccount() == null) {
            traffic.setSourceAccount(traffic.user.getAccount());
        }

        if (traffic.getVerb().equals("HASH") || traffic.getVerb().equals("DLT")) {
            traffic.user.setAmount("0");
        }

        String query = "INSERT INTO " + tableName + " (operation, source, destination, amount, fail) VALUES ('" + traffic.getVerb() + "', '" + traffic.getSourceAccount() + "', '" + traffic.getDestinationAccount() + "', '" + traffic.user.getAmount() + "', '" + 0 + "');";
        int rs = stmt.executeUpdate(query);

        if (rs == 0) {
            traffic.setFail(true);
            traffic.setFailMessage("sqlInsertOplog failed");
            return traffic;
        }

        traffic.setMessage("insert successful");

        return traffic;
    }

    //the query methods
    private Traffic sqlQueryLedger(Statement stmt, String tableName, Traffic traffic) throws SQLException {
        String QUERY = "SELECT _id, account, amount FROM " + tableName + "  WHERE account= '" + traffic.user.getAccount() + "';";

        ResultSet rs = stmt.executeQuery(QUERY);

        Boolean isResultsEmpty = true;
        while (rs.next()) {
            // Retrieve by column name
            traffic.setId(rs.getInt("_id"));
            traffic.user.setAccount(rs.getString("account"));
            traffic.user.setAmount(rs.getString("amount"));
            isResultsEmpty = false;
        }

        if (isResultsEmpty.equals(true)) {
            traffic.setFail(true);
            traffic.setFailMessage("sqlQueryLedger no results");
        }

        return traffic;
    }

    private Traffic sqlQueryUser(Statement stmt, String tableName, Traffic traffic) throws SQLException {
        String QUERY = "SELECT _id, account, password FROM " + tableName + "  WHERE account= '" + traffic.user.getAccount() + "';";

        ResultSet rs = stmt.executeQuery(QUERY);

        Boolean isResultsEmpty = true;
        while (rs.next()) {
            // Retrieve by column name
            traffic.setId(rs.getInt("_id"));
            traffic.user.setAccount(rs.getString("account"));
            traffic.user.setPassword(rs.getString("password"));
            isResultsEmpty = false;
        }

        if (isResultsEmpty.equals(true)) {
            traffic.setFail(true);
            traffic.setFailMessage("sqlQueryUser no results");
        }

        return traffic;
    }

    private Traffic sqlQueryHash(Statement stmt, String tableName, Traffic traffic) throws SQLException {
        String QUERY = "SELECT iteration, previousHash, hash, ledger, timestamp FROM " + tableName + "  WHERE iteration= '" + traffic.hash.getIteration() + "';";

        ResultSet rs = stmt.executeQuery(QUERY);

        Boolean isResultsEmpty = true;
        while (rs.next()) {
            // Retrieve by column name
            traffic.hash.setIteration(rs.getInt("iteration"));
            traffic.hash.setPreviousHash(rs.getString("previousHash"));
            traffic.hash.setHash(rs.getString("hash"));
            traffic.hash.setLedger(rs.getString("ledger"));
            traffic.hash.setTimestamp(rs.getString("timestamp"));
            isResultsEmpty = false;
        }

        if (isResultsEmpty.equals(true)) {
            traffic.setFail(true);
            traffic.setFailMessage("sqlQueryHash no results");
        }

        return traffic;
    }

    private Traffic sqlQueryOplog(Statement stmt, String tableName, Traffic traffic) throws SQLException {
        String QUERY = "SELECT _id, operation, source, destination, amount, fail FROM " + tableName + "  WHERE _id= '" + traffic.getPayload() + "';";

        ResultSet rs = stmt.executeQuery(QUERY);

        Boolean isResultsEmpty = true;
        while (rs.next()) {
            // Retrieve by column name
            System.out.print("ID: " + rs.getInt("_id"));
            traffic.setId(rs.getInt("_id"));
            System.out.print(", Account: " + rs.getString("account"));
            traffic.user.setAccount(rs.getString("account"));
            System.out.print(", Amount: " + rs.getString("amount"));
            traffic.user.setAmount(rs.getString("amount"));


            isResultsEmpty = false;
        }

        if (isResultsEmpty.equals(true)) {
            traffic.setFail(true);
            traffic.setFailMessage("sqlQueryOplog no results");
        }

        return traffic;
    }

    //the update methods
    private Traffic sqlUpdateLedger(Statement stmt, String tableName, Traffic traffic) throws SQLException {
        logger.debug("running insert");
        String query = "UPDATE " + tableName + "  SET amount=" + traffic.user.getAmount() + " WHERE account='" + traffic.user.getAccount() + "';";
        int rs = stmt.executeUpdate(query);

        if (rs == 0) {
            traffic.setFail(true);
            traffic.setFailMessage("sqlUpdateLedger failed");
            return traffic;
        }

        traffic.setMessage("update successful");

        return traffic;
    }

    private Traffic sqlUpdateUser(Statement stmt, String tableName, Traffic traffic) throws SQLException {
        logger.debug("running insert");
        String query = "UPDATE " + tableName + "  SET password=" + traffic.user.getPassword() + " WHERE account='" + traffic.user.getAccount() + "';";
        int rs = stmt.executeUpdate(query);

        if (rs == 0) {
            traffic.setFail(true);
            traffic.setFailMessage("sqlUpdateUser failed");
            return traffic;
        }

        traffic.setMessage("update successful");

        return traffic;
    }

    //no use case
    private Traffic sqlUpdateHash(Statement stmt, String tableName, Traffic traffic) throws SQLException {
        logger.debug("running insert");
        String query = "UPDATE " + tableName + "  SET amount=" + traffic.user.getAmount() + " WHERE account='" + traffic.user.getAccount() + "';";
        int rs = stmt.executeUpdate(query);

        if (rs == 0) {
            traffic.setFail(true);
            traffic.setFailMessage("sqlUpdateHash failed");
            return traffic;
        }

        traffic.setMessage("update successful");

        return traffic;
    }

    //no use case
    private Traffic sqlUpdateOplog(Statement stmt, String tableName, Traffic traffic) throws SQLException {
        logger.debug("running insert");
        String query = "UPDATE " + tableName + "  SET amount=" + traffic.user.getAmount() + " WHERE account='" + traffic.user.getAccount() + "';";
        int rs = stmt.executeUpdate(query);

        if (rs == 0) {
            traffic.setFail(true);
            traffic.setFailMessage("sqlUpdateOplog failed");
            return traffic;
        }

        traffic.setMessage("update successful");

        return traffic;
    }

    //the delete methods
    private Traffic sqlDeleteLedger(Statement stmt, String tableName, Traffic traffic) throws SQLException {
        logger.debug("running delete");
        String query = "DELETE FROM " + tableName + " WHERE account='" + traffic.user.getAccount() + "';";
        int rs = stmt.executeUpdate(query);

        if (rs == 0) {
            traffic.setFail(true);
            traffic.setFailMessage("sqlDeleteLedger failed");
            return traffic;
        }

        traffic.setMessage("delete successful");

        return traffic;
    }

    private Traffic sqlDeleteUser(Statement stmt, String tableName, Traffic traffic) throws SQLException {
        logger.debug("running delete");
        String query = "DELETE FROM " + tableName + " WHERE account='" + traffic.user.getAccount() + "';";
        int rs = stmt.executeUpdate(query);

        if (rs == 0) {
            traffic.setFail(true);
            traffic.setFailMessage("sqlDeleteUser failed");
            return traffic;
        }

        traffic.setMessage("delete successful");

        return traffic;
    }

    private Traffic sqlDeleteHash(Statement stmt, String tableName, Traffic traffic) throws SQLException {
        logger.debug("running delete");
        String query = "DELETE FROM " + tableName + " WHERE iteration='" + traffic.hash.getIteration() + "';";
        int rs = stmt.executeUpdate(query);

        if (rs == 0) {
            traffic.setFail(true);
            traffic.setFailMessage("sqlDeleteHash failed");
            return traffic;
        }

        traffic.setMessage("delete successful");

        return traffic;
    }

    private Traffic sqlDeleteOplog(Statement stmt, String tableName, Traffic traffic) throws SQLException {
        logger.debug("running delete");
        String query = "DELETE FROM " + tableName + " WHERE iteration='" + traffic.hash.getIteration() + "';";
        int rs = stmt.executeUpdate(query);

        if (rs == 0) {
            traffic.setFail(true);
            traffic.setFailMessage("sqlDeleteOplog failed");
            return traffic;
        }

        traffic.setMessage("delete successful");

        return traffic;
    }

    //    private ArrayList<User> sqlQueryAll(Statement stmt, Traffic traffic) throws SQLException {
////        String QUERY = "SELECT _id, account, amount FROM ledger;";
////        ResultSet rs = stmt.executeQuery(QUERY);
////
////        ArrayList<User> userArrayList =
////
////        while (rs.next()) {
////            // Retrieve by column name
////            System.out.print("ID: " + rs.getInt("_id"));
////            traffic.setId(rs.getInt("_id"));
////            System.out.print(", Account: " + rs.getString("account"));
////            traffic.user.setAccount(rs.getString("account"));
////            System.out.print(", Amount: " + rs.getString("amount"));
////            traffic.user.setAmount(rs.getString("amount"));
////        }
//
//        ArrayList<User> userArrayList = new ArrayList<User>();
//        return userArrayList;
//    }

//    private Traffic sqlGetMostRecent(Statement stmt, String tableName, Traffic traffic) throws SQLException {
//        Timestamp timestamp = new Timestamp(System.currentTimeMillis());
//        String myTime = new SimpleDateFormat("yyyy-MM-dd HH:mm:ss").format(timestamp);
//
//        String QUERY = "SELECT previousHash, hash, iteration from hashHistory ORDER BY iteration DESC LIMIT 1;";
//        ResultSet rs = stmt.executeQuery(QUERY);
//
//        if (rs.next()) {
////            while (rs.next()) {
//            // Retrieve by column name
//
//            System.out.print(", PreviousHash: " + rs.getString("previousHash"));
//            traffic.hash.setPreviousHash(rs.getString("previousHash"));
//            System.out.print(", Hash: " + rs.getString("hash"));
//            traffic.hash.setHash(rs.getString("hash"));
//            System.out.print(", Iteration: " + rs.getString("iteration"));
//            traffic.hash.setIteration(rs.getInt("iteration"));
////            }
//        } else {
//            logger.info("sqlGetMostRecent result is empty");
//        }
//
//        return traffic;
//    }
}

