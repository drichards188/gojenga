package main

import (
	"context"
	"gojenga"
)

func main() {

	//type User struct {
	//	Account  string
	//	Password string
	//}

	type Ledger struct {
		Account string
		Amount  string
	}

	//user := User{
	//	Account:  "david",
	//	Password: "54321",
	//}

	//ledger := Ledger{
	//	Account: "david",
	//	Amount:  "200",
	//}

	//query := Query{
	//	TableName: "hashHistory",
	//	Key:       "Iteration",
	//	Value:     "1",
	//}

	ctx := context.Background()

	//ctx, cancelCtx := context.WithCancel(ctx)
	gojenga.StartServer("8070", ctx)
	//time.Sleep(time.Second * 2)
	//cancelCtx()
}
