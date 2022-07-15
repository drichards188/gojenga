package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/drichards188/gojenga/src/lib/gjLib"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"io"
	"log"
	"net/http"
)

func testingFunc() (throwError bool) {
	logger = gjLib.InitializeLogger()
	ctx := context.Background()

	traffic := gjLib.Traffic{SourceAccount: "david", Table: "dynamoTest", Role: "test"}

	resp, err := FindUserAccount(traffic, ctx)
	if err != nil {
		logger.Warn(fmt.Sprintf("gjDeposit test error: %s", err))
		return true
	}

	logger.Debug(fmt.Sprintf("gjDeposit test returned: %s", resp))

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

	if jsonResponse.Verb == "USER" {
		results, err := FindUserAccount(jsonResponse, ctx)
		if err != nil {
			log.Println(err)
			return "CRT error"
		}
		return results
	}

	return "crypto error"
}

func FindUserAccount(jsonResponse gjLib.Traffic, ctx context.Context) (string, error) {
	Account := jsonResponse.SourceAccount
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

	return string(mapB), nil
}
