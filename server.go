package gojenga

import (
	_ "bytes"
	"context"
	"encoding/json"
	"fmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	//"go.opentelemetry.io/otel/exporters/jaeger"
	//"go.opentelemetry.io/otel/sdk/resource"
	"io"
	"log"
	"net/http"
	"time"
	//"go.opentelemetry.io/otel"
	//"go.opentelemetry.io/otel/attribute"
	//tracesdk "go.opentelemetry.io/otel/sdk/trace"
	//semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	//_ "go.opentelemetry.io/otel/trace"
)

const (
	service     = "blockchain"
	environment = "alpha"
	id          = 1
	verion      = "1.0.9"
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

func StartServer(port string, ctx context.Context) {

	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, _ = config.Build()
	defer logger.Sync()

	logger.Debug(fmt.Sprintf("version: %s", verion))

	port = ":" + port

	http.HandleFunc("/crypto", crypto)

	logger.Debug(fmt.Sprintf("--> Listening on %s", port))

	log.Fatal(http.ListenAndServe(port, nil))
}

//the main worker function. called when node gets an http request
func crypto(w http.ResponseWriter, req *http.Request) {

	//tracerProvider gets trace info to Jaeger. One per node but I'm not sure if it matters
	myTp, err := tracerProvider("http://localhost:14268/api/traces")
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
		//PUT handles most commands particularly transaction
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
		//delete is self explanatory and untested lately
	case "DELETE":
		tr := otel.Tracer("crypto-called")
		ctx, span := tr.Start(ctx, "got-delete")
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
		results, err := createUser(jsonResponse, ctx)
		if err != nil {
			log.Println(err)
			return "CRT error"
		}
		return results
	}

	return "POST failed"
}

func handleGet(req *http.Request, ctx context.Context) (results string) {
	var jsonResponse Traffic

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
		_, err := createUser(jsonResponse, ctx)
		if err != nil {
			logger.Debug(fmt.Sprintf("--> %s", err))
			return fmt.Sprintf("PING error: %s", err)
		}
	}

	return "GET failed"
}

func handlePut(req *http.Request, ctx context.Context) (results string) {
	var jsonResponse Traffic

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
		results, err := transaction(jsonResponse, ctx)
		if err != nil {
			logger.Debug(fmt.Sprintf("--> %s", err))
			return fmt.Sprintf("TRAN error: %s", err)
		}
		return results
	} else if jsonResponse.Verb == "ADD" {
		results, err := deposit(jsonResponse, ctx)
		if err != nil {
			logger.Debug(fmt.Sprintf("--> %s", err))
			return fmt.Sprintf("CRT error: %s", err)
		}
		return results
	} else if jsonResponse.Verb == "LOGIN" {
		results, err := Login(jsonResponse, ctx)
		if err != nil {
			logger.Debug(fmt.Sprintf("--> %s", err))
			return fmt.Sprintf("ADD error: %s", err)
		}
		return results
	} else if jsonResponse.Verb == "QUERY" {
		results, err := findUser(jsonResponse.SourceAccount, ctx)
		if err != nil {
			logger.Debug(fmt.Sprintf("--> %s", err))
			return fmt.Sprintf("QUERY error: %s", err)
		}
		return results
	} else if jsonResponse.Verb == "USER" {
		results, err := findUserAccount(jsonResponse.SourceAccount, ctx)
		if err != nil {
			logger.Debug(fmt.Sprintf("--> %s", err))
			return fmt.Sprintf("USER error: %s", err)
		}
		return results
	} else if jsonResponse.Verb == "DLT" {
		results, err := deleteUser(jsonResponse, ctx)
		if err != nil {
			logger.Debug(fmt.Sprintf("--> %s", err))
			return fmt.Sprintf("DLT error: %s", err)
		}
		return results
	}

	return "PUT failed"
}

func handleDelete(req *http.Request, ctx context.Context) (results string) {
	var traffic Traffic

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

	_, err = deleteUser(traffic, ctx)
	if err != nil {
		logger.Debug(fmt.Sprintf("--> %s", err))
		return fmt.Sprintf("DLT error: %s", err)
	}

	logger.Debug(fmt.Sprintf("%s", jsonMap["name"]))
	return "delete successful"
}
