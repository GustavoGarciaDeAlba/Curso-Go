package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c) // funcion para cerrar canales
	}()

	//recibir
	for v := range c { // Con un for se pueden obtener todos los valores de un canal
		fmt.Println(v)
	}

	fmt.Println("Finalizando...")

}
