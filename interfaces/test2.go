package interfaces

// START OMIT

import "errors"

type mockValidadorFalha struct{}

type mockValidadorSucesso struct{}

func (f *mockValidadorFalha) ValidaArgs(a, b int) error { // HLxxx
	return errors.New("test-error") // HLxxx
} // HLxxx

func (f *mockValidadorSucesso) ValidaArgs(a, b int) error { // HLxxx
	return nil // HLxxx
} // HLxxx

// END OMIT

func TestAbs(t *testing.T) {
	err := somaNaturais(1, 2, &mockValidadorSucesso{}) // HLxxx
	if err != nil {                                    // HLxxx
		t.Errorf("Erro inválido!") // HLxxx
	} // HLxxx

	err = somaNaturais(-1, 2, &mockValidadorFalha{}) // HLxxx
	if err == nil {                                  // HLxxx
		t.Errorf("Deveria retornar erro!") // HLxxx
	} // HLxxx
}
