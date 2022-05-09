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
	"strconv"
	"time"
)

var logger *zap.Logger

const (
	service     = "deposit"
	environment = "alpha"
	id          = 4
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

	if jsonResponse.Verb == "ADD" {
		results, err := Deposit(jsonResponse, ctx)
		if err != nil {
			log.Println(err)
			return "ADD error"
		}
		return results
	}

	return "crypto error"
}

func Deposit(jsonResponse gjLib.Traffic, ctx context.Context) (string, error) {
	var results string
	tr := otel.Tracer("crypto-trace")
	_, span := tr.Start(ctx, "createUser")
	span.SetAttributes(attribute.Key("testset").String("value"))
	defer span.End()
	//response := lakeCreateUser(jsonResponse.SourceAccount)
	//response := acceptCreateUser(jsonResponse)
	fmt.Println("-->data ping results: " + results)
	//return response

	//mongoResult := queryMongo(jsonResponse)
	resultMap, err := gjLib.RunDynamoGetItem(gjLib.Query{TableName: "ledger", Key: "Account", Value: jsonResponse.SourceAccount})
	if err != nil {
		return "--> " + resultMap["msg"], errors.New("--> " + resultMap["msg"])
	}

	//if resultMap["message"] != "No Match" {
	//data := jsonResponse.SourceAccount + jsonResponse.Amount

	//mongoResult := queryMongo(jsonResponse)
	time.Sleep(1 * time.Second)
	//resultMap := mongoResult.Map()
	finalAmount, err := strconv.Atoi(jsonResponse.Amount)
	if err != nil {
		fmt.Println(err)
	}
	Amount1, err := strconv.Atoi(resultMap["Amount"])
	if err != nil {
		fmt.Println(err)
	}
	finalAmount1 := Amount1 + finalAmount

	intFinalAmount1 := strconv.Itoa(finalAmount1)

	//hashLedger(data)
	resp, err := gjLib.RunDynamoCreateItem("ledger", gjLib.Ledger{Account: jsonResponse.SourceAccount, Amount: intFinalAmount1})
	if err != nil {
		return "--> " + resp["msg"], errors.New("--> " + resp["msg"])
	}
	//updateMongo(jsonResponse.SourceAccount, intFinalAmount1)
	//writeToMongo("ledger", jsonResponse.SourceAccount, jsonResponse.Amount)
	return jsonResponse.SourceAccount + " resp", errors.New(jsonResponse.SourceAccount + " resp")
	//}

	return "Account not found", errors.New("Account not found")
}
