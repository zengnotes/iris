package main

import (
	"github.com/zengnotes/iris"
	"github.com/zengnotes/iris/adaptors/httprouter"
	"github.com/zengnotes/iris/middleware/pprof"
)

func main() {
	app := iris.New()
	app.Adapt(iris.DevLogger())
	app.Adapt(httprouter.New())

	app.Get("/", func(ctx *iris.Context) {
		ctx.HTML(iris.StatusOK, "<h1> Please click <a href='/debug/pprof'>here</a>")
	})

	app.Get("/debug/pprof/*action", pprof.New())
	//                              ___________
	// Note:
	// if you prefer gorillamux adaptor, then
	// the wildcard for gorilla mux (as you already know) is '{action:.*}' instead of '*action'.
	app.Listen(":8080")
}
