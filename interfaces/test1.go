package interfaces

func TestAbs(t *testing.T) {
	err := somaNaturais(1, 2) // HLxxx
	if err != nil {           // HLxxx
		t.Errorf("Erro inválido!") // HLxxx
	} // HLxxx

	err = somaNaturais(-1, 2) // HLxxx
	if err == nil {           // HLxxx
		t.Errorf("Deveria retornar erro!") // HLxxx
	} // HLxxx
}
