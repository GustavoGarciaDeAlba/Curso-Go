package main

import (
	"fmt"
)

type persona struct {
	nombre   string
	apellido string
	age      int
}

type agenteSecreto struct {
	persona
	lpm bool
}

// Interfaces
type humano interface {
	presentarse()
}

// El tipo agenteSecreto implementa la interfaz humano

func main() {
	// Structs
	practicaEstructuras()
	y := []int{1, 2, 3, 4, 5}
	x := foo(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("El valor almacenado en la variable es", x)
	z := foo(y...) // el operador... despliega los valores del slice para poder utilizarlo en la funcion con parametros variables
	fmt.Println("El valor almacenado en la variable es", z)

	//Metodos
	p5 := agenteSecreto{
		persona: persona{
			nombre:   "Foo",
			apellido: "Bar",
			age:      10,
		},
		lpm: true,
	}
	p5.presentarse()

	p := persona{
		nombre:   "Carito",
		apellido: "Guz",
		age:      12,
	}
	bar(p5)
	bar(p)

	// Funciones anonimas
	func(x int) {
		fmt.Println("El significado de la vida es:", x)
	}(42)

	//Callbacks
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := sum(ii...)
	fmt.Println("all numbers", s)

	s2 := even(sum, ii...)
	fmt.Println("even numbers", s2)
}

func practicaEstructuras() {
	p1 := persona{
		nombre:   "Gustavo",
		apellido: "Garcia",
		age:      22,
	}
	p2 := persona{
		nombre:   "Pepe",
		apellido: "Pollo",
		age:      30,
	}
	p3 := agenteSecreto{
		persona: persona{
			nombre:   "Mariana",
			apellido: "Padilla",
			age:      22,
		},
		lpm: true,
	}
	//Struct anonima
	p4 := struct {
		first string
		last  string
		age   int
	}{
		first: "Juan",
		last:  "Perez",
		age:   33,
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)
	fmt.Println(p4)

	fmt.Println(p1.nombre, p1.apellido, p1.age)
	fmt.Println(p2.nombre, p2.apellido, p2.age)
	fmt.Println(p3.nombre, p3.apellido, p3.age, p3.lpm)
}

// Para funciones con parametros variables se escriben de la siguiente manera de la siguiente manera de la siguiente
func foo(x ...int) int {
	//convierte los datos en un slice de enteros
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	suma := 0
	for i, v := range x {
		suma += v
		fmt.Println("El valor en el indice", i, "suma", v, "al total, quedando", suma)
	}
	fmt.Println("El total es", suma)
	return suma
}

// Metodos = func (r receptor) identificador (parametros) (retornos) {codigo}
func (a agenteSecreto) presentarse() {
	fmt.Println("Hola, soy", a.nombre, a.apellido, "y tengo", a.age, "anios. Soy un agente secreto")
}

func (p persona) presentarse() {
	fmt.Println("Hola, soy humano, y me llamo", p.nombre, p.apellido, "y tengo", p.age, "anios")
}

func bar(h humano) {
	switch h.(type) {
	case persona:
		fmt.Println("Fui pasado a al funcion bar (persona)", h.(persona).nombre)
	case agenteSecreto:
		fmt.Println("Fui pasado a al funcion bar (agente secreto)", h.(agenteSecreto).nombre)
	}

}

func sum(xi ...int) int {
	fmt.Printf("%T\n", xi)
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

//Callbacks
func even(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}
