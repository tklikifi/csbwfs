package main

import (
	"github.com/kataras/iris"
)

type Calculation struct {
	First 	int
	Second 	int
	Result  int
}

func main() {
    api := iris.New()
    api.Get("/add", func(ctx *iris.Context) {
    	if err := ctx.Render("add.html", nil); err != nil {
    		iris.Logger.Printf(err.Error())
    	}
    })
    api.Post("/add", func(ctx *iris.Context) {
		c := Calculation{}
		err := ctx.ReadForm(&c)
		if err != nil {
    		iris.Logger.Printf(err.Error())
		}
		c.Result = c.First+c.Second
    	if err := ctx.Render("add.html", map[string]interface{}{"First": c.First, "Second": c.Second, "Result": c.Result}); err != nil {
    		iris.Logger.Printf(err.Error())
    	}
	})
    api.Get("/multiply", func(ctx *iris.Context) {
    	if err := ctx.Render("multiply.html", nil); err != nil {
    		iris.Logger.Printf(err.Error())
    	}
    })
    api.Post("/multiply", func(ctx *iris.Context) {
		c := Calculation{}
		err := ctx.ReadForm(&c)
		if err != nil {
    		iris.Logger.Printf(err.Error())
		}
		c.Result = c.First*c.Second
    	if err := ctx.Render("multiply.html", map[string]interface{}{"First": c.First, "Second": c.Second, "Result": c.Result}); err != nil {
    		iris.Logger.Printf(err.Error())
    	}
	})
    api.Listen(":8888")
}
