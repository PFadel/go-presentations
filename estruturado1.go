package main

import "fmt"

// START OMIT
type erroA struct {
	message string
}

func (e erroA) Error() string {
	return e.message
}

type erroB struct {
	message string
}

func (e erroB) Error() string {
	return e.message
}

// END OMIT

func somaNaturais(a, b int) (int, error) {
	if a < 0 {
		return 0, erroA{fmt.Sprintf("Mensagem super descritiva formatada pelos deuses %d", a)} // HLxxx
	}
	if b < 0 {
		return 0, erroB{fmt.Sprintf("Mensagem super descritiva formatada pelos deuses %d", b)} // HLxxx
	}
	return a + b, nil
}

func main() {
	_, err := somaNaturais(1, -2)
	if err != nil {
		if _, ok := err.(erroA); ok { // HLxxx
			// Tratamento super seguro
			println(err.Error())
		} else if _, ok := err.(erroB); ok { // HLxxx
			// Tratamento super seguro diferente
			println(err.Error())
		} else { // HLxxx
			println("Erro desconhecido!")
		}
	}
	println("Deu tudo certo no final")
}
