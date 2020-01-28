package main

func somaNaturais(a, b int) (int, error) { // HLxxx
	if err := validaArgs(a, b); err != nil { // HLxxx
		return 0, err // HLxxx
	} // HLxxx
	return a + b, nil // HLxxx
}
