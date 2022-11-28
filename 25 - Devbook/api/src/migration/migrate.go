package migration

import "api/src/models"

func Migrate(){
	AutoMigration(&models.Usuario{})
}