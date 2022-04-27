package main

import "fmt"

type persona struct {
	nombre string
}
type humano interface {
	hablar()
}

func main() {
	p := persona{
		nombre: "Gustavo",
	}
	diAlgo(&p)
}

func diAlgo(h humano) {
	h.hablar()
}

func (p *persona) hablar() {
	fmt.Println("Holaaaaaaa", p.nombre)
}
