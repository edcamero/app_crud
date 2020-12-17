package router

import (
	"github.com/kataras/iris/v12"
)

func AddRutas(app *iris.Application) {

	api := app.Party("/api")
	api.Post("/alumno/registrar", controllers.Create)
}
