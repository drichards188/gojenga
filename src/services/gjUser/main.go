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
	service     = "user"
	environment = "alpha"
	id          = 8
	version     = "1.0.10"
)

func FindUserAccount(Account string, ctx context.Context) (string, error) {

	tr := otel.Tracer("crypto-trace")
	ctx, span := tr.Start(ctx, "findUser")
	span.SetAttributes(attribute.Key("testset").String("value"))
	defer span.End()
	//response := lakeFindUser(Account, ctx)
	//fmt.Println("-->data ping results: " + results)
	//return response
	//traffic := Traffic{Account: Account, Role: "USER"}

	//mongoResult := queryMongo(traffic)
	resultMap, err := gjLib.RunDynamoGetItem(gjLib.Query{TableName: "users", Key: "Account", Value: Account})
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
	theAmount := resultMap["Password"]

	mapD := map[string]string{"Account": theAccount, "Password": theAmount}
	mapB, _ := json.Marshal(mapD)

	fmt.Println(string(mapB))

	//c.Write([]byte(mapB))

	fmt.Println(theAccount)

	return string(mapB), errors.New(string(mapB))
}
