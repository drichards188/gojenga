package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

//func testingFunc() (throwError bool) {
//	logger = InitializeLogger()
//	ctx := context.Background()
//
//	traffic := Traffic{SourceAccount: "david", Table: "dynamoTest", Role: "test"}
//
//	resp, err := FindUser(traffic, ctx)
//	if err != nil {
//		logger.Warn(fmt.Sprintf("gjQuery test error: %s", err))
//		return true
//	}
//
//	logger.Debug(fmt.Sprintf("gjQuery test returned: %s", resp))
//
//	return false
//}

func FindUser(jsonResponse Traffic, ctx context.Context) (string, error) {

	Account := jsonResponse.SourceAccount

	tr := otel.Tracer("crypto-trace")
	ctx, span := tr.Start(ctx, "findUser")
	span.SetAttributes(attribute.Key("testset").String("value"))
	defer span.End()

	//mongoResult := queryMongo(traffic)
	resultMap, err := RunDynamoGetItem(Query{TableName: "ledger", Key: "Account", Value: Account})
	if err != nil {
		return "--> " + resultMap["msg"], errors.New("--> " + resultMap["msg"])
	}

	//if mongoResult["message"] == "No Match" {
	//	return "Account Not Found"
	//}

	fmt.Print("Your gjQuery result ")
	//var resultMap primitive.M
	//
	//resultMap = mongoResult.Map()

	msg := resultMap["message"]
	if msg == "No Match" {
		return "Account Not Found", errors.New("Account Not Found")
	}

	theAccount := resultMap["Account"]
	theAmount := resultMap["Amount"]

	mapD := map[string]string{"Account": theAccount, "Amount": theAmount}
	mapB, _ := json.Marshal(mapD)

	fmt.Println(string(mapB))

	//c.Write([]byte(mapB))

	fmt.Println(theAccount)

	return string(mapB), nil
}
