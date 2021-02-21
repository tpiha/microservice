FROM golang:alpine

WORKDIR /microservice

ADD . .

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon -command="./microservice"