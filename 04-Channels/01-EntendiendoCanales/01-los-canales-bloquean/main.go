package main

import "fmt"

func main() {
	//unbuffered channel (canal sin bufer) - con un canal sin buffer siempre se necesita una gorutina emisora y otra receptora
	ca := make(chan int)
	ca <- 42
	fmt.Println(<-ca)
}
