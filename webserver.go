package main

import (
	"github.com/kataras/iris"
)

func main() {
    api := iris.New()
    api.Get("/", hello)
    api.Listen(":8888")
}

func hello(ctx *iris.Context){
   ctx.Write("Hello Web!")
}
