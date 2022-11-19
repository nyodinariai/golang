package repository

import (
	"api/src/models"
	"log"

	"gorm.io/gorm"
)

type Usuarios struct{
	db *gorm.DB
}

//Cria um repositorio de usuarios
func NovoRepositorioDeUsuarios(db *gorm.DB) *Usuarios{
	return &Usuarios{db}
}

func (repositorio Usuarios) Criar(usuario models.Usuario) (uint, error){
	
	statement := repositorio.db.Create(&usuario)
	if err := statement.Error; err != nil{
		log.Fatal((err))
	}

	return uint(usuario.ID), nil
	
}