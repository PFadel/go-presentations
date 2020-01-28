package main

import "errors"

func validaArgs(a, b int) error {
	if a < 0 {
		return errors.New("A precisa ser positivo não nulo")
	}
	if b < 0 {
		return errors.New("B precisa ser positivo não nulo")
	}
	return nil
}

func somaNaturais(a, b int) (int, error) { // HLxxx
	if err := validaArgs(a, b); err != nil { // HLxxx
		return 0, err // HLxxx
	} // HLxxx
	return a + b, nil // HLxxx
} // HLxxx
