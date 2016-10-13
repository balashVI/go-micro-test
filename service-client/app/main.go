package main

import (
	"fmt"

	"github.com/balashVI/go-micro-test/proto"
	"github.com/micro/go-micro"
	"golang.org/x/net/context"
)

type App struct {
	HeathCheckClient proto.HealthCheckClient
}

var app = new(App)

func init() {
	service := micro.NewService(
		micro.Name("go.micro-test.healthcheck"),
		micro.Version("latest"),
	)
	service.Init()

	app.HeathCheckClient = proto.NewHealthCheckClient("go.micro-test.healthcheck", service.Client())
}

func main() {
	rsp, err := app.HeathCheckClient.Ping(context.TODO(), &proto.PingRequest{})

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rsp.Message)
}
