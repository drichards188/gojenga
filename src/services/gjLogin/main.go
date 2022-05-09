package main

import (
	"context"
	"errors"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"gojenga/src/lib/gjLib"
)

const (
	service     = "login"
	environment = "alpha"
	id          = 5
	version     = "1.0.10"
)

func Login(jsonResponse gjLib.Traffic, ctx context.Context) (results string, err error) {
	tr := otel.Tracer("crypto-trace")
	_, span := tr.Start(ctx, "gjLogin")
	span.SetAttributes(attribute.Key("testset").String("value"))
	defer span.End()
	//response := lakeCreateUser(jsonResponse.Account)
	//response := acceptCreateUser(jsonResponse)
	//logger.Debug("-->data ping results: " + results)
	//return response

	jsonResponse.Role = "USER"

	resultMap, err := gjLib.RunDynamoGetItem(gjLib.Query{TableName: "users", Key: "Account", Value: jsonResponse.SourceAccount})
	if err != nil {
		return "--> User already exists", errors.New("--> User already exists")
	}

	if resultMap["code"] != "1" {

		if jsonResponse.SourceAccount == resultMap["Account"] && jsonResponse.Password == resultMap["Password"] {
			rMap := make(map[string]string)

			rMap["token"] = "thisisthetoken"

			return "{\"token\":\"thisisthetoken\"}", nil
		}

		//writeToMongo("ledger", jsonResponse.Account, jsonResponse.Amount)
		return jsonResponse.SourceAccount + " updated", nil
	}

	return "account not found", errors.New("account not found")
}
