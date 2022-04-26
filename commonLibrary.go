package main

import (
	"context"
	"crypto/sha1"
	"encoding/json"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"log"
	"strconv"
	"time"
)

//works
//hashHistory should have ledger items it is changing, hash of said ledger, hash of prev hash and current hash, time
func transaction(jsonResponse Traffic, ctx context.Context) (results string, err error) {

	source, err := findUser(jsonResponse.SourceAccount, ctx)
	if err != nil {
		log.Println(err)
		return "source findUser failed in transaction", errors.New("source findUser failed in transaction")
	}
	destination, err := findUser(jsonResponse.DestinationAccount, ctx)
	if err != nil {
		log.Println(err)
		return "destination findUser failed in transaction", errors.New("destination findUser failed in transaction")
	}

	if source == "Account Not Found" {
		return "Account1 Not Found", errors.New("account1 not found")
	}

	if destination == "Account Not Found" {
		return "Account2 Not Found", errors.New("account2 not found")
	}

	sourceAccount := jsonResponse.SourceAccount
	transactionAmount := jsonResponse.Amount
	destinationAccount := jsonResponse.DestinationAccount

	traffic := Traffic{SourceAccount: destinationAccount}

	mongoResult, err := queryMongo(jsonResponse)
	if err != nil {
		log.Println(err)
		return "query error in transaction", errors.New("query error in transaction")
	}
	time.Sleep(1 * time.Second)
	resultMap := mongoResult.Map()
	mongoResult2, err := queryMongo(traffic)
	if err != nil {
		log.Println(err)
		return "query error in transaction", errors.New("query error in transaction")
	}
	resultMap2 := mongoResult2.Map()
	finalAmount, err := strconv.Atoi(transactionAmount)
	if err != nil {
		logger.Error(fmt.Sprintf("%s", err))
	}
	SourceAmount, err := strconv.Atoi(resultMap["Amount"].(string))
	if err != nil {
		logger.Error(fmt.Sprintf("%s", err))
	}
	DestinationAmount, err := strconv.Atoi(resultMap2["Amount"].(string))
	if err != nil {
		logger.Error(fmt.Sprintf("%s", err))
	}
	finalAmount1 := SourceAmount - finalAmount
	finalAmount2 := DestinationAmount + finalAmount

	finalSourceAmount := strconv.Itoa(finalAmount1)
	finalDestinationAmount := strconv.Itoa(finalAmount2)

	data := sourceAccount + destinationAccount + transactionAmount
	tr := otel.Tracer("crypto-trace")
	ctx, span := tr.Start(ctx, "handle-transaction")
	span.SetAttributes(attribute.Key("testset").String("value"))
	defer span.End()

	_, err = hashLedger(data)
	if err != nil {
		log.Println(err)
		return "hashLedger fail in transaction", errors.New("hashLedger fail in transaction")
	}

	if resultMap["message"] != "No Match" {
		updateMongo(sourceAccount, finalSourceAmount)
		updateMongo(destinationAccount, finalDestinationAmount)

		return "Transaction Successful", nil
	} else {
		return "Transaction Failed", errors.New("transaction Failed")
	}
}

//works
func findUser(Account string, ctx context.Context) (results string, err error) {

	tr := otel.Tracer("crypto-trace")
	ctx, span := tr.Start(ctx, "findUser")
	span.SetAttributes(attribute.Key("testset").String("value"))
	defer span.End()

	traffic := Traffic{SourceAccount: Account}

	mongoResult, err := queryMongo(traffic)
	if err != nil {
		log.Println(err)
		return "query error in findUser", errors.New("query error in findUser")
	}

	var resultMap primitive.M

	resultMap = mongoResult.Map()

	msg := resultMap["message"]
	if msg == "No Match" {
		return "account Not Found", errors.New("account not found")
	}

	theAccount := resultMap["Account"].(string)
	theAmount := resultMap["Amount"].(string)

	mapD := map[string]string{"Account": theAccount, "Amount": theAmount}
	mapB, _ := json.Marshal(mapD)

	logger.Debug(string(mapB))

	//c.Write([]byte(mapB))

	logger.Debug(theAccount)

	return string(mapB), nil
}

func findUserAccount(Account string, ctx context.Context) (results string, err error) {
	tr := otel.Tracer("crypto-trace")
	ctx, span := tr.Start(ctx, "findUser")
	span.SetAttributes(attribute.Key("testset").String("value"))
	defer span.End()

	traffic := Traffic{SourceAccount: Account, Role: "USER"}

	mongoResult, err := queryMongo(traffic)
	if err != nil {
		log.Println(err)
		return "query error in findAccount", errors.New("query error in findAccount")
	}

	//if mongoResult["message"] == "No Match" {
	//	return "Account Not Found"
	//}

	//fmt.Print("Your query result ")
	var resultMap primitive.M

	resultMap = mongoResult.Map()

	msg := resultMap["message"]
	if msg == "No Match" {
		return "account not found", errors.New("account not found")
	}

	theAccount := resultMap["Account"].(string)
	theAmount := resultMap["Password"].(string)

	mapD := map[string]string{"Account": theAccount, "Password": theAmount}
	mapB, _ := json.Marshal(mapD)

	logger.Debug(string(mapB))

	//c.Write([]byte(mapB))

	//logger.Info(theAccount)

	return string(mapB), nil
}

