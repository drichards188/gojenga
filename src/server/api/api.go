package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/drichards188/gojenga/src/lib/gjLib"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"io"
	"log"
	"net/http"
)

func HandlePost(req *http.Request, ctx context.Context) (results string) {
	var jsonResponse Traffic

	tr := otel.Tracer("crypto-called")
	_, span := tr.Start(ctx, "handle-post")
	span.SetAttributes(attribute.Key("testset").String("value"))
	defer span.End()

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
		results, err := CreateUser(gjLib.Traffic(jsonResponse), ctx)
		if err != nil {
			log.Println(err)
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			return "CRT error"
		}
		return results
	}

	return "POST failed"
}

func HandleGet(req *http.Request, ctx context.Context) (results string) {
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
		_, err := CreateUser(jsonResponse, ctx)
		if err != nil {
			logger.Debug(fmt.Sprintf("--> %s", err))
			return fmt.Sprintf("PING error: %s", err)
		}
	}

	return "GET failed"
}

func HandlePut(req *http.Request, ctx context.Context) (results string) {
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
		results, err := Transaction(jsonResponse, ctx)
		if err != nil {
			logger.Debug(fmt.Sprintf("--> %s", err))
			return fmt.Sprintf("TRAN error: %s", err)
		}
		return results
	} else if jsonResponse.Verb == "ADD" {
		results, err := Deposit(jsonResponse, ctx)
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
		results, err := FindUser(jsonResponse, ctx)
		if err != nil {
			logger.Debug(fmt.Sprintf("--> %s", err))
			return fmt.Sprintf("QUERY error: %s", err)
		}
		return results
	} else if jsonResponse.Verb == "USER" {
		results, err := FindUserAccount(jsonResponse, ctx)
		if err != nil {
			logger.Debug(fmt.Sprintf("--> %s", err))
			return fmt.Sprintf("USER error: %s", err)
		}
		return results
	} else if jsonResponse.Verb == "DLT" {
		results, err := DeleteUser(jsonResponse, ctx)
		if err != nil {
			logger.Debug(fmt.Sprintf("--> %s", err))
			return fmt.Sprintf("DLT error: %s", err)
		}
		return results
	}

	return "PUT failed"
}

func HandleDelete(req *http.Request, ctx context.Context) (results string) {
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

	_, err = DeleteUser(traffic, ctx)
	if err != nil {
		logger.Debug(fmt.Sprintf("--> %s", err))
		return fmt.Sprintf("DLT error: %s", err)
	}

	logger.Debug(fmt.Sprintf("%s", jsonMap["name"]))
	return "gjDelete successful"
}
