package controllers

import (
	"api/src/banco"
	"api/src/models"
	"api/src/repository"
	"api/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

//Criar usuario insere um usuario no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request){
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario models.Usuario
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil{
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

		if erro = usuario.Preparar("cadastro"); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return 
	}

	repositorio := repository.NovoRepositorioDeUsuarios(db)
	usuario.ID, erro = repositorio.Criar(usuario)
	if erro != nil{
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusCreated, usuario)
}

//BuscarUsuarios busca todos os usuarios salvos no banco
func BuscarUsuarios(w http.ResponseWriter, r *http.Request){
	nomeOuNick := strings.ToLower(r.URL.Query().Get("usuario"))
		db, erro := banco.Conectar()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return 
	}

	repositorio := repository.NovoRepositorioDeUsuarios(db)
	usuarios, erro := repositorio.Buscar(nomeOuNick)
		if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return 
	}

	responses.JSON(w, http.StatusOK, usuarios)
}

//Busca um usuario no banco
func BuscarUsuario(w http.ResponseWriter, r *http.Request){
	parametros := mux.Vars(r)

	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil{
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}
		db, erro := banco.Conectar()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return 
	}

	repositorio := repository.NovoRepositorioDeUsuarios(db)
	usuario, erro := repositorio.BuscarPorID(uint(usuarioID))
		if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return 
	}
	responses.JSON(w, http.StatusOK, usuario)

}

// Atualiza um usuario no banco
func AtualizarUsuario(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Atualizar um Usuario"))
}

//Deleta um usuario do banco
func DeletarUsuario(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Deletar um Usuario"))
}