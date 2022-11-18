package banco

import (
	"api/src/config"
	"api/src/models"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Conectar()(*gorm.DB, error){

	db, err := gorm.Open(mysql.Open(config.StringConexaoBanco), &gorm.Config{})
	if err != nil{
		log.Fatalln(err)
		return nil, err
	}

	db.AutoMigrate(&models.Usuario{})

	fmt.Println("Conexão Realizada com sucesso")

	return db, nil

}