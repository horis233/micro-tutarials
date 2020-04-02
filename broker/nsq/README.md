# NSQ Pubsub

[broker](https://godoc.org/github.com/micro/go-micro/broker#Broker) is an interface for PubSub

## prerequisite：

Install NSQ，following offical website[Instal;](https://nsq.io/deployment/installing.html)，or using [Docker](https://nsq.io/deployment/docker.html)

我们假设读者的NSQ工作在本地标准地址：

- NSQ：127.0.0.1:4150

## How to run

### Client

```bash
cd cli
go run client.go
```

### Server

```bash
cd srv
go run server.go
```
