package repository

import (
	"api/src/models"
	"fmt"
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

//Todos os usuarios que atendem um filtro de nome ou nick
func (repositorio Usuarios) Buscar(nomeOuNick string)([]models.Usuario, error){
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick)
	var usuarios []models.Usuario
	statement := repositorio.db.Select("ID, nome, nick, email").Where("nome like ?", nomeOuNick).Or("nick like ?", nomeOuNick).Find(&usuarios)
	if erro := statement.Error; erro != nil{
		log.Fatal((erro))
	}

	return usuarios, nil
	
	}

	func (repositorio Usuarios) BuscarPorID(id uint)(models.Usuario, error){
		var usuarios = models.Usuario{}
		statement := repositorio.db.Select("ID, nome, nick, email").First(&usuarios, id)

		if erro := statement.Error; erro != nil{
		log.Fatal((erro))
	}

	return usuarios, nil
	}
