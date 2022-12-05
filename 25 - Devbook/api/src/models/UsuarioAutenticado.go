package models

type UsuarioAutenticado struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}