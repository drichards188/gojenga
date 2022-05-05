package gojenga

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
	"strconv"
	"time"
)

////works
////hashHistory should have ledger items it is changing, hash of said ledger, hash of prev hash and current hash, time
func transaction(jsonResponse Traffic, ctx context.Context) (string, error) {

	user1, err := RunDynamoGetItem(Query{TableName: "ledger", Key: "Account", Value: jsonResponse.SourceAccount})
	if err != nil {
		return "--> " + user1["msg"], errors.New("--> " + user1["msg"])
	}
	user2, err := RunDynamoGetItem(Query{TableName: "ledger", Key: "Account", Value: jsonResponse.DestinationAccount})
	if err != nil {
		return "--> " + user2["msg"], errors.New("--> " + user1["msg"])
	}

	Account := jsonResponse.SourceAccount
	Amount := jsonResponse.Amount
	Account2 := jsonResponse.DestinationAccount

	//traffic := Traffic{Account: Account2}

	time.Sleep(1 * time.Second)

	//mongoResult2 := queryMongo(traffic)
	//resultMap2 := mongoResult2.Map()
	finalAmount, err := strconv.Atoi(Amount)
	if err != nil {
		fmt.Println(err)
	}
	Amount1, err := strconv.Atoi(user1["Amount"])
	if err != nil {
		fmt.Println(err)
	}
	Amount2, err := strconv.Atoi(user2["Amount"])
	if err != nil {
		fmt.Println(err)
	}
	finalAmount1 := Amount1 - finalAmount
	finalAmount2 := Amount2 + finalAmount

	intFinalAmount1 := strconv.Itoa(finalAmount1)
	intFinalAmount2 := strconv.Itoa(finalAmount2)

	//data := Account + Account2 + Amount
	tr := otel.Tracer("crypto-trace")
	ctx, span := tr.Start(ctx, "handle-transaction")
	span.SetAttributes(attribute.Key("testset").String("value"))
	defer span.End()
	//lakeResponse := acceptTransaction(jsonResponse, ctx)
	////lakeResponse := lakeTransaction(jsonResponse)
	//return lakeResponse

	//hashLedger(data)

	r, err := RunDynamoCreateItem("ledger", Ledger{Account: Account, Amount: intFinalAmount1})
	if err != nil {
		return "--> " + r["msg"], errors.New("--> " + r["msg"])
	}
	r, err = RunDynamoCreateItem("ledger", Ledger{Account: Account2, Amount: intFinalAmount2})
	if err != nil {
		return "--> " + r["msg"], errors.New("--> " + r["msg"])
	}

	return "Transaction Successful", errors.New("Transaction Succesful")
}

////works
func findUser(Account string, ctx context.Context) (string, error) {
	tr := otel.Tracer("crypto-trace")
	ctx, span := tr.Start(ctx, "findUser")
	span.SetAttributes(attribute.Key("testset").String("value"))
	defer span.End()
	//response := lakeFindUser(Account, ctx)
	//fmt.Println("-->data ping results: " + results)
	//return response
	//traffic := Traffic{Account: Account}

	//mongoResult := queryMongo(traffic)
	resultMap, err := RunDynamoGetItem(Query{TableName: "ledger", Key: "Account", Value: Account})
	if err != nil {
		return "--> " + resultMap["msg"], errors.New("--> " + resultMap["msg"])
	}

	//if mongoResult["message"] == "No Match" {
	//	return "Account Not Found"
	//}

	fmt.Print("Your query result ")
	//var resultMap primitive.M
	//
	//resultMap = mongoResult.Map()

	msg := resultMap["message"]
	if msg == "No Match" {
		return "Account Not Found", errors.New("Account Not Found")
	}

	theAccount := resultMap["Account"]
	theAmount := resultMap["Amount"]

	mapD := map[string]string{"Account": theAccount, "Amount": theAmount}
	mapB, _ := json.Marshal(mapD)

	fmt.Println(string(mapB))

	//c.Write([]byte(mapB))

	fmt.Println(theAccount)

	return string(mapB), errors.New(string(mapB))
}

