FROM golang:alpine

WORKDIR /microservice

ADD . .

RUN apk add --no-cache \
        libc6-compat

RUN apk update && apk add --virtual build-dependencies build-base gcc

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon -graceful-kill=true -command="./microservice"
