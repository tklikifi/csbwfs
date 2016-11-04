package main

import (
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
	ctx.MustRender("add.html", map[string]interface{}{"a": a, "b": b, "c": a+b})
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
	ctx.MustRender("multiply.html", map[string]interface{}{"a": a, "b": b, "c": a*b})
}
