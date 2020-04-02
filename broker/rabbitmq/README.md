# Broker

[broker](https://godoc.org/github.com/micro/go-micro/broker#Broker) is an interface for PubSub

## Contents

- main.go - main excute two go runtines for 10s，one is in charge of publish another is for subscription.

- plugin.go integrate rabbitmq plugin

## How to run

```bash
go run main.go plugin.go --broker=rabbitmq
```
