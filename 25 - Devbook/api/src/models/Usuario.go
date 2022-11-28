package models

import (
	"api/src/security"
	"errors"
	"strings"

	"github.com/badoux/checkmail"
	"gorm.io/gorm"
)

// Usuario representa um usuário utilizando a rede social
type Usuario struct{
	gorm.Model
	Nome string `json:"nome"`
	Nick string `json:"nick"`
	Email string `gorm:"type:varchar(255); unique_index" json:"email"`
	Senha string `json:"senha,omitempty"`
}

// Preparar vai chamar os métodos para validar e formatar o usuário recebido
func (usuario *Usuario) Preparar(etapa string) error {
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}

	if erro := usuario.formatar(etapa); erro != nil {
		return erro
	}

	return nil
}

func (usuario *Usuario) validar(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("o nome é obrigatório e não pode estar em branco")
	}

	if usuario.Nick == "" {
		return errors.New("o nick é obrigatório e não pode estar em branco")
	}

	if usuario.Email == "" {
		return errors.New("o e-mail é obrigatório e não pode estar em branco")
	}

	if erro := checkmail.ValidateFormat(usuario.Email); erro != nil{
		return errors.New("o e-mail inserido é inválido")
	}

	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("o e-mail é obrigatório e não pode estar em branco")
	}

	return nil
}

func (usuario *Usuario) formatar(etapa string) error {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)

	if etapa == "cadastro"{
		senhaComHash, erro := security.Hash(usuario.Senha)
		if erro != nil{
			return erro
		}
		usuario.Senha = string(senhaComHash)
	}

	return nil
}
