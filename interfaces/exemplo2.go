package main

type Validador interface { // HLxxx
	ValidaArgs(a, b int) error // HLxxx
} // HLxxx

func somaNaturais(a, b int, v Validador) (int, error) {
	if err := v.ValidaArgs(a, b); err != nil {
		return 0, err
	}
	return a + b, nil
}
