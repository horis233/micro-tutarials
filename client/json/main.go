package main

import (
	"context"
	"fmt"

	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/metadata"
	"github.com/micro/go-micro/v2/registry/mdns"
)

// Request parameter structure as long as the other party's service can recognize it
type whatEverReq struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

// The response structure can be as long as our service can recognize it
type whatEverRsp struct {
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func main() {
	cli := client.NewClient(
		// Just register with the directory service
		client.Registry(mdns.NewRegistry()),
	)

	// call the structure of the target service
	req := cli.NewRequest("go.micro.rpc.example", "Example.Call",
		&whatEverReq{
			Name: "John",
		},
		// When you are not sure about the service of the other party, you need to use JSON format instead of protobuf
		client.WithContentType("application/json"))

	// Custom metadata
	ctx := metadata.NewContext(context.Background(), map[string]string{
		"X-User-Id": "john",
		"X-From-Id": "script",
	})

	rsp := &whatEverRsp{}

	// Call the service
	if err := cli.Call(ctx, req, rsp); err != nil {
		fmt.Println("call err: ", err, rsp)
		return
	}

	fmt.Println("rsp: ", rsp.Message)
}