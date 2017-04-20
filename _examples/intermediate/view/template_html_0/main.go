package main

import (
	"github.com/zengnotes/iris"
	"github.com/zengnotes/iris/adaptors/httprouter"
	"github.com/zengnotes/iris/adaptors/view"
)

func main() {
	app := iris.New(iris.Configuration{Gzip: false, Charset: "UTF-8"}) // defaults to these

	app.Adapt(iris.DevLogger())
	app.Adapt(httprouter.New())

	app.Adapt(view.HTML("./templates", ".html"))

	app.Get("/hi", hi)
	app.Listen(":8080")
}

func hi(ctx *iris.Context) {
	ctx.MustRender("hi.html", struct{ Name string }{Name: "iris"})
}
