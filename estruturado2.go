package main

import "fmt"

// START OMIT
type erroNaConexao struct {
	message string
}

func (e erroNaConexao) Error() string {
	return e.message
}

func populaDb() error {
	queue, err := conectaNaFila(args)
	if err != nil {
		return erroNaConexao{"Não foi possível se conectar à fila"} // HLxxx
	}
	db, err := conectaNoDB(args2)
	if err != nil {
		return erroNaConexao{"Não foi possível se conectar ao banco de dados"} // HLxxx
	}
	return nil
}

// END OMIT
