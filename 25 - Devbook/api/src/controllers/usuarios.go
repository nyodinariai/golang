package controllers

import "net/http"

// Criari usuario insere um usuario no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Criando Usuario"))
}

//BuscarUsuarios busca todos os usuarios salvos no banco
func BuscarUsuarios(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Buscar Usuarios"))
}

//Busca um usuario no banco
func BuscarUsuario(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Buscar um Usuario"))
}

// Atualiza um usuario no banco
func AtualizarUsuario(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Atualizar um Usuario"))
}

//Deleta um usuario do banco
func DeletarUsuario(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Deletar um Usuario"))
}