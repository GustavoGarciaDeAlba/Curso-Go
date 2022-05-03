package main

import (
	"fmt"

	"CursoGo/06-Documentacion/Ninja-nivel12/Ejercicio1-PaquetePerro/perro"
)

type dog struct {
	name string
	age  int
}

func main() {
	d1 := dog{name: "Chester", age: perro.Age(2)}

	fmt.Println(d1)
}
