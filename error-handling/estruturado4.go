package main

import "fmt"

// START OMIT

type statusCodeInesperado interface {
	Error() string // HLxxx
	StatusCode() int
}

type erroNoRequest interface {
	Error() string // HLxxx
}

// END OMIT
