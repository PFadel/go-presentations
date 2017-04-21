package main

// START OMIT
type requestTimeOut struct {
	message string
}

func (e requestTimeOut) Error() string {
	return e.message
}

type erroInternoNoServidor struct {
	message string
	status  int
}

func (e erroInternoNoServidor) Error() string {
	return e.message
}

func (e erroInternoNoServidor) StatusCode() int {
	return e.status
}

// END OMIT
