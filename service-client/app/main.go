package main

import (
	"fmt"

	"github.com/balashVI/go-micro-test/proto"
	"github.com/micro/go-micro"
	"golang.org/x/net/context"
	"strconv"
	"time"
)

type App struct {
	ToDoServiceClient proto.ToDoServiceClient
}

var app = new(App)

func init() {

	service := micro.NewService(
		micro.Name("go.micro-test.todo"),
		micro.Version("latest"),
	)
	service.Init()

	app.ToDoServiceClient = proto.NewToDoServiceClient("go.micro-test.todo", service.Client())
}

func main() {
	for {
		time.Sleep(3 * time.Second)
		rsp, err := app.ToDoServiceClient.Ping(context.TODO(), new(proto.Empty))

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(rsp.Message)

			listResp, _ := app.ToDoServiceClient.List(context.TODO(), new(proto.ListRequest))
			fmt.Println(listResp.Todos)

			newToDo := proto.ToDo{
				Id:      int64(len(listResp.Todos)),
				Message: "ToDo #" + strconv.Itoa(len(listResp.Todos)),
				Done:    false,
			}
			app.ToDoServiceClient.Add(context.TODO(), &newToDo)
		}
	}
}