func findUserAccount(Account string, ctx context.Context) (string, error) {

	tr := otel.Tracer("crypto-trace")
	ctx, span := tr.Start(ctx, "findUser")
	span.SetAttributes(attribute.Key("testset").String("value"))
	defer span.End()
	//response := lakeFindUser(Account, ctx)
	//fmt.Println("-->data ping results: " + results)
	//return response
	//traffic := Traffic{Account: Account, Role: "USER"}

	//mongoResult := queryMongo(traffic)
	resultMap, err := RunDynamoGetItem(Query{TableName: "users", Key: "Account", Value: Account})
	if err != nil {
		return "--> " + resultMap["msg"], errors.New("--> " + resultMap["msg"])
	}

	//if mongoResult["message"] == "No Match" {
	//	return "Account Not Found"
	//}

	fmt.Print("Your query result ")
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

func deleteUser(jsonResponse Traffic, ctx context.Context) (string, error) {
	tr := otel.Tracer("crypto-trace")
	ctx, span := tr.Start(ctx, "deleteUser")
	span.SetAttributes(attribute.Key("testset").String("value"))
	defer span.End()
	r := RunDynamoDeleteItem("ledger", jsonResponse.SourceAccount)
	if r["code"] == "1" {
		return "--> " + r["msg"], errors.New("--> " + r["msg"])
	}
	RunDynamoDeleteItem("users", jsonResponse.SourceAccount)
	if r["code"] == "1" {
		return "--> " + r["msg"], errors.New("--> " + r["msg"])
	}

	//deleteMongo(jsonResponse, ctx)
	//jsonResponse.Role = "USER"
	//deleteMongo(jsonResponse, ctx)
	return "Delete Succesful", errors.New("Delete Succesful")
}

//works
func createUser(jsonResponse Traffic, ctx context.Context) (string, error) {
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

	r, err := RunDynamoGetItem(Query{TableName: "users", Key: "Account", Value: jsonResponse.SourceAccount})
	if err == nil {
		return "--> User already exists", errors.New("--> User already exists")
	}

	r, err = RunDynamoCreateItem("users", User{Account: jsonResponse.SourceAccount, Password: jsonResponse.Password})
	if err != nil {
		return "--> " + r["msg"], errors.New("--> " + r["msg"])
	}

	r, err = RunDynamoCreateItem("ledger", Ledger{Account: jsonResponse.SourceAccount, Amount: jsonResponse.Amount})
	if err != nil {
		return "--> " + r["msg"], errors.New("--> " + r["msg"])
	}

	return r["msg"], nil
}

func Login(jsonResponse Traffic, ctx context.Context) (results string, err error) {
	tr := otel.Tracer("crypto-trace")
	_, span := tr.Start(ctx, "login")
	span.SetAttributes(attribute.Key("testset").String("value"))
	defer span.End()
	//response := lakeCreateUser(jsonResponse.Account)
	//response := acceptCreateUser(jsonResponse)
	//logger.Debug("-->data ping results: " + results)
	//return response

	jsonResponse.Role = "USER"

	resultMap, err := RunDynamoGetItem(Query{TableName: "users", Key: "Account", Value: jsonResponse.SourceAccount})
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

func deposit(jsonResponse Traffic, ctx context.Context) (string, error) {
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
	resultMap, err := RunDynamoGetItem(Query{TableName: "ledger", Key: "Account", Value: jsonResponse.SourceAccount})
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
	resp, err := RunDynamoCreateItem("ledger", Ledger{Account: jsonResponse.SourceAccount, Amount: intFinalAmount1})
	if err != nil {
		return "--> " + resp["msg"], errors.New("--> " + resp["msg"])
	}
	//updateMongo(jsonResponse.SourceAccount, intFinalAmount1)
	//writeToMongo("ledger", jsonResponse.SourceAccount, jsonResponse.Amount)
	return jsonResponse.SourceAccount + " resp", errors.New(jsonResponse.SourceAccount + " resp")
	//}

	return "Account not found", errors.New("Account not found")
}

//func HashLedger(data string) (results string, err error) {
//	theHash := ""
//	var iteration int
//	hashResult, err := queryMongoAll("hashHistory")
//	if err != nil {
//		log.Println(err)
//		return "query error in hashLedger", errors.New("query error in hashLedger")
//	}
//	if hashResult != nil {
//		iteration = len(hashResult)
//
//		mongoHash, err := queryHash(iteration)
//		if err != nil {
//			log.Println(err)
//			return "query error in hashLedger", errors.New("query error in hashLedger")
//		}
//
//		//queryMongo(Account)
//
//		//fmt.Print("Your query result ")
//		hashMap := mongoHash.Map()
//		if _, ok := hashMap["Message"]; ok {
//			//do something here
//
//			if hashMap["Message"].(string) == "No Match" {
//				theHash = "000000"
//			}
//		} else {
//			theHash = hashMap["Hash"].(string)
//			//logger.Info(theHash)
//		}
//	}
//
//	s := data + theHash
//
//	h := sha1.New()
//
//	h.Write([]byte(s))
//
//	bs := h.Sum(nil)
//
//	currentTime := time.Now()
//	hashTime := currentTime.Format("2006-01-02") + currentTime.Format(" 15:04:05")
//
//	//logger.Info(s)
//	prettyHash := fmt.Sprintf("%x", bs)
//	//logger.Info(prettyHash)
//	//logger.Info(hashTime)
//	logger.Debug(fmt.Sprintf("--> ledger hashed: %s", prettyHash))
//
//	_, err = writeToHashHistory("hashHistory", prettyHash, hashTime, iteration+1, theHash, data)
//	if err != nil {
//		log.Println("write error in hashLedger")
//		return "write error in hashLedger", nil
//	}
//	return "hashLedger succesful", nil
//}

func tracerProvider(url string) (*tracesdk.TracerProvider, error) {
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
