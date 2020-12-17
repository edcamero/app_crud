package router

import (
	Controllers "github.com/edcamero/app_crud/controllers"
	"github.com/kataras/iris/v12"
)

func AddRutas(app *iris.Application) {

	app.Post("add", Controllers.CreateAlumno)

	app.Get("/add", func(ctx iris.Context) {
		ctx.View("add.html")
	})

	app.Get("/list", Controllers.AllAlumnos)

	app.Get("/update", Controllers.GetAlumno)

	app.Post("/update", Controllers.UpdateAlumno2)
	app.Get("/delete", Controllers.DeleteAlumno)
}
