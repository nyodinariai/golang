package services

import (
	"api/src/banco"
	"api/src/models"
	"api/src/repository"
)


func CadastrarUsuario(usuario *models.Usuario) (uint, error){

	db, erro := banco.Conectar()
	if erro != nil {
		return 0, erro
	}
	
	repositorio := repository.NovoRepositorioDeUsuarios(db)
	usuarioCadastrado, erro := repositorio.Criar(usuario)
	if erro != nil{
		return 0, erro
	}

	return usuarioCadastrado, nil
}

func BuscarPorEmail(email string) (models.Usuario, error){
		db, erro := banco.Conectar()
		if erro != nil {
		return models.Usuario{}, erro
		}

		repositorio := repository.NovoRepositorioDeUsuarios(db)
		usuario, erro := repositorio.BuscarPorEmail(email)
		if erro != nil{
		return usuario, erro
		}

		return usuario, nil
}