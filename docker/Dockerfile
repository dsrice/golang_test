FROM golang:1.19.2

RUN apt-get update

# appディレクトリの作成
RUN mkdir /go/src/app

WORKDIR /go/src/app
ADD . /go/src/app