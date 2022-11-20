package services

import (
	"api/src/models"
	"api/src/repository"
)

type UsuarioService interface{
	Save(models.Usuario) (models.Usuario, error)
}

type usuarioService struct{
	usuarioRepository repository.UsuarioRepository
}

func (u usuarioService) NovoUsuarioService(r repository.UsuarioRepository) UsuarioService{
	return usuarioService{
		usuarioRepository: r,
	}
}

func (u usuarioService) Save(usuario models.Usuario) (models.Usuario, error){

	return u.usuarioRepository.Save(usuario)
}