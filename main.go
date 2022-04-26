package main

import (
	"context"
)

func main() {
	ctx := context.Background()
	//ctx, cancelCtx := context.WithCancel(ctx)
	startServer("8070", ctx)
	//time.Sleep(time.Second * 2)
	//cancelCtx()
}
