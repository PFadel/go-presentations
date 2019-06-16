package main

import "fmt"

// START OMIT

type erroNaConexao interface {
	Error() string
	ParseMensagem() string // HLxxx
}

type erroNoDB struct {
	message string
}

func (e erroNoDB) Error() string {
	return e.message
}

func (e erroNoDB) ParseMensagem() string { // HLxxx
	return fmt.Sprintf("Houve um erro na conexão ao DB,
						mensagem do driver: %s", e.message)
}

// END OMIT
