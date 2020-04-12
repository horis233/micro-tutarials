package main

import (
	"context"
	"fmt"
	"log"

	proto "github.com/micro-in-cn/tutorials/examples/client/rpc/proto"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/errors"
)

type Example struct{}

type Foo struct{}

func (e *Example) Call(ctx context.Context, req *proto.CallRequest, rsp *proto.CallResponse) error {
	log.Printf("Receive Example.Call request %v\n", req)
	fmt.Printf("%v\n", req)

	if len(req.Name) == 0 {
		return errors.BadRequest("go.micro.rpc.example", "no content")
	}

	rsp.Message = "RPC Call received your request " + req.Name
	return nil
}

func (f *Foo) Bar(ctx context.Context, req *proto.EmptyRequest, rsp *proto.EmptyResponse) error {
	log.Print("Received Foo.Bar request")
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.rpc.example"),
	)

	service.Init()

	// registry example interface
	proto.RegisterExampleHandler(service.Server(), new(Example))

	// registry foo interface
	proto.RegisterFooHandler(service.Server(), new(Foo))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}