package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Sistema operativo:", runtime.GOOS)
	fmt.Println("Arquitectura:", runtime.GOARCH)
}
