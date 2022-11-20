package repository

import (
	"api/src/models"
	"log"

	"gorm.io/gorm"
)

type usuarioRepository struct{
	DB *gorm.DB
}

type UsuarioRepository interface{
	Save(models.Usuario) (models.Usuario, error)
}

//Cria um repositorio de usuarios
func NovoRepositorioDeUsuarios(db *gorm.DB) *usuarioRepository{
	return &usuarioRepository{
		DB: db,
	}
}

func (u usuarioRepository) Migrate() error{
	log.Print("[UsuarioRepository]...Migrate")
	return u.DB.AutoMigrate(&models.Usuario{})
}

func (u usuarioRepository) Save(usuario models.Usuario) (models.Usuario, error){
	log.Print("[UsuarioRepository]...Save")
	err := u.DB.Create(&usuario).Error
	return usuario, err
}