func deleteUser(jsonResponse Traffic, ctx context.Context) (results string, err error) {
	tr := otel.Tracer("crypto-trace")
	ctx, span := tr.Start(ctx, "deleteUser")
	span.SetAttributes(attribute.Key("testset").String("value"))
	defer span.End()
	deleteMongo(jsonResponse, ctx)
	jsonResponse.Role = "USER"
	deleteMongo(jsonResponse, ctx)
	return "delete user succes", nil
}

//works
func createUser(jsonResponse Traffic, ctx context.Context) (results string, err error) {
	tr := otel.Tracer("crypto-trace")
	_, span := tr.Start(ctx, "createUser")
	span.SetAttributes(attribute.Key("testset").String("value"))
	defer span.End()

	logger.Debug("--> data ping results: " + results)

	traffic := Traffic{SourceAccount: jsonResponse.SourceAccount, Password: jsonResponse.Password}
	mongoResult, err := queryMongo(traffic)
	if err != nil {
		resultMap := mongoResult.Map()
		if resultMap["message"] == "No Match" {
			data := jsonResponse.SourceAccount + jsonResponse.Amount
			_, err = hashLedger(data)
			if err != nil {
				log.Println("hashLedger error in createUser")
				return "hashledger error in createUser", nil
			}
			_, err = writeToMongo("users", jsonResponse.SourceAccount, "", traffic)
			if err != nil {
				log.Println("write error in createUser")
				return "write error in createUser", nil
			}
			_, err = writeToMongo("ledger", jsonResponse.SourceAccount, jsonResponse.Amount, traffic)
			if err != nil {
				log.Println("write error in createUser")
				return "write error in createUser", nil
			}
			return jsonResponse.SourceAccount + " created", nil
		}
	}

	return "account already exists", errors.New("account already exists")
}

func deposit(jsonResponse Traffic, ctx context.Context) (results string, err error) {
	tr := otel.Tracer("crypto-trace")
	_, span := tr.Start(ctx, "createUser")
	span.SetAttributes(attribute.Key("testset").String("value"))
	defer span.End()
	//response := lakeCreateUser(jsonResponse.Account)
	//response := acceptCreateUser(jsonResponse)
	logger.Debug("-->data ping results: " + results)
	//return response

	mongoResult, err := queryMongo(jsonResponse)
	if err != nil {
		log.Println(err)
		return "query error in deposit", errors.New("query error in deposit")
	}
	resultMap := mongoResult.Map()
	if resultMap["message"] != "No Match" {
		data := jsonResponse.SourceAccount + jsonResponse.Amount

		mongoResult, err := queryMongo(jsonResponse)
		if err != nil {
			log.Println(err)
			return "query error in deposit", errors.New("query error in deposit")
		}
		time.Sleep(1 * time.Second)
		resultMap := mongoResult.Map()
		finalAmount, err := strconv.Atoi(jsonResponse.Amount)
		if err != nil {
			logger.Error(fmt.Sprintf("%s", err))
		}
		Amount1, err := strconv.Atoi(resultMap["Amount"].(string))
		if err != nil {
			logger.Error(fmt.Sprintf("%s", err))
		}
		finalAmount1 := Amount1 + finalAmount

		intFinalAmount1 := strconv.Itoa(finalAmount1)

		_, err = hashLedger(data)
		if err != nil {
			log.Println(err)
			return "hashLedger error in deposit", errors.New("hashLedger error in deposit")
		}
		updateMongo(jsonResponse.SourceAccount, intFinalAmount1)
		//writeToMongo("ledger", jsonResponse.Account, jsonResponse.Amount)
		return jsonResponse.SourceAccount + " updated", nil
	}

	return "account not found", errors.New("account not found")
}

func hashLedger(data string) (results string, err error) {
	theHash := ""
	var iteration int
	hashResult, err := queryMongoAll("hashHistory")
	if err != nil {
		log.Println(err)
		return "query error in hashLedger", errors.New("query error in hashLedger")
	}
	if hashResult != nil {
		iteration = len(hashResult)

		mongoHash, err := queryHash(iteration)
		if err != nil {
			log.Println(err)
			return "query error in hashLedger", errors.New("query error in hashLedger")
		}

		//queryMongo(Account)

		//fmt.Print("Your query result ")
		hashMap := mongoHash.Map()
		if _, ok := hashMap["Message"]; ok {
			//do something here

			if hashMap["Message"].(string) == "No Match" {
				theHash = "000000"
			}
		} else {
			theHash = hashMap["Hash"].(string)
			//logger.Info(theHash)
		}
	}

	s := data + theHash

	h := sha1.New()

	h.Write([]byte(s))

	bs := h.Sum(nil)

	currentTime := time.Now()
	hashTime := currentTime.Format("2006-01-02") + currentTime.Format(" 15:04:05")

	//logger.Info(s)
	prettyHash := fmt.Sprintf("%x", bs)
	//logger.Info(prettyHash)
	//logger.Info(hashTime)
	logger.Debug(fmt.Sprintf("--> ledger hashed: %s", prettyHash))

	_, err = writeToHashHistory("hashHistory", prettyHash, hashTime, iteration+1, theHash, data)
	if err != nil {
		log.Println("write error in hashLedger")
		return "write error in hashLedger", nil
	}
	return "hashLedger succesful", nil
}

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
