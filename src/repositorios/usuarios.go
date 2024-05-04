package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

// Usuarios representa um repositório de um usuarios
type Usuarios struct {
	db *sql.DB
}

// NovoRepositoriosDeUsuarios cria um repositório de usuarios
func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {

	return &Usuarios{db}
}

// Criar insere um usuário no banco de dados
func (u Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {

	return 0, nil

}
