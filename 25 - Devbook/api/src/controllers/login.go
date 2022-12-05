package controllers

import (
	"api/src/auth"
	"api/src/models"
	"api/src/responses"
	"api/src/security"
	"api/src/services"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

func Login(w http.ResponseWriter, r *http.Request){
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

	var usuarioAutenticado models.Usuario
	usuarioAutenticado, erro = services.BuscarPorEmail(usuario.Email)
	if erro != nil{
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}


	if erro = security.VerificarSenha(usuarioAutenticado.Senha, usuario.Senha); erro != nil{
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	token, erro := auth.Token(usuarioAutenticado.ID)
	if erro != nil{
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	usuarioID := strconv.FormatUint(uint64(usuarioAutenticado.ID), 10)

	responses.JSON(w, http.StatusOK, models.UsuarioAutenticado{ID: usuarioID, Token: token})
}