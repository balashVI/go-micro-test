package main

import (
	"github.com/balashVI/go-micro-test/proto"
	"github.com/balashVI/go-micro-test/service-client/app/handlers"
	"github.com/kataras/iris"
	"github.com/micro/go-micro"
)

type App struct {
	ToDoServiceClient proto.ToDoServiceClient

	ToDoHandler handlers.ToDoHandlerInterface
}

var app = new(App)

func init() {
	service := micro.NewService(
		micro.Name("go.micro-test.todo"),
		micro.Version("latest"),
	)
	service.Init()

	app.ToDoServiceClient = proto.NewToDoServiceClient("go.micro-test.todo", service.Client())

	app.ToDoHandler = handlers.ToDoHandler{ToDoServiceClient: app.ToDoServiceClient}
}

func main() {
	iris.Get("/ping", app.ToDoHandler.Ping)

	iris.Get("/list", app.ToDoHandler.List)

	todo := iris.Party("/todo")
	{
		todo.Get("/:id", app.ToDoHandler.Get)
		todo.Put("/", app.ToDoHandler.Add)
	}

	iris.Listen(":8080")
}
