package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("Número de CPUs:", runtime.NumCPU())
	fmt.Println("Número de Goroutinas:", runtime.NumGoroutine())

	contador := 0

	const gs = 10
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := contador
			runtime.Gosched()
			v++
			contador = v
			wg.Done()
		}()
		fmt.Println("Número de Goroutinas:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Valor final de contador:", contador)
}
