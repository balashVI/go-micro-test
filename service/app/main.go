package main

import (
	"fmt"

	"github.com/micro/go-micro"
	"github.com/balashVI/go-micro-test/proto"
	"golang.org/x/net/context"
)

type HealthCheck struct{}

func (hc HealthCheck) Ping(ctx context.Context, req *proto.PingRequest, rsp *proto.PingResponse) error{
	rsp.Message = "Pong"
	return nil
}

func main() {
	// Create a new service
	service := micro.NewService(
		micro.Name("go.micro-test.healthcheck"),
		micro.Version("latest"),
	)

	// Init will parse the command line flags
	service.Init()

	// Register handler
	proto.RegisterHealthCheckHandler(service.Server(), new(HealthCheck))

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
