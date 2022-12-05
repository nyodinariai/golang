package repository

import (
	"api/src/models"
	"fmt"

	"gorm.io/gorm"
)

type Usuarios struct{
	db *gorm.DB
}


//Cria um repositorio de usuarios
func NovoRepositorioDeUsuarios(db *gorm.DB) *Usuarios{
	return &Usuarios{db}
}

func (repositorio Usuarios) Criar(usuario *models.Usuario) (uint, error){
	
	if erro := repositorio.db.Create(&usuario).Error; erro != nil{
		return 0, erro
	}

	return uint(usuario.ID), nil
	
}

//Todos os usuarios que atendem um filtro de nome ou nick
func (repositorio Usuarios) Buscar(nomeOuNick string)([]models.Usuario, error){
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick)
	var usuarios []models.Usuario
	statement := repositorio.db.Select("ID, nome, nick, email").Where("nome like ?", nomeOuNick).Or("nick like ?", nomeOuNick).Find(&usuarios)
	if erro := statement.Error; erro != nil{
		return nil, erro
	}

	return usuarios, nil
	
	}

	func (repositorio Usuarios) BuscarPorID(id uint)(models.Usuario, error){
		var usuario = models.Usuario{}
		statement := repositorio.db.Select("ID, nome, nick, email").First(&usuario, id)
		if erro := statement.Error; erro != nil{
			return usuario, erro			
		}

	return usuario, nil
	}

func (repositorio Usuarios) AtualizarUsuario(usuarioID uint, usuario models.Usuario)(models.Usuario, error){
	statement := repositorio.db.Model(&usuario).Where("ID = ?", usuarioID).Updates(models.Usuario{Nome: usuario.Nome, Nick: usuario.Nick, Email: usuario.Email})
	if erro := statement.Error; erro != nil{
		return usuario, erro
	}

	return usuario, nil
	
}

func (repositorio Usuarios) Deletar(ID uint) error{
	var usuario models.Usuario
	statement := repositorio.db.Delete(&usuario, ID)
	if erro := statement.Error; erro != nil  {
		return erro
	}

	return nil
}

//Busca um usuario por e-mail e retorna um usuario por email, id e senha
func (repositorio Usuarios) BuscarPorEmail(email string) (models.Usuario, error) {
	var usuario models.Usuario
	statement := repositorio.db.Model(&models.Usuario{}).Where("email = ?", email).Find(&usuario)
	if erro := statement.Error; erro != nil{
		return usuario, erro
	}

	return usuario, nil
}
