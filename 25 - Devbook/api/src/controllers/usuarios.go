package controllers

import (
	"api/src/auth"
	"api/src/banco"
	"api/src/models"
	"api/src/repository"
	"api/src/responses"
	"api/src/services"
	"encoding/json"
	"errors"
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

	usuario.ID, erro = services.CadastrarUsuario(&usuario)
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
	parametros := mux.Vars(r)

	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil{
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	usuarioIDToken, erro := auth.ExtrairUsuarioId(r)
	if erro != nil{
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	if usuarioID != uint64(usuarioIDToken){
		responses.Erro(w, http.StatusForbidden, errors.New("não é possível atualizar um usuario que não seja o seu"))
		return
	}

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

	if erro = usuario.Preparar("edicao"); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}
	db, erro := banco.Conectar()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return 
	}
	

	repositorio := repository.NovoRepositorioDeUsuarios(db)
	repositorio.AtualizarUsuario(uint(usuarioID), usuario)

	responses.JSON(w, http.StatusOK, usuario)
}

//Deleta um usuario do banco
func DeletarUsuario(w http.ResponseWriter, r *http.Request){
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
	}
	
	repositorio := repository.NovoRepositorioDeUsuarios(db)
	if erro = repositorio.Deletar(uint(usuarioID)); erro != nil{
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}


	responses.JSON(w, http.StatusAccepted, nil)
}