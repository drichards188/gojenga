package gjCreateUser

import (
	"context"
	"errors"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"gojenga/lib/gjLib"
)

const (
	service     = "blockchain"
	environment = "alpha"
	id          = 1
	verion      = "1.0.10"
)

func TracerProvider(url string) (*tracesdk.TracerProvider, error) {
	// Create the Jaeger exporter
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}
	tp := tracesdk.NewTracerProvider(
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in a Resource.
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(service),
			attribute.String("environment", environment),
			attribute.Int64("ID", id),
		)),
	)
	return tp, nil
}

func CreateUser(jsonResponse gjLib.Traffic, ctx context.Context) (string, error) {
	tr := otel.Tracer("crypto-trace")
	_, span := tr.Start(ctx, "createUser")
	span.SetAttributes(attribute.Key("testset").String("value"))
	defer span.End()
	//response := lakeCreateUser(jsonResponse.SourceAccount)
	//response := acceptCreateUser(jsonResponse)
	//fmt.Println("-->data ping results: " + results)
	//return response

	//traffic := Traffic{Account: jsonResponse.SourceAccount, Password: jsonResponse.Password}
	//resultMap := RunDynamoGetItem(Query{TableName: "users", Key: "Account", Value: jsonResponse.SourceAccount})
	//if resultMap["message"] == "No Match" {
	//	data := jsonResponse.SourceAccount + jsonResponse.Amount
	//	hashLedger(data)
	//	writeToMongo("users", jsonResponse.SourceAccount, "", traffic)
	//	writeToMongo("ledger", jsonResponse.SourceAccount, jsonResponse.Amount, traffic)
	//	return jsonResponse.SourceAccount + " created"
	//}

	r, err := gjLib.RunDynamoGetItem(gjLib.Query{TableName: "users", Key: "Account", Value: jsonResponse.SourceAccount})
	if err == nil {
		return "--> User already exists", errors.New("--> User already exists")
	}

	r, err = gjLib.RunDynamoCreateItem("users", gjLib.User{Account: jsonResponse.SourceAccount, Password: jsonResponse.Password})
	if err != nil {
		return "--> " + r["msg"], errors.New("--> " + r["msg"])
	}

	r, err = gjLib.RunDynamoCreateItem("ledger", gjLib.Ledger{Account: jsonResponse.SourceAccount, Amount: jsonResponse.Amount})
	if err != nil {
		return "--> " + r["msg"], errors.New("--> " + r["msg"])
	}

	return r["msg"], nil
}
