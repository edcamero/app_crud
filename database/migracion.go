package database

import (
	"github.com/edcamero/app_crud/database"
	"github.com/edcamero/app_crud/models"
)

func main() {
	db := database.GetConnection()
	alumno := models.Alumno{}
	db.AutoMigrate(&alumno)
}
