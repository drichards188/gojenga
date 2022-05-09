package main

import (
	"context"
	"errors"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"gojenga/src/lib/gjLib"
)

const (
	service     = "delete"
	environment = "alpha"
	id          = 3
	version     = "1.0.10"
)

func DeleteUser(jsonResponse gjLib.Traffic, ctx context.Context) (string, error) {
	tr := otel.Tracer("crypto-trace")
	ctx, span := tr.Start(ctx, "deleteUser")
	span.SetAttributes(attribute.Key("testset").String("value"))
	defer span.End()
	r := gjLib.RunDynamoDeleteItem("ledger", jsonResponse.SourceAccount)
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
