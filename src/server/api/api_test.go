package api

import (
	"context"
	"github.com/drichards188/gojenga/src/lib/gjLib"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	"go.uber.org/zap"
	"net/http"
	"reflect"
	"testing"
)

func TestCreateUser(t *testing.T) {
	type args struct {
		jsonResponse gjLib.Traffic
		ctx          context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateUser(tt.args.jsonResponse, tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CreateUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteUser(t *testing.T) {
	type args struct {
		jsonResponse gjLib.Traffic
		ctx          context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeleteUser(tt.args.jsonResponse, tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DeleteUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeposit(t *testing.T) {
	type args struct {
		jsonResponse gjLib.Traffic
		ctx          context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Deposit(tt.args.jsonResponse, tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Deposit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Deposit() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindUser(t *testing.T) {
	type args struct {
		jsonResponse gjLib.Traffic
		ctx          context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindUser(tt.args.jsonResponse, tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FindUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindUserAccount(t *testing.T) {
	type args struct {
		jsonResponse gjLib.Traffic
		ctx          context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindUserAccount(tt.args.jsonResponse, tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindUserAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FindUserAccount() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenRanString(t *testing.T) {
	type args struct {
		length int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenRanString(tt.args.length); got != tt.want {
				t.Errorf("GenRanString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInitializeLogger(t *testing.T) {
	tests := []struct {
		name string
		want *zap.Logger
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitializeLogger(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitializeLogger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogin(t *testing.T) {
	type args struct {
		jsonResponse gjLib.Traffic
		ctx          context.Context
	}
	tests := []struct {
		name        string
		args        args
		wantResults string
		wantErr     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResults, err := Login(tt.args.jsonResponse, tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResults != tt.wantResults {
				t.Errorf("Login() gotResults = %v, want %v", gotResults, tt.wantResults)
			}
		})
	}
}

func TestRollbackCreateUser(t *testing.T) {
	type args struct {
		jsonResponse gjLib.Traffic
		ctx          context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RollbackCreateUser(tt.args.jsonResponse, tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("RollbackCreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RollbackCreateUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRollbackDeposit(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RollbackDeposit()
		})
	}
}

func TestRunDynamoCreateItem(t *testing.T) {
	type args struct {
		tableName string
		item      T
	}
	tests := []struct {
		name     string
		args     args
		wantResp map[string]string
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := RunDynamoCreateItem(tt.args.tableName, tt.args.item)
			if (err != nil) != tt.wantErr {
				t.Errorf("RunDynamoCreateItem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("RunDynamoCreateItem() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestRunDynamoCreateTable(t *testing.T) {
	type args struct {
		tableName string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RunDynamoCreateTable(tt.args.tableName)
		})
	}
}

func TestRunDynamoDeleteItem(t *testing.T) {
	type args struct {
		tableName string
		value     string
	}
	tests := []struct {
		name     string
		args     args
		wantResp map[string]string
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := RunDynamoDeleteItem(tt.args.tableName, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("RunDynamoDeleteItem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("RunDynamoDeleteItem() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestRunDynamoGetItem(t *testing.T) {
	type args struct {
		query Query
	}
	tests := []struct {
		name     string
		args     args
		wantResp map[string]string
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := RunDynamoGetItem(tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("RunDynamoGetItem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("RunDynamoGetItem() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestStartServer(t *testing.T) {
	type args struct {
		port   string
		config Config
		crypto func(w http.ResponseWriter, req *http.Request)
		ctx    context.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			StartServer(tt.args.port, tt.args.config, tt.args.crypto, tt.args.ctx)
		})
	}
}

func TestStringWithCharset(t *testing.T) {
	type args struct {
		length  int
		charset string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringWithCharset(tt.args.length, tt.args.charset); got != tt.want {
				t.Errorf("StringWithCharset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTracerProvider(t *testing.T) {
	type args struct {
		url    string
		config Config
	}
	tests := []struct {
		name    string
		args    args
		want    *tracesdk.TracerProvider
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TracerProvider(tt.args.url, tt.args.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("TracerProvider() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TracerProvider() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransaction(t *testing.T) {
	type args struct {
		jsonResponse gjLib.Traffic
		ctx          context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Transaction(tt.args.jsonResponse, tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Transaction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Transaction() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransactionRollback(t *testing.T) {
	type args struct {
		jsonResponse gjLib.Traffic
		ctx          context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TransactionRollback(tt.args.jsonResponse, tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("TransactionRollback() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("TransactionRollback() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_crypto(t *testing.T) {
	type args struct {
		w   http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			crypto(tt.args.w, tt.args.req)
		})
	}
}

func Test_crypto1(t *testing.T) {
	type args struct {
		w   http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			crypto(tt.args.w, tt.args.req)
		})
	}
}

func Test_crypto2(t *testing.T) {
	type args struct {
		w   http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			crypto(tt.args.w, tt.args.req)
		})
	}
}

func Test_crypto3(t *testing.T) {
	type args struct {
		w   http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			crypto(tt.args.w, tt.args.req)
		})
	}
}

func Test_crypto4(t *testing.T) {
	type args struct {
		w   http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			crypto(tt.args.w, tt.args.req)
		})
	}
}

func Test_crypto5(t *testing.T) {
	type args struct {
		w   http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			crypto(tt.args.w, tt.args.req)
		})
	}
}

func Test_crypto6(t *testing.T) {
	type args struct {
		w   http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			crypto(tt.args.w, tt.args.req)
		})
	}
}

func Test_handleCrypto(t *testing.T) {
	type args struct {
		req *http.Request
		ctx context.Context
	}
	tests := []struct {
		name        string
		args        args
		wantResults string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResults := handleCrypto(tt.args.req, tt.args.ctx); gotResults != tt.wantResults {
				t.Errorf("handleCrypto() = %v, want %v", gotResults, tt.wantResults)
			}
		})
	}
}

func Test_handleCrypto1(t *testing.T) {
	type args struct {
		req *http.Request
		ctx context.Context
	}
	tests := []struct {
		name        string
		args        args
		wantResults string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResults := handleCrypto(tt.args.req, tt.args.ctx); gotResults != tt.wantResults {
				t.Errorf("handleCrypto() = %v, want %v", gotResults, tt.wantResults)
			}
		})
	}
}

func Test_handleCrypto2(t *testing.T) {
	type args struct {
		req *http.Request
		ctx context.Context
	}
	tests := []struct {
		name        string
		args        args
		wantResults string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResults := handleCrypto(tt.args.req, tt.args.ctx); gotResults != tt.wantResults {
				t.Errorf("handleCrypto() = %v, want %v", gotResults, tt.wantResults)
			}
		})
	}
}

func Test_handleCrypto3(t *testing.T) {
	type args struct {
		req *http.Request
		ctx context.Context
	}
	tests := []struct {
		name        string
		args        args
		wantResults string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResults := handleCrypto(tt.args.req, tt.args.ctx); gotResults != tt.wantResults {
				t.Errorf("handleCrypto() = %v, want %v", gotResults, tt.wantResults)
			}
		})
	}
}

func Test_handleCrypto4(t *testing.T) {
	type args struct {
		req *http.Request
		ctx context.Context
	}
	tests := []struct {
		name        string
		args        args
		wantResults string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResults := handleCrypto(tt.args.req, tt.args.ctx); gotResults != tt.wantResults {
				t.Errorf("handleCrypto() = %v, want %v", gotResults, tt.wantResults)
			}
		})
	}
}

func Test_handleCrypto5(t *testing.T) {
	type args struct {
		req *http.Request
		ctx context.Context
	}
	tests := []struct {
		name        string
		args        args
		wantResults string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResults := handleCrypto(tt.args.req, tt.args.ctx); gotResults != tt.wantResults {
				t.Errorf("handleCrypto() = %v, want %v", gotResults, tt.wantResults)
			}
		})
	}
}

func Test_handleCrypto6(t *testing.T) {
	type args struct {
		req *http.Request
		ctx context.Context
	}
	tests := []struct {
		name        string
		args        args
		wantResults string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResults := handleCrypto(tt.args.req, tt.args.ctx); gotResults != tt.wantResults {
				t.Errorf("handleCrypto() = %v, want %v", gotResults, tt.wantResults)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_main1(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_main2(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_main3(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_main4(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_main5(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_main6(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_main7(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_testingFunc(t *testing.T) {
	tests := []struct {
		name           string
		wantThrowError bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotThrowError := testingFunc(); gotThrowError != tt.wantThrowError {
				t.Errorf("testingFunc() = %v, want %v", gotThrowError, tt.wantThrowError)
			}
		})
	}
}

func Test_testingFunc1(t *testing.T) {
	tests := []struct {
		name           string
		wantThrowError bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotThrowError := testingFunc(); gotThrowError != tt.wantThrowError {
				t.Errorf("testingFunc() = %v, want %v", gotThrowError, tt.wantThrowError)
			}
		})
	}
}

func Test_testingFunc2(t *testing.T) {
	tests := []struct {
		name           string
		wantThrowError bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotThrowError := testingFunc(); gotThrowError != tt.wantThrowError {
				t.Errorf("testingFunc() = %v, want %v", gotThrowError, tt.wantThrowError)
			}
		})
	}
}

func Test_testingFunc3(t *testing.T) {
	tests := []struct {
		name           string
		wantThrowError bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotThrowError := testingFunc(); gotThrowError != tt.wantThrowError {
				t.Errorf("testingFunc() = %v, want %v", gotThrowError, tt.wantThrowError)
			}
		})
	}
}

func Test_testingFunc4(t *testing.T) {
	tests := []struct {
		name           string
		wantThrowError bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotThrowError := testingFunc(); gotThrowError != tt.wantThrowError {
				t.Errorf("testingFunc() = %v, want %v", gotThrowError, tt.wantThrowError)
			}
		})
	}
}

func Test_testingFunc5(t *testing.T) {
	tests := []struct {
		name           string
		wantThrowError bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotThrowError := testingFunc(); gotThrowError != tt.wantThrowError {
				t.Errorf("testingFunc() = %v, want %v", gotThrowError, tt.wantThrowError)
			}
		})
	}
}

func Test_testingFunc6(t *testing.T) {
	tests := []struct {
		name           string
		wantThrowError bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotThrowError := testingFunc(); gotThrowError != tt.wantThrowError {
				t.Errorf("testingFunc() = %v, want %v", gotThrowError, tt.wantThrowError)
			}
		})
	}
}
