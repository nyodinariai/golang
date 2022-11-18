package banco

import (
	"api/src/config"
	"api/src/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Conectar()(*gorm.DB){

	db, err := gorm.Open(mysql.Open(config.StringConexaoBanco), &gorm.Config{})
	if err != nil{
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Usuario{})

	return db

}