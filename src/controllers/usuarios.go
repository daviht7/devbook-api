package controller

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// CriarUsuario insere um usuario no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		log.Fatal(erro)
	}

	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		log.Fatal(erro)
	}

	db, erro := banco.Conectar()
	if erro != nil {
		log.Fatal(erro)
	}

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	repositorio.Criar(usuario)

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
