package main

import (
	"context"
	"encoding/json"

	proto "github.com/micro-in-cn/tutorials/examples/micro-api/rpc/proto"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
)

type Example struct{}

func (e *Example) Call(ctx context.Context, req *proto.CallRequest, rsp *proto.CallResponse) error {
	log.Info("Example.Call interface get your request，return success")

	b, _ := json.Marshal(map[string]string{
		"message": "We have already got your request, " + req.Name,
	})

	// Set return value
	rsp.Message = string(b)

	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.retry.example"),
	)

	service.Init()

	// Registry example handler
	proto.RegisterExampleHandler(service.Server(), new(Example))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}