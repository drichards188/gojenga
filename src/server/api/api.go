package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"io"
	"log"
	"net/http"
)

func HandlePost(req *http.Request, ctx context.Context) (results string, err error) {
	var jsonResponse Traffic

	tr := otel.Tracer("crypto-called")
	ctx, span := tr.Start(ctx, "handle-post")
	span.SetAttributes(attribute.Key("my-version").String("1,0,1"))
	defer span.End()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		logger.Debug(fmt.Sprintf("--> %s", err))
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return "", errors.New("POST ReadAll Error")
	}

	err = json.Unmarshal([]byte(body), &jsonResponse)
	if err != nil {
		logger.Debug(fmt.Sprintf("--> %s", err))
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return "", errors.New("POST Unmarshal error")
	}

	if jsonResponse.Verb == "CRT" {
		results, err := CreateUser(Traffic(jsonResponse), ctx)
		if err != nil {
			logger.Debug(fmt.Sprintf("--> %s", err))
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			return "", errors.New("CRT error")
		}
		return results, nil
	}

	return "", errors.New("POST failed")
}

func HandleGet(req *http.Request, ctx context.Context) (results string) {
	var jsonResponse Traffic

	tr := otel.Tracer("mempool-trace")
	ctx, span := tr.Start(ctx, "handle-get")
	span.SetAttributes(attribute.Key("my-version").String("1,0,1"))
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

	_, err = DeleteUser(traffic, ctx)
	if err != nil {
		logger.Debug(fmt.Sprintf("--> %s", err))
		return fmt.Sprintf("DLT error: %s", err)
	}

	logger.Debug(fmt.Sprintf("%s", jsonMap["name"]))
	return "gjDelete successful"
}
