package decir

import (
	"fmt"
	"testing"
)

func TestSaludar(t *testing.T) {
	s := Saludar("Gustavo")
	if s != "Bienvenido Gustavo" {
		t.Error("Expected", "Bienvenido Gustavo", "Got", s)
	}
}

func ExampleSaludar() {
	fmt.Println(Saludar("Gustavo"))
	//Output:
	// Bienvenido Gustavo
}

func BenchmarkSaludar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Saludar("Gustavo")
	}
}
