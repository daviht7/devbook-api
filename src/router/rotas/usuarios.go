package rotas

import (
	controller "api/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{

	{
		URI:                "/usuarios",
		Metodo:             http.MethodPost,
		Funcao:             controller.CriarUsuario,
		RequerAutenticacao: false,
	},
	{
		URI:                "/usuarios",
		Metodo:             http.MethodGet,
		Funcao:             controller.BuscarUsuarios,
		RequerAutenticacao: false,
	},
	{
		URI:                "/usuarios/{usuarioId}",
		Metodo:             http.MethodGet,
		Funcao:             controller.Buscarusuario,
		RequerAutenticacao: false,
	},
	{
		URI:                "/usuarios/{usuarioId}",
		Metodo:             http.MethodPut,
		Funcao:             controller.AtualizarUsuario,
		RequerAutenticacao: false,
	},
	{
		URI:                "/usuarios/{usuarioId}",
		Metodo:             http.MethodDelete,
		Funcao:             controller.DeletarUsuario,
		RequerAutenticacao: false,
	},
}
