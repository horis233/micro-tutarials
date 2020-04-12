package main

import (
	"context"

	proto "github.com/micro-in-cn/tutorials/examples/micro-api/rpc/proto"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/client/grpc"
	"github.com/micro/go-micro/v2/errors"
	log "github.com/micro/go-micro/v2/logger"
)

func main() {
	cli := grpc.NewClient(
		// Set the number of retyr based on requirement
		client.Retries(4),
		client.Retry(func(ctx context.Context, req client.Request, retryCount int, err error) (b bool, e error) {
			// Retry on error
			if err != nil {
				// task error
				if err2, ok := err.(*errors.Error); ok {
					// Suppose that any code greater than 1000 is a error
					if err2.Code > 1000 {
						log.Infof("[ERR] Request error, business exception, no retry, err: %s", err)
						return false, nil
					}
				}

				log.Infof("[ERR] Request error, retry% d, will retry soon, err: %s", retryCount, err)
				return true, nil
			}

			return false, nil
		}),
	)

	// Create Client
	greeter := proto.NewExampleService("go.micro.retry.example", cli)

	// Call greeter service
	for i := 0; i < 10; i++ {
		rsp, err := greeter.Call(context.TODO(), &proto.CallRequest{Name: "Micro tutorial"})
		if err != nil {
			log.Infof("[ERR] It is the %d time, Error on request:%s", i, err)
			continue
		}

		log.Infof("[INF] It is the %d time, request result，%v", i, rsp.Message)
	}
}