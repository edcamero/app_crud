package main

import (
	"github.com/edcamero/app_crud/models"
)

func database() {
	db := database.GetDatabase()
	alumno := models.Alumno{}
	db.AutoMigrate(&alumno)
}
