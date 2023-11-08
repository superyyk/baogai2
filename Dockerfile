FROM golang:latest

MAINTAINER YourName

RUN mkdir -p /anywhere/YourProject
WORKDIR /anywhere/YourProject
COPY . /anywhere/YourProject

ENV GO111MODULE=on
ENV GOPROXY="https://goproxy.cn" # 更改代理到国内服务器
RUN go mod download

RUN go build main.go

EXPOSE 8080 # 你项目运行的端口

ENTRYPOINT  ["./main"] # 运行你项目的命令
