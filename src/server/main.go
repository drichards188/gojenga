package main

import (
	"context"
	"fmt"
	"github.com/drichards188/gojenga/src/lib/gjLib"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.uber.org/zap"
	"log"
	"net/http"
	"server/api"
	"time"

	_ "bytes"
	"go.opentelemetry.io/otel/sdk/trace"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
)

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

const (
	service     = "gojenga"
	environment = "alpha"
	id          = 1
	version     = "1.0.10"
)

var tp *trace.TracerProvider
var logger *zap.Logger

func testingFunc() (throwError bool) {
	logger = gjLib.InitializeLogger()
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
		span.SetAttributes(attribute.Key("my-version").String("1,0,1"))
		defer span.End()
		w.WriteHeader(http.StatusOK)
		results := api.HandleGet(req, ctx)
		_, err := w.Write([]byte(`{"response":` + results + `}`))
		if err != nil {
			logger.Debug(fmt.Sprintf("--> %s", err))
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			return
		}
		//POST does create account and runs the discovery command for the master or GET does
	case "POST":
		tr := otel.Tracer("crypto-called")
		ctx, span := tr.Start(ctx, "got-post")
		span.SetAttributes(attribute.Key("my-version").String("1,0,1"))
		defer span.End()
		w.WriteHeader(http.StatusCreated)

		results, err := api.HandlePost(req, ctx)
		if err != nil {
			logger.Debug(fmt.Sprintf("--> %s", err))
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			return
		}

		_, err = w.Write([]byte(`{"response":` + results + `}`))
		if err != nil {
			logger.Debug(fmt.Sprintf("--> %s", err))
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}

		if err != nil {
			logger.Debug(fmt.Sprintf("--> %s", err))
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			return
		}
		//PUT handles most commands particularly gjTransaction
	case "PUT":
		tr := otel.Tracer("crypto-called")
		ctx, span := tr.Start(ctx, "got-put")
		span.SetAttributes(attribute.Key("my-version").String("1,0,1"))
		defer span.End()
		w.WriteHeader(http.StatusAccepted)
		results := api.HandlePut(req, ctx)
		_, err := w.Write([]byte(`{"response":` + results + `}`))
		if err != nil {
			logger.Debug(fmt.Sprintf("--> %s", err))
			return
		}
		//gjDelete is self explanatory and untested lately
	case "DELETE":
		tr := otel.Tracer("crypto-called")
		ctx, span := tr.Start(ctx, "got-gjDelete")
		span.SetAttributes(attribute.Key("my-version").String("1,0,1"))
		defer span.End()
		w.WriteHeader(http.StatusOK)
		results := api.HandleDelete(req, ctx)
		_, err := w.Write([]byte(`{"response":` + results + `}`))
		if err != nil {
			logger.Debug(fmt.Sprintf("--> %s", err))
			return
		}
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}
