FROM golang:alpine

WORKDIR /microservice

ADD . .

RUN go mod download

RUN ulimit -s 1048576

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon -command="./microservice"