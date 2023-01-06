package api

import (
	"context"
	"errors"
	"fmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"math/rand"
	"time"
)

//charset is to generate random passwords and usernames for testing
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

//func testingFunc() (throwError bool) {
//	logger = InitializeLogger()
//	ctx := context.Background()
//	randAccount := GenRanString(6)
//
//	traffic := Traffic{SourceAccount: randAccount, Table: "dynamoTest", Role: "test"}
//
//	resp, err := CreateUser(traffic, ctx)
//	if err != nil {
//		logger.Warn(fmt.Sprintf("gjCreateFunction test error: %s", err))
//		return true
//	}
//
//	logger.Debug(fmt.Sprintf("gjCreateFunction test returned: %s", resp))
//
//	return false
//}

func CreateUser(jsonResponse Traffic, ctx context.Context) (string, error) {
	tr := otel.Tracer("crypto-trace")
	_, span := tr.Start(ctx, "createUser")
	span.SetAttributes(attribute.Key("testset").String("value"))
	defer span.End()

	if jsonResponse.Role == "ROLL" {
		resp, err := RollbackCreateUser(jsonResponse, ctx)

		if err != nil {
			logger.Debug(fmt.Sprintf("--> %s", err))
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			return "", errors.New("CRT error")
		}
		return resp, nil
	}

	if jsonResponse.Test == true {
		r, err := RunDynamoCreateItem(jsonResponse.Table, User{Account: jsonResponse.SourceAccount, Password: jsonResponse.SourceAccount}, ctx)
		if err != nil {
			return "--> " + r.msg, errors.New("--> " + r.msg)
		}
		return r.msg, nil
	}

	jsonResponse.Table = "users"

	r, err := RunDynamoGetItem(Query{TableName: jsonResponse.Table, Key: "Account", Value: jsonResponse.SourceAccount}, ctx)
	if err == nil {
		return "--> User already exists", errors.New("--> User already exists")
	}

	r, err = RunDynamoCreateItem("users", User{Account: jsonResponse.SourceAccount, Password: jsonResponse.Password}, ctx)
	if err != nil {
		return "--> " + r.msg, errors.New("--> " + r.msg)
	}

	r, err = RunDynamoCreateItem("ledger", Ledger{Account: jsonResponse.SourceAccount, Amount: jsonResponse.Amount}, ctx)
	if err != nil {
		return "--> " + r.msg, errors.New("--> " + r.msg)
	}

	return r.msg, nil
}

func RollbackCreateUser(jsonResponse Traffic, ctx context.Context) (string, error) {
	tr := otel.Tracer("crypto-trace")
	_, span := tr.Start(ctx, "rollbackCreateUser")
	span.SetAttributes(attribute.Key("testset").String("value"))
	defer span.End()

	if jsonResponse.Test == true {
		jsonResponse.Table = "usersTest"
		r, err := RunDynamoCreateItem(jsonResponse.Table, User{Account: jsonResponse.SourceAccount, Password: jsonResponse.SourceAccount}, ctx)
		if err != nil {
			return "--> " + r.msg, errors.New("--> " + r.msg)
		}

		r, err = RunDynamoDeleteItem(jsonResponse.Table, jsonResponse.SourceAccount, ctx)
		if err != nil {
			return "--> " + r.msg, errors.New("--> " + r.msg)
		}
		return r.msg, nil
	} else {
		r, err := RunDynamoDeleteItem("users", jsonResponse.SourceAccount, ctx)
		if err != nil {
			return "--> " + r.msg, errors.New("--> " + r.msg)
		}

		r, err = RunDynamoDeleteItem("ledger", jsonResponse.SourceAccount, ctx)
		if err != nil {
			return "--> " + r.msg, errors.New("--> " + r.msg)
		}
	}

	return "rollback createUser complete", nil
}
