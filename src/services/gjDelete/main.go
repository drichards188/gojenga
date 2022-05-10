package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/drichards188/gojenga/src/lib/gjLib"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
	"io"
	"log"
	"net/http"
)

var logger *zap.Logger

const (
	service     = "delete"
	environment = "alpha"
	id          = 3
	version     = "1.0.10"
)

func testingFunc() (throwError bool) {
	return false
}

func main() {
	ctx := context.Background()

	config := gjLib.Config{
		Service:     service,
		Environment: environment,
		Id:          id,
		Version:     version,
	}

	//ctx, cancelCtx := context.WithCancel(ctx)
	gjLib.StartServer("8070", config, crypto, ctx)
	//time.Sleep(time.Second * 2)
	//cancelCtx()
}

func crypto(w http.ResponseWriter, req *http.Request) {
	ctx := context.Background()
	w.WriteHeader(http.StatusOK)
	results := handleCrypto(req, ctx)
	_, err := w.Write([]byte(`{"response":` + results + `}`))
	if err != nil {
		logger.Debug(fmt.Sprintf("--> %s", err))
		return
	}
}

func handleCrypto(req *http.Request, ctx context.Context) (results string) {
	var jsonResponse gjLib.Traffic

	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal([]byte(body), &jsonResponse)
	if err != nil {
		logger.Debug(fmt.Sprintf("--> %s", err))
		return fmt.Sprintf("PUT unmarshal error: %s", err)
	}

	if jsonResponse.Verb == "DLT" {
		results, err := DeleteUser(jsonResponse, ctx)
		if err != nil {
			log.Println(err)
			return "DLT error"
		}
		return results
	}

	return "crypto error"
}

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
