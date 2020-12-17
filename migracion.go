package main

import (
	"github.com/edcamero/app_crud/database"
	"github.com/edcamero/app_crud/models"
)

func main() {
	db := database.GetDatabase()
	alumno := models.Alumno{}
	db.AutoMigrate(&alumno)
}
