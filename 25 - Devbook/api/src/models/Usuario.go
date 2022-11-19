package models

import "gorm.io/gorm"

// Usuario representa um usuário utilizando a rede social
type Usuario struct{
	gorm.Model
	Nome string 
	Nick string 
	Email string 
	Senha string 
}