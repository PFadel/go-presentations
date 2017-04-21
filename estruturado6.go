package main

// START OMIT
func main() {
	_, err := executaGet("url", nil)
	if err != nil {
		if _, ok := err.(statusCodeInesperado); ok { // HLxxx
			println("Tratamento 1")
		} else if _, ok := err.(erroNoRequest); ok { // HLxxx
			println("Tratamento 2")
		} else { // HLxxx
			...
		}
	}
}

// END OMIT