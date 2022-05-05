//Package decir nos permite saludar
package decir

import "fmt"

// Saludar nos deja saludar a una persona
func Saludar(s string) string {
	return fmt.Sprint("Bienvenido ", s)
}
