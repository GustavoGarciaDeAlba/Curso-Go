package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	wg.Add(1)

	go foo()
	bar()

	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	wg.Wait()

	//Retornar una variable al usar una go routine
	ch := make(chan int) //Creacion de un canal
	go func() {
		ch <- doSomething(5)
	}()
	fmt.Println(<-ch)

	//Condicion de carrera
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	contador := 0

	const gs = 100
	var waitG sync.WaitGroup
	waitG.Add(gs)
	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := contador
			v++
			runtime.Gosched()
			contador = v
			mu.Unlock()
			waitG.Done()
		}()
		fmt.Println("Goroutines\t", runtime.NumGoroutine())
	}
	waitG.Wait()
	fmt.Println("Cuenta:", contador)
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}

func doSomething(x int) int {
	return x * 5
}
