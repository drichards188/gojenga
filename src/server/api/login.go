package api

import (
	"context"
	"errors"
	"github.com/drichards188/gojenga/src/lib/gjLib"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

//func testingFunc() (throwError bool) {
//	logger = gjLib.InitializeLogger()
//	ctx := context.Background()
//
//	traffic := gjLib.Traffic{SourceAccount: "david", Password: "54321", Table: "dynamoTest"}
//	resp, err := Login(traffic, ctx)
//	if err != nil {
//		logger.Warn(fmt.Sprintf("gjLogin test error: %s", err))
//		return true
//	}
//
//	logger.Debug(fmt.Sprintf("gjLogin test returned: %s", resp))
//
//	return false
//}

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

	resultMap, err := gjLib.RunDynamoGetItem(gjLib.Query{TableName: jsonResponse.Table, Key: "Account", Value: jsonResponse.SourceAccount})
	if err != nil {
		return "--> User does not exist login fail", errors.New("--> User does not exist login fail")
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
