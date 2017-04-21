package main

import "errors"

func somaNaturais(a, b int) (int, error) {
	if a < 0 {
		return 0, errors.New("O primeiro argumento precisa ser natural") // HLxxx
	}
	if b < 0 {
		return 0, errors.New("O segundo argumento precisa ser natural") // HLxxx
	}
	return a + b, nil
}

func main() {
	_, err := somaNaturais(-1, 2)
	if err != nil {
		if err.Error() == "O primeiro argumento precisa ser natural" { // HLxxx
			println("Tratamento super seguro")
		} else if err.Error() == "O segundo argumento precisa ser natural" { // HLxxx
			println("Tratamento super seguro diferente")
		} else { // HLxxx
			println("Erro desconhecido!")
		}
	}
	println("Deu tudo certo no final")
}
