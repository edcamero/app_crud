package main

import (
	"os"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081" //localhost
	}

	tmpl := iris.HTML("./views", ".html")

	tmpl.Reload(true)

	tmpl.AddFunc("greet", func(s string) string {
		return "Greetings " + s + "!"
	})
	app.RegisterView(tmpl)
	app.Get("/", func(ctx iris.Context) {
		//ctx.ViewData("message", "Crud de estudiante")
		ctx.View("index.html")
	})
	app.Get("/add", func(ctx iris.Context) {
		ctx.View("add.html")
	})
	app.Get("/list", func(ctx iris.Context) {
		ctx.View("list.html")
	})

	app.Get("/update", func(ctx iris.Context) {
		ctx.View("update.html")
	})

	app.Listen(":" + port)
}
