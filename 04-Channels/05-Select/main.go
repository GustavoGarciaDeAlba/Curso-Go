package main

import "fmt"

func main() {
	par := make(chan int)
	impar := make(chan int)
	salir := make(chan int)

	//Enviar
	go enviar(par, impar, salir)

	//Recibir
	recibir(par, impar, salir)

	fmt.Println("Finalizando...")
}

func enviar(par, impar, salir chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			par <- i
		} else {
			impar <- i
		}
	}
	salir <- 0
}
func recibir(par, impar, salir <-chan int) {
	for {
		select {
		case v := <-par:
			fmt.Println("Desde el canal par:", v)
		case v := <-impar:
			fmt.Println("Desde el canal impar:", v)
		case v := <-salir:
			fmt.Println("Desde el canal salir:", v)
			return
		}
	}
}
