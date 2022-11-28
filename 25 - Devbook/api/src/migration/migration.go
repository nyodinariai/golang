package migration

import (
	"api/src/banco"
)

func AutoMigration(model interface{}){
	db, erro := banco.Conectar()
	if erro != nil{
		return
	}
	db.AutoMigrate(&model)
}
