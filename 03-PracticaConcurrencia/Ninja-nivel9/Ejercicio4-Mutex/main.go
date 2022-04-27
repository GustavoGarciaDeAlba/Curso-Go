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
	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := contador
			runtime.Gosched()
			v++
			contador = v
			fmt.Println("Valor de contador:", contador)
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Valor final de contador:", contador)
}
