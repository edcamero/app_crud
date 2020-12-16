package database

import (
	"app_crud/models"

	"github.com/edcamero/app_crud/backend/database"
)

func main() {
	db := database.GetConnection()
	alumno := models.Alumno{}
	db.AutoMigrate(&alumno)
}
