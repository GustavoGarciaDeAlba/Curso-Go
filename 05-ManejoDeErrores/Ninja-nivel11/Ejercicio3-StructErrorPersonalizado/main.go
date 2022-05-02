package main

import (
	"fmt"
)

type errorPer struct {
	info string
}

func (ep errorPer) Error() string {
	return fmt.Sprintf("Error: %v", ep.info)
}

func main() {
	e1 := errorPer{
		info: "Hola, este es el error",
	}
	foo(e1)
}

func foo(e error) {
	fmt.Println("foo", e)
}
