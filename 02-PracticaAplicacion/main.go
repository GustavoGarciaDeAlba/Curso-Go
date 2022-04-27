package main

import (
	"encoding/json"
	"fmt"
	"sort"

	"golang.org/x/crypto/bcrypt"
)

type persona struct {
	Nombre   string `json:"Nombre"`
	Apellido string `json:"Apellido"`
	Edad     int    `json:"Edad"`
}

type PorEdad []persona

func (a PorEdad) Len() int           { return len(a) }
func (a PorEdad) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a PorEdad) Less(i, j int) bool { return a[i].Edad < a[j].Edad }

func main() {

	fmt.Print("\n=================Practica Punteros=================\n")
	// Punteros
	p1 := persona{Nombre: "Pepe"}
	fmt.Println("Nombre antes: ", p1.Nombre)
	cambiame(&p1)
	fmt.Println("Nombre despues: ", p1.Nombre)

	fmt.Print("\n=================Practica JSON Marshal=================\n")
	//Practica Marshal
	practicaMarshal()

	fmt.Print("\n=================Practica JSON Unmarshal=================\n")
	//Practica Unmarshal
	practicaUnmarshal()

	fmt.Print("\n=================Practica Sort=================\n")
	s := []int{5, 2, 6, 3, 1, 4}
	fmt.Println(s)
	sort.Ints(s)
	fmt.Println(s)
	xs := []string{"James", "Q", "M", "Moneypenny", "Dr. No"}
	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)

	fmt.Print("\n=================Practica Sort Personalizado=================\n")
	p2 := persona{"Moneypenny", "A", 27}
	p3 := persona{"Q", "B", 64}
	p4 := persona{"M", "C", 56}
	personas := []persona{p1, p2, p3, p4}
	fmt.Println(personas)
	sort.Sort(PorEdad(personas))
	fmt.Println(personas)

	fmt.Print("\n=================Practica Bcrypt=================\n")
	password := `clave123`
	bs, err := bcrypt.GenerateFromPassword([]byte(password), 4)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(password)
	fmt.Println(bs)
}

//Funcion para repasar punteros
func cambiame(p *persona) {
	(*p).Nombre = "Gustavo"
}

//Practica Marshal
func practicaMarshal() {
	p1 := persona{
		Nombre:   "James",
		Apellido: "Bond",
		Edad:     32,
	}

	p2 := persona{
		Nombre:   "Miss",
		Apellido: "Moneypenny",
		Edad:     27,
	}

	personas := []persona{p1, p2}

	fmt.Println(personas)

	bs, err := json.Marshal(personas) // de objeto a JSON
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}

// Practica Unmarshal
func practicaUnmarshal() {
	s := `[{"Nombre":"James","Apellido":"Bond","Edad":32},{"Nombre":"Miss","Apellido":"Moneypenny","Edad":27}]`
	bs := []byte(s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	var personas []persona
	err := json.Unmarshal(bs, &personas)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Toda la data", personas)

	for i, v := range personas {
		fmt.Println("\nPERSONA NÚMERO", i+1)
		fmt.Println(v.Nombre, v.Apellido, v.Edad)
	}
}
