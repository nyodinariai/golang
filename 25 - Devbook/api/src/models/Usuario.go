package models

import "gorm.io/gorm"

// Usuario representa um usuário utilizando a rede social
type Usuario struct{
	gorm.Model
	Nome string `json:"nome,omitempty"`
	Nick string `json:"nick,omitempty"`
	Email string `json:"email,omitempty"`
	Senha string `json:"senha,omitempty"`
}