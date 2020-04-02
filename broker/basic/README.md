# Broker

[broker](https://godoc.org/github.com/micro/go-micro/broker#Broker) is an interface for PubSub

## Contents

- main.go - main excute two go runtines for 10s，one is in charge of publish another is for subscription.

## How to run

If you want to run default http broke, you can run the following command：

```bash
go run main.go
```

If you want to other kind of broke，for example `nats`，you can run the following command：

```bash
export MICRO_BROKER=nats
go run main.go
```

Or：

```bash
go run main.go --broker=nats
```

Or：

```bash
go run main.go --broker=nats --broker_address=127.0.0.1:4222
```
