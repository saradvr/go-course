package main

import "fmt"

// Función recursiva es la que se llama a sí misma

func factorial(numero int) int {

	if numero == 1 {
		return 1
	}

	return factorial(numero-1) * numero
}

func main() {
	resultado := factorial(3)

	fmt.Println("El factorial es:", resultado)
}
