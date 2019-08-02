FROM alpine

MAINTAINER try satria try.satria.a@gmail.com

ADD . /go/src/cdigital

ARG appname=cdigital
ARG http_port=8080




RUN mkdir -p /go/src/cdigital
COPY . /go/src/cdigital
WORKDIR /go/src/cdigital

CMD ["./main"]


EXPOSE $http_port