package gjUser

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
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
