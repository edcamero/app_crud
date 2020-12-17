package controllers

import (
	"fmt"

	"github.com/edcamero/app_crud/database"
	"github.com/edcamero/app_crud/models"
	"github.com/kataras/iris/v12"
)

func AllAlumnos(ctx iris.Context) []models.Alumno{} {
	alumnos := []models.Alumno{}
	db := database.GetDatabase()
	db.Find(&alumnos)
	// Consulta a la DB - SELECT * FROM contacts WHERE ID = ?
return alumnos
}

func CreateAlumno( ctx iris.Context)){
	alumnos := []models.Alumno{}
	err := json.NewDecoder(ctx.Request().Body).Decode(&user)
	if err != nil {
		// Sí hay algun error al guardar los datos se devolvera un error 500
		fmt.Println(err)
		view.SendErr(ctx, http.StatusInternalServerError)
		return
	}
	// Se codifica el nuevo registro y se devuelve
	j, _ := json.Marshal(user)
	view.SendResponse(ctx, http.StatusCreated, j)
}