package main

import "errors"

// START OMIT
const (
	erroA = "O primeiro argumento precisa ser natural"
	erroB = "O segundo argumento precisa ser natural"
)

func somaNaturais(a, b int) (int, error) {
	if a < 0 {
		return 0, errors.New(erroA) // HLxxx
	}
	if b < 0 {
		return 0, errors.New(erroB) // HLxxx
	}
	return a + b, nil
}

// END OMIT

func main() {
	_, err := somaNaturais(1, -2)
	if err != nil {
		if err.Error() == erroA { // HLxxx
			println("Tratamento super seguro")
		} else if err.Error() == erroB { // HLxxx
			println("Tratamento super seguro diferente")
		} else { // HLxxx
			println("Erro desconhecido!")
		}
	}
	println("Deu tudo certo no final")
}
