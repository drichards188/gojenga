package api

import (
	"context"
	"errors"
	"github.com/drichards188/gojenga/src/lib/gjLib"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
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
//	logger = gjLib.InitializeLogger()
//	ctx := context.Background()
//	randAccount := GenRanString(6)
//
//	traffic := gjLib.Traffic{SourceAccount: randAccount, Table: "dynamoTest", Role: "test"}
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

func RollbackCreateUser(jsonResponse gjLib.Traffic, ctx context.Context) (string, error) {
	tr := otel.Tracer("crypto-trace")
	_, span := tr.Start(ctx, "rollbackCreateUser")
	span.SetAttributes(attribute.Key("testset").String("value"))
	defer span.End()

	if jsonResponse.Role == "test" {
		r, err := gjLib.RunDynamoCreateItem(jsonResponse.Table, gjLib.User{Account: jsonResponse.SourceAccount, Password: jsonResponse.SourceAccount})
		if err != nil {
			return "--> " + r["msg"], errors.New("--> " + r["msg"])
		}
		return r["msg"], nil
	} else {
		r, err := gjLib.RunDynamoDeleteItem("users", jsonResponse.SourceAccount)
		if err != nil {
			return "--> " + r["msg"], errors.New("--> " + r["msg"])
		}

		r, err = gjLib.RunDynamoDeleteItem("ledger", jsonResponse.SourceAccount)
		if err != nil {
			return "--> " + r["msg"], errors.New("--> " + r["msg"])
		}
	}

	return "rollback createUser complete", nil
}
