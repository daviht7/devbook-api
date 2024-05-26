package middlewares

import (
	"api/src/autenticacao"
	"api/src/respostas"
	"fmt"
	"net/http"
)

// Autenticar verifica se o usuario fazendo a requisicao está autenticado
func Autenticar(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		if erro := autenticacao.ValidarToken(r); erro != nil {
			respostas.Erro(w, http.StatusUnauthorized, erro)
			return
		}
		fmt.Println("Autenticando...")
		next(w, r)
	}

}
