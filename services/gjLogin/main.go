package gjLogin

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
