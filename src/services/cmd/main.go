package main

import (
	"context"
	"encoding/json"
	"fmt"
	"gjLib"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.uber.org/zap"

	"io"
	"log"
	"net/http"
	"time"

	_ "bytes"
	"go.opentelemetry.io/otel/sdk/trace"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
)

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

const (
	service     = "blockchain"
	environment = "alpha"
	id          = 1
	version     = "1.0.10"
)

type Traffic struct {
	Name               string
	Amount             string
	SourceAccount      string
	DestinationAccount string
	Verb               string
	Role               string
	Port               string
	Payload            string
	Password           string
}

var tp *trace.TracerProvider
var logger *zap.Logger

func testingFunc() (throwError bool) {
	return false
}

func TracerProvider(url string) (*tracesdk.TracerProvider, error) {
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

//the main worker function. called when node gets an http request
func crypto(w http.ResponseWriter, req *http.Request) {

	//tracerProvider gets trace info to Jaeger. One per node but I'm not sure if it matters
	myTp, err := TracerProvider("http://localhost:14268/api/traces")
	if err != nil {
		log.Fatal(err)
	}

	tp = myTp

	otel.SetTracerProvider(tp)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	defer func(ctx context.Context) {
		// Do not make the application hang when it is shutdown.
		ctx, cancel = context.WithTimeout(ctx, time.Second*5)
		defer cancel()
		//if err := tp.Shutdown(ctx); err != nil {
		//	log.Fatal(err)
		//}
	}(ctx)

	//beginning of trace path
	tr := tp.Tracer("crypto-trace")
	_, span := tr.Start(ctx, "receive-request")
	defer span.End()

	//this looks for the http verb and routes accordingly. the start of the second trace
	switch req.Method {
	//outside of normal GET, this is the execute command as well. Might do disco or GET does it
	case "GET":
		tr := otel.Tracer("mempool-trace")
		ctx, span := tr.Start(ctx, "got-get")
		span.SetAttributes(attribute.Key("testset").String("value"))
		defer span.End()
		w.WriteHeader(http.StatusOK)
		results := handleGet(req, ctx)
		_, err := w.Write([]byte(`{"response":` + results + `}`))
		if err != nil {
			logger.Debug(fmt.Sprintf("--> %s", err))
			return
		}
		//POST does create account and runs the discovery command for the master or GET does
	case "POST":
		tr := otel.Tracer("crypto-called")
		ctx, span := tr.Start(ctx, "got-post")
		span.SetAttributes(attribute.Key("testset").String("value"))
		defer span.End()
		w.WriteHeader(http.StatusCreated)
		results := handlePost(req, ctx)
		_, err := w.Write([]byte(`{"response":` + results + `}`))
		if err != nil {
			logger.Debug(fmt.Sprintf("--> %s", err))
			return
		}
		//PUT handles most commands particularly gjTransaction
	case "PUT":
		tr := otel.Tracer("crypto-called")
		ctx, span := tr.Start(ctx, "got-put")
		span.SetAttributes(attribute.Key("testset").String("value"))
		defer span.End()
		w.WriteHeader(http.StatusAccepted)
		results := handlePut(req, ctx)
		_, err := w.Write([]byte(`{"response":` + results + `}`))
		if err != nil {
			logger.Debug(fmt.Sprintf("--> %s", err))
			return
		}
		//gjDelete is self explanatory and untested lately
	case "DELETE":
		tr := otel.Tracer("crypto-called")
		ctx, span := tr.Start(ctx, "got-gjDelete")
		span.SetAttributes(attribute.Key("testset").String("value"))
		defer span.End()
		w.WriteHeader(http.StatusOK)
		results := handleDelete(req, ctx)
		_, err := w.Write([]byte(`{"response":` + results + `}`))
		if err != nil {
			logger.Debug(fmt.Sprintf("--> %s", err))
			return
		}
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func handlePost(req *http.Request, ctx context.Context) (results string) {
	var jsonResponse Traffic

	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal([]byte(body), &jsonResponse)
	if err != nil {
		logger.Debug(fmt.Sprintf("--> %s", err))
		return ""
	}

	if jsonResponse.Verb == "CRT" {
		results, err := gjCreateUser.CreateUser(gjLib.Traffic(jsonResponse), ctx)
		if err != nil {
			log.Println(err)
			return "CRT error"
		}
		return results
	}

	return "POST failed"
}

func handleGet(req *http.Request, ctx context.Context) (results string) {
	var jsonResponse gjLib.Traffic

	tr := otel.Tracer("mempool-trace")
	ctx, span := tr.Start(ctx, "handle-get")
	span.SetAttributes(attribute.Key("testset").String("value"))
	defer span.End()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal([]byte(body), &jsonResponse)
	if err != nil {
		logger.Debug(fmt.Sprintf("--> %s", err))
		return "unmarshal error in handleGet"
	}

	if jsonResponse.Verb == "PING" {
		_, err := gjCreateUser.CreateUser(jsonResponse, ctx)
		if err != nil {
			logger.Debug(fmt.Sprintf("--> %s", err))
			return fmt.Sprintf("PING error: %s", err)
		}
	}

	return "GET failed"
}

func handlePut(req *http.Request, ctx context.Context) (results string) {
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

	if jsonResponse.Verb == "TRAN" {
		results, err := gjTransaction.Transaction(jsonResponse, ctx)
		if err != nil {
			logger.Debug(fmt.Sprintf("--> %s", err))
			return fmt.Sprintf("TRAN error: %s", err)
		}
		return results
	} else if jsonResponse.Verb == "ADD" {
		results, err := gjDeposit.Deposit(jsonResponse, ctx)
		if err != nil {
			logger.Debug(fmt.Sprintf("--> %s", err))
			return fmt.Sprintf("CRT error: %s", err)
		}
		return results
	} else if jsonResponse.Verb == "LOGIN" {
		results, err := gjLogin.Login(jsonResponse, ctx)
		if err != nil {
			logger.Debug(fmt.Sprintf("--> %s", err))
			return fmt.Sprintf("ADD error: %s", err)
		}
		return results
	} else if jsonResponse.Verb == "QUERY" {
		results, err := gjQuery.FindUser(jsonResponse.SourceAccount, ctx)
		if err != nil {
			logger.Debug(fmt.Sprintf("--> %s", err))
			return fmt.Sprintf("QUERY error: %s", err)
		}
		return results
	} else if jsonResponse.Verb == "USER" {
		results, err := gjUser.FindUserAccount(jsonResponse.SourceAccount, ctx)
		if err != nil {
			logger.Debug(fmt.Sprintf("--> %s", err))
			return fmt.Sprintf("USER error: %s", err)
		}
		return results
	} else if jsonResponse.Verb == "DLT" {
		results, err := gjDelete.DeleteUser(jsonResponse, ctx)
		if err != nil {
			logger.Debug(fmt.Sprintf("--> %s", err))
			return fmt.Sprintf("DLT error: %s", err)
		}
		return results
	}

	return "PUT failed"
}

func handleDelete(req *http.Request, ctx context.Context) (results string) {
	var traffic gjLib.Traffic

	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var jsonMap map[string]interface{}
	err = json.Unmarshal([]byte(body), &traffic)
	if err != nil {
		logger.Debug(fmt.Sprintf("--> %s", err))
		return fmt.Sprintf("DLT error: %s", err)
	}

	_, err = gjDelete.DeleteUser(traffic, ctx)
	if err != nil {
		logger.Debug(fmt.Sprintf("--> %s", err))
		return fmt.Sprintf("DLT error: %s", err)
	}

	logger.Debug(fmt.Sprintf("%s", jsonMap["name"]))
	return "gjDelete successful"
}
