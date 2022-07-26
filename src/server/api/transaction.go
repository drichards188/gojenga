package api

import (
	"context"
	"errors"
	"fmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"strconv"
	"time"
)

//func testingFunc() (throwError bool) {
//	logger = InitializeLogger()
//	ctx := context.Background()
//
//	traffic := Traffic{SourceAccount: "david", DestinationAccount: "allie", Table: "dynamoTest", Role: "test", Amount: "12"}
//
//	resp, err := Transaction(traffic, ctx)
//	if err != nil {
//		logger.Warn(fmt.Sprintf("gjDeposit test error: %s", err))
//		return true
//	}
//
//	logger.Debug(fmt.Sprintf("gjDeposit test returned: %s", resp))
//
//	return false
//}

func Transaction(jsonResponse Traffic, ctx context.Context) (string, error) {
	user1, err := RunDynamoGetItem(Query{TableName: "ledger", Key: "Account", Value: jsonResponse.SourceAccount})
	if err != nil {
		return "--> " + user1["msg"], errors.New("--> " + user1["msg"])
	}
	user2, err := RunDynamoGetItem(Query{TableName: "ledger", Key: "Account", Value: jsonResponse.DestinationAccount})
	if err != nil {
		return "--> " + user2["msg"], errors.New("--> " + user1["msg"])
	}

	Account := jsonResponse.SourceAccount
	Amount := jsonResponse.Amount
	Account2 := jsonResponse.DestinationAccount

	//traffic := Traffic{Account: Account2}

	time.Sleep(1 * time.Second)

	//mongoResult2 := queryMongo(traffic)
	//resultMap2 := mongoResult2.Map()
	finalAmount, err := strconv.Atoi(Amount)
	if err != nil {
		fmt.Println(err)
	}
	Amount1, err := strconv.Atoi(user1["Amount"])
	if err != nil {
		fmt.Println(err)
	}
	Amount2, err := strconv.Atoi(user2["Amount"])
	if err != nil {
		fmt.Println(err)
	}
	finalAmount1 := Amount1 - finalAmount
	finalAmount2 := Amount2 + finalAmount

	intFinalAmount1 := strconv.Itoa(finalAmount1)
	intFinalAmount2 := strconv.Itoa(finalAmount2)

	//data := Account + Account2 + Amount
	tr := otel.Tracer("crypto-trace")
	ctx, span := tr.Start(ctx, "handle-gjTransaction")
	span.SetAttributes(attribute.Key("testset").String("value"))
	defer span.End()
	//lakeResponse := acceptTransaction(jsonResponse, ctx)
	////lakeResponse := lakeTransaction(jsonResponse)
	//return lakeResponse

	//hashLedger(data)

	r, err := RunDynamoCreateItem("ledger", Ledger{Account: Account, Amount: intFinalAmount1})
	if err != nil {
		return "--> " + r["msg"], errors.New("--> " + r["msg"])
	}
	r, err = RunDynamoCreateItem("ledger", Ledger{Account: Account2, Amount: intFinalAmount2})
	if err != nil {
		return "--> " + r["msg"], errors.New("--> " + r["msg"])
	}

	return "Transaction Successful", nil
}

func TransactionRollback(jsonResponse Traffic, ctx context.Context) (string, error) {
	user1, err := RunDynamoGetItem(Query{TableName: "ledger", Key: "Account", Value: jsonResponse.SourceAccount})
	if err != nil {
		return "--> " + user1["msg"], errors.New("--> " + user1["msg"])
	}
	user2, err := RunDynamoGetItem(Query{TableName: "ledger", Key: "Account", Value: jsonResponse.DestinationAccount})
	if err != nil {
		return "--> " + user2["msg"], errors.New("--> " + user1["msg"])
	}

	Account := jsonResponse.SourceAccount
	Amount := jsonResponse.Amount
	Account2 := jsonResponse.DestinationAccount

	//traffic := Traffic{Account: Account2}

	time.Sleep(1 * time.Second)

	//mongoResult2 := queryMongo(traffic)
	//resultMap2 := mongoResult2.Map()
	finalAmount, err := strconv.Atoi(Amount)
	if err != nil {
		fmt.Println(err)
	}
	Amount1, err := strconv.Atoi(user1["Amount"])
	if err != nil {
		fmt.Println(err)
	}
	Amount2, err := strconv.Atoi(user2["Amount"])
	if err != nil {
		fmt.Println(err)
	}
	finalAmount1 := Amount1 + finalAmount
	finalAmount2 := Amount2 - finalAmount

	intFinalAmount1 := strconv.Itoa(finalAmount1)
	intFinalAmount2 := strconv.Itoa(finalAmount2)

	//data := Account + Account2 + Amount
	tr := otel.Tracer("crypto-trace")
	ctx, span := tr.Start(ctx, "handle-gjTransaction")
	span.SetAttributes(attribute.Key("testset").String("value"))
	defer span.End()
	//lakeResponse := acceptTransaction(jsonResponse, ctx)
	////lakeResponse := lakeTransaction(jsonResponse)
	//return lakeResponse

	//hashLedger(data)

	r, err := RunDynamoCreateItem("ledger", Ledger{Account: Account, Amount: intFinalAmount1})
	if err != nil {
		return "--> " + r["msg"], errors.New("--> " + r["msg"])
	}
	r, err = RunDynamoCreateItem("ledger", Ledger{Account: Account2, Amount: intFinalAmount2})
	if err != nil {
		return "--> " + r["msg"], errors.New("--> " + r["msg"])
	}

	return "Transaction Rollback Successful", nil
}
