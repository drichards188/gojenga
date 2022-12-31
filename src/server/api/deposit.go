package api

import (
	"context"
	"errors"
	"fmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"strconv"
	"time"
)

//func testingFunc() (throwError bool) {
//	logger = InitializeLogger()
//	ctx := context.Background()
//
//	traffic := Traffic{SourceAccount: "david", Table: "dynamoTest", Role: "test", Amount: "12"}
//
//	resp, err := Deposit(traffic, ctx)
//	if err != nil {
//		logger.Warn(fmt.Sprintf("gjDeposit test error: %s", err))
//		return true
//	}
//
//	logger.Debug(fmt.Sprintf("gjDeposit test returned: %s", resp))
//
//	return false
//}

func Deposit(jsonResponse Traffic, ctx context.Context) (string, error) {
	var results string
	tr := otel.Tracer("crypto-trace")
	ctx, span := tr.Start(ctx, "createUser")
	span.SetAttributes(attribute.Key("testset").String("value"))
	defer span.End()

	//response := lakeCreateUser(jsonResponse.SourceAccount)
	//response := acceptCreateUser(jsonResponse)
	fmt.Println("-->data ping results: " + results)
	//return response

	//mongoResult := queryMongo(jsonResponse)
	resultMap, err := RunDynamoGetItem(Query{TableName: "ledger", Key: "Account", Value: jsonResponse.SourceAccount}, ctx)
	if err != nil {
		logger.Debug(fmt.Sprintf("--> %s", err))
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return "--> " + resultMap.msg, errors.New("--> " + resultMap.msg)

	}

	//if resultMap["message"] != "No Match" {
	//data := jsonResponse.SourceAccount + jsonResponse.Amount

	//mongoResult := queryMongo(jsonResponse)
	time.Sleep(1 * time.Second)
	//resultMap := mongoResult.Map()
	finalAmount, err := strconv.Atoi(jsonResponse.Amount)
	if err != nil {
		fmt.Println(err)
	}
	Amount1, err := strconv.Atoi(resultMap.data["Amount"])
	if err != nil {
		fmt.Println(err)
	}
	finalAmount1 := Amount1 + finalAmount

	intFinalAmount1 := strconv.Itoa(finalAmount1)

	//hashLedger(data)
	resp, err := RunDynamoCreateItem("ledger", Ledger{Account: jsonResponse.SourceAccount, Amount: intFinalAmount1}, ctx)
	if err != nil {
		return "--> " + resp.msg, errors.New("--> " + resp.msg)
	}
	//updateMongo(jsonResponse.SourceAccount, intFinalAmount1)
	//writeToMongo("ledger", jsonResponse.SourceAccount, jsonResponse.Amount)
	return jsonResponse.SourceAccount + " resp", nil
	//}

	return "Account not found", errors.New("Account not found")
}

func RollbackDeposit() {

}
