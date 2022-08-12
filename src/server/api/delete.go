package api

import (
	"context"
	"errors"
	"fmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

//func testingFunc() (throwError bool) {
//	logger = InitializeLogger()
//	return false
//}

func DeleteUser(jsonResponse Traffic, ctx context.Context) (string, error) {
	tr := otel.Tracer("crypto-trace")
	ctx, span := tr.Start(ctx, "deleteUser")
	span.SetAttributes(attribute.Key("testset").String("value"))
	defer span.End()
	r, err := RunDynamoDeleteItem("ledger", jsonResponse.SourceAccount, ctx)
	if err != nil {
		fmt.Println("error in DeleteUser")
	}
	if r["code"] == "1" {
		return "--> " + r["msg"], errors.New("--> " + r["msg"])
	}
	RunDynamoDeleteItem("users", jsonResponse.SourceAccount, ctx)
	if r["code"] == "1" {
		return "--> " + r["msg"], errors.New("--> " + r["msg"])
	}

	//deleteMongo(jsonResponse, ctx)
	//jsonResponse.Role = "USER"
	//deleteMongo(jsonResponse, ctx)
	return "Delete Succesful", errors.New("Delete Succesful")
}
