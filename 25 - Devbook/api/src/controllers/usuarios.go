package controllers

import (
	"api/src/banco"
	"api/src/models"
	"api/src/repository"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// Criari usuario insere um usuario no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request){
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		log.Fatal(erro)
	}

	var usuario models.Usuario
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil{
		log.Fatal(erro)
	}

	db, erro := banco.Conectar()
	if erro != nil {
		log.Fatal(erro)
	}

	repositorio := repository.NovoRepositorioDeUsuarios(db)
	repositorio.Criar(usuario)
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