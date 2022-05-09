package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"gojenga/src/lib/gjLib"
)

const (
	service     = "query"
	environment = "alpha"
	id          = 6
	version     = "1.0.10"
)

func FindUser(Account string, ctx context.Context) (string, error) {
	tr := otel.Tracer("crypto-trace")
	ctx, span := tr.Start(ctx, "findUser")
	span.SetAttributes(attribute.Key("testset").String("value"))
	defer span.End()

	//mongoResult := queryMongo(traffic)
	resultMap, err := gjLib.RunDynamoGetItem(gjLib.Query{TableName: "ledger", Key: "Account", Value: Account})
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

	return string(mapB), errors.New(string(mapB))
}
