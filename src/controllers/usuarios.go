package controller

import "net/http"

// CriarUsuario insere um usuario no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("criando usuario!"))

}

// BuscarUsuarios busca todos os usuarios
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("buscando todos usuarios!"))

}

// Buscarusuario busca um usuario especifico por id
func Buscarusuario(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("buscando usuario!"))

}

// AtualizarUsuario altera informações de um usuario
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("atualizar usuario!"))

}

// DeletarUsuario exclui as informações de um usuario
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("deletar usuario!"))

}
