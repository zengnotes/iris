package main

import (
	"github.com/zengnotes/iris"
	"github.com/zengnotes/iris/adaptors/httprouter"
)

const host = "127.0.0.1:443"

func main() {
	app := iris.New()
	// output startup banner and error logs on os.Stdout
	app.Adapt(iris.DevLogger())
	// set the router, you can choose gorillamux too
	app.Adapt(httprouter.New())

	app.Get("/", func(ctx *iris.Context) {
		ctx.Writef("Hello from the server")
	})

	app.ListenUNIX("/tmp/srv.sock", 0666)
}
