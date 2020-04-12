package main

import (
	"context"
	"time"

	proto "github.com/micro-in-cn/tutorials/examples/micro-api/rpc/proto"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/errors"
	log "github.com/micro/go-micro/v2/logger"
)

type Example struct{}

func (e *Example) Call(ctx context.Context, req *proto.CallRequest, rsp *proto.CallResponse) error {
	log.Info("Example.Call interface get request，return error")
	if time.Now().Second()%3 == 0 {
		return errors.New("some_id", "some_biz_detail", 1001)
	}

	return errors.New("some_id", "some_detail", 999)
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.retry.example"),
	)

	service.Init()

	// registry example handler
	proto.RegisterExampleHandler(service.Server(), new(Example))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}