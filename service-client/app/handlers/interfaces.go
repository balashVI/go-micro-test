package handlers

import "github.com/kataras/iris"

type ToDoHandlerInterface interface {
	Ping (c *iris.Context)

	List (c *iris.Context)
	Get (c *iris.Context)
	Add (c *iris.Context)
}