package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var contador int64

	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&contador, 1)
			fmt.Println(atomic.LoadInt64(&contador))
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("El valor final de incremento es:", contador)
}
