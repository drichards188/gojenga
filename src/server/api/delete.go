package api

import (
	"context"
	"errors"
	"fmt"
	"github.com/drichards188/gojenga/src/lib/gjLib"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

//func testingFunc() (throwError bool) {
//	logger = gjLib.InitializeLogger()
//	return false
//}

func DeleteUser(jsonResponse gjLib.Traffic, ctx context.Context) (string, error) {
	tr := otel.Tracer("crypto-trace")
	ctx, span := tr.Start(ctx, "deleteUser")
	span.SetAttributes(attribute.Key("testset").String("value"))
	defer span.End()
	r, err := gjLib.RunDynamoDeleteItem("ledger", jsonResponse.SourceAccount)
	if err != nil {
		fmt.Println("error in DeleteUser")
	}
	if r["code"] == "1" {
		return "--> " + r["msg"], errors.New("--> " + r["msg"])
	}
	gjLib.RunDynamoDeleteItem("users", jsonResponse.SourceAccount)
	if r["code"] == "1" {
		return "--> " + r["msg"], errors.New("--> " + r["msg"])
	}

	//deleteMongo(jsonResponse, ctx)
	//jsonResponse.Role = "USER"
	//deleteMongo(jsonResponse, ctx)
	return "Delete Succesful", errors.New("Delete Succesful")
}
