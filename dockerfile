FROM golang:alpine

MAINTAINER try satria try.satria.a@gmail.com

ADD . /go/src/cdigital

ARG appname=cdigital
ARG http_port=8080

RUN apk update && \
    apk upgrade && \
    apk add git
	
RUN go get -d github.com/gorilla/mux


RUN mkdir -p /go/src/cdigital
COPY . /go/src/cdigital
WORKDIR /go/src/cdigital

RUN go build -o main .
CMD ["./main"]

EXPOSE $http_port