package rotas

import (
	"api/src/controllers"
	"net/http"
)

type UsuarioRoutes interface{
	AddUser (http.ResponseWriter, *http.Request)
}


type usuarioRoute struct{
	usuarioController controllers.UsuarioController
}

func NovoUsuarioRouter (c controllers.UsuarioController) UsuarioRoutes{
	return usuarioRoute{
		usuarioController: c,
	}
}

func (u usuarioRoute) AddUser(w http.ResponseWriter, r *http.Request){

}
var rotasUsuarios = []Rota{
	{
		Uri: "/usuarios",
		Metodo: http.MethodPost,
		Funcao: AddUser(),
		RequerAutenticacao: false,
	},
	{
		Uri: "/usuarios",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscarUsuarios,
		RequerAutenticacao: false,
	},
	{
		Uri: "/usuarios/{usuarioId}",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscarUsuario,
		RequerAutenticacao: false,
	},
		{
		Uri: "/usuarios/{usuarioId}",
		Metodo: http.MethodPut,
		Funcao: controllers.AtualizarUsuario,
		RequerAutenticacao: false,
	},
		{
		Uri: "/usuarios/{usuarioId}",
		Metodo: http.MethodDelete,
		Funcao: controllers.DeletarUsuario,
		RequerAutenticacao: false,
	},
}
