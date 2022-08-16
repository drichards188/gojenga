package api

import (
	"context"
	"errors"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

//func testingFunc() (throwError bool) {
//	logger = InitializeLogger()
//	ctx := context.Background()
//
//	traffic := Traffic{SourceAccount: "david", Password: "54321", Table: "dynamoTest"}
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

func Login(jsonResponse Traffic, ctx context.Context) (results string, err error) {
	tr := otel.Tracer("crypto-trace")
	ctx, span := tr.Start(ctx, "gjLogin")
	span.SetAttributes(attribute.Key("testset").String("value"))
	defer span.End()
	//response := lakeCreateUser(jsonResponse.Account)
	//response := acceptCreateUser(jsonResponse)
	//logger.Debug("-->data ping results: " + results)
	//return response

	jsonResponse.Role = "USER"

	resultMap, err := RunDynamoGetItem(Query{TableName: jsonResponse.Table, Key: "Account", Value: jsonResponse.SourceAccount}, ctx)
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
