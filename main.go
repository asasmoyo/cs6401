package main

import (
	"github.com/kataras/iris"
	"os"

	"github.com/asasmoyo/cs6401/tasks"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	iris.Get("/", func(ctx *iris.Context) {
		ctx.WriteString("this is /")
	})

	iris.Get("/first", func(ctx *iris.Context) {
		ctx.WriteString(tasks.First())
	})

	iris.Listen(":" + port)
}
