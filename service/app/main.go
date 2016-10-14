package main

import (
	"fmt"

	"github.com/balashVI/go-micro-test/proto"
	"github.com/balashVI/go-micro-test/service/app/services"
	"github.com/micro/go-micro"
)

func main() {
	// Create a new service
	service := micro.NewService(
		micro.Name("go.micro-test.todo"),
		micro.Version("latest"),
	)

	// Init will parse the command line flags
	service.Init()

	// Register handler
	proto.RegisterToDoServiceHandler(service.Server(), new(services.ToDoService))

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
