package repository

import (
	"api/src/models"

	"gorm.io/gorm"
)

type usuarios struct{
	db *gorm.DB
}

//Cria um repositorio de usuarios
func NovoRepositorioDeUsuarios(db *gorm.DB) *usuarios{
	return &usuarios{db}
}

func (u usuarios) Criar(usuario models.Usuario) (uint64, error){
	
	return 0, nil
}