package controllers

import (
	"api/src/models"
	"api/src/responses"
	"api/src/services"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type UsuarioController interface{
	AddUser (http.ResponseWriter, *http.Request)
}


type usuarioController struct{
	usuarioService services.UsuarioService
}

func NovoUsuarioController (s services.UsuarioService) UsuarioController{
	return usuarioController{
		usuarioService: s,
	}
}

func (u usuarioController) AddUser(w http.ResponseWriter, r *http.Request){
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

	usuario, erro = u.usuarioService.Save(usuario)
	if erro != nil{
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	responses.JSON(w, http.StatusCreated, usuario)
}

// Criar usuario insere um usuario no banco de dados
// func CriarUsuario(w http.ResponseWriter, r *http.Request){
// 	corpoRequest, erro := ioutil.ReadAll(r.Body)
// 	if erro != nil {
// 		responses.Erro(w, http.StatusUnprocessableEntity, erro)
// 		return
// 	}

// 	var usuario models.Usuario
// 	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil{
// 		responses.Erro(w, http.StatusBadRequest, erro)
// 		return
// 	}

// 	db, erro := banco.Conectar()
// 	if erro != nil {
// 		responses.Erro(w, http.StatusInternalServerError, erro)
// 		return 
// 	}

// 	repositorio := repository.NovoRepositorioDeUsuarios(db)
// 	usuario.ID, erro = repositorio.Criar(usuario)
// 	if erro != nil{
// 		responses.Erro(w, http.StatusInternalServerError, erro)
// 		return
// 	}

// 	responses.JSON(w, http.StatusCreated, usuario)
// }

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