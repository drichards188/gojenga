# syntax=docker/dockerfile:1

FROM golang:1.18-alpine

WORKDIR /app

COPY src/go.mod ./
COPY src/go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /gojenga .

EXPOSE 8070

CMD [ "/gojenga" ]