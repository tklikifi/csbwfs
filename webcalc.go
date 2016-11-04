package main

import (
	"fmt"
	"log"
	"github.com/kataras/iris"
)

func main() {
    api := iris.New()
    api.Get("/add", add)
    api.Get("/multiply", multiply)
    api.Listen(":8888")
}

func add(ctx *iris.Context){
	a, err := ctx.URLParamInt("first")
	if err != nil {
		log.Print(err)
		ctx.Write("Invalid input")
		return
	}
	b, err := ctx.URLParamInt("second")
	if err != nil {
		log.Print(err)
		ctx.Write("Invalid input")
		return
	}
	ctx.Write(fmt.Sprintf("%d + %d = %d", a, b, a+b))
}

func multiply(ctx *iris.Context){
	a, err := ctx.URLParamInt("first")
	if err != nil {
		log.Print(err)
		ctx.Write("Invalid input")
		return
	}
	b, err := ctx.URLParamInt("second")
	if err != nil {
		log.Print(err)
		ctx.Write("Invalid input")
		return
	}
	ctx.Write(fmt.Sprintf("%d * %d = %d", a, b, a*b))
}
