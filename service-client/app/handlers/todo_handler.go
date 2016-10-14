package handlers

import (
	"github.com/balashVI/go-micro-test/proto"
	"github.com/kataras/iris"
	"golang.org/x/net/context"
)

type ToDoHandler struct {
	ToDoServiceClient proto.ToDoServiceClient
}

func (h ToDoHandler) Ping(c *iris.Context) {
	resp, err := h.ToDoServiceClient.Ping(context.TODO(), new(proto.Empty))
	if err != nil {
		c.JSON(iris.StatusInternalServerError, iris.Map{"error": err.Error()})
	} else {
		c.JSON(iris.StatusOK, resp)
	}
}

func (h ToDoHandler) List(c *iris.Context) {
	page, _ := c.URLParamInt("page")
	count, _ := c.URLParamInt("count")

	req := proto.ListRequest{
		Count: int32(count),
		Page: int32(page),
	}
	resp, err := h.ToDoServiceClient.List(context.TODO(), &req)
	if err != nil {
		c.JSON(iris.StatusInternalServerError, iris.Map{"error": err.Error()})
	} else {
		c.JSON(iris.StatusOK, resp)
	}
}

func (h ToDoHandler) Get(c *iris.Context) {
	id, err := c.ParamInt64("id")
	if err != nil {
		c.JSON(iris.StatusBadRequest, iris.Map{"error": err.Error()})
	}

	req := proto.GetRequest{Id: id}
	resp, err := h.ToDoServiceClient.Get(context.TODO(), &req)
	if err != nil {
		c.JSON(iris.StatusInternalServerError, iris.Map{"error": err.Error()})
	} else {
		c.JSON(iris.StatusOK, resp)
	}
}

func (h ToDoHandler) Add(c *iris.Context) {
	toDo := proto.ToDo{}
	err := c.ReadJSON(&toDo)
	if err != nil {
		c.JSON(iris.StatusBadRequest, iris.Map{"error": err.Error()})
	}

	resp, err := h.ToDoServiceClient.Add(context.TODO(), &toDo)
	if err != nil {
		c.JSON(iris.StatusInternalServerError, iris.Map{"error": err.Error()})
	} else {
		c.JSON(iris.StatusOK, resp)
	}
}
