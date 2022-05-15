package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/drichards188/gojenga/src/lib/gjLib"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/trace"
	"go.uber.org/zap"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var tp *trace.TracerProvider
var logger *zap.Logger

const (
	service     = "createUser"
	environment = "alpha"
	id          = 2
	version     = "1.0.11"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func GenRanString(length int) string {
	return StringWithCharset(length, charset)
}

func testingFunc() (throwError bool) {
	logger = gjLib.InitializeLogger()
	ctx := context.Background()
	randAccount := GenRanString(6)

	traffic := gjLib.Traffic{SourceAccount: randAccount, Table: "dynamoTest", Role: "test"}

	resp, err := CreateUser(traffic, ctx)
	if err != nil {
		logger.Warn(fmt.Sprintf("gjCreateFunction test error: %s", err))
		return true
	}

	logger.Debug(fmt.Sprintf("gjCreateFunction test returned: %s", resp))

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

	if jsonResponse.Verb == "CRT" {
		results, err := CreateUser(jsonResponse, ctx)
		if err != nil {
			log.Println(err)
			return "CRT error"
		}
		return results
	}

	return "crypto error"
}

func CreateUser(jsonResponse gjLib.Traffic, ctx context.Context) (string, error) {
	tr := otel.Tracer("crypto-trace")
	_, span := tr.Start(ctx, "createUser")
	span.SetAttributes(attribute.Key("testset").String("value"))
	defer span.End()

	if jsonResponse.Role == "test" {
		r, err := gjLib.RunDynamoCreateItem(jsonResponse.Table, gjLib.User{Account: jsonResponse.SourceAccount, Password: jsonResponse.SourceAccount})
		if err != nil {
			return "--> " + r["msg"], errors.New("--> " + r["msg"])
		}
		return r["msg"], nil
	}

	r, err := gjLib.RunDynamoGetItem(gjLib.Query{TableName: jsonResponse.Table, Key: "Account", Value: jsonResponse.SourceAccount})
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
