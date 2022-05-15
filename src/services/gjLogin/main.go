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
	service     = "login"
	environment = "alpha"
	id          = 5
	version     = "1.0.10"
)

func testingFunc() (throwError bool) {
	logger = gjLib.InitializeLogger()
	ctx := context.Background()

	traffic := gjLib.Traffic{SourceAccount: "david", Password: "54321", Table: "dynamoTest"}
	resp, err := Login(traffic, ctx)
	if err != nil {
		logger.Warn(fmt.Sprintf("gjLogin test error: %s", err))
		return true
	}

	logger.Debug(fmt.Sprintf("gjLogin test returned: %s", resp))

	return false
}

func main() {
	logger = gjLib.InitializeLogger()

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

	if jsonResponse.Verb == "LOGIN" {
		results, err := Login(jsonResponse, ctx)
		if err != nil {
			log.Println(err)
			return "LOGIN error"
		}
		return results
	}

	return "crypto error"
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
