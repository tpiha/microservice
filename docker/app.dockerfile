FROM golang:alpine

WORKDIR /microservice

ADD . .

RUN apk add --no-cache \
        libc6-compat

RUN apk update && apk add --virtual build-dependencies build-base gcc wget git

RUN go mod download

# RUN go get github.com/githubnemo/CompileDaemon

# ENTRYPOINT CompileDaemon -command="./microservice"

# ENTRYPOINT /microservice/microservice
ENTRYPOINT go build -tags netgo -a -v && ./microservice