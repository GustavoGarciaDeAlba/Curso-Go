package main

import "fmt"

func main() {
	// Canal bidireccional
	ca := make(chan int, 2) //Para hacer un canal send only es: chan<- int. Para hacer lo receive only es: <-chan int.

	ca <- 42 //Aqui se envia
	ca <- 43

	fmt.Println(<-ca) // Aqui se recibe
	fmt.Println(<-ca)
	fmt.Printf("%T", ca)
	//Solamente se pueden asignar canales de general a especifico
}
