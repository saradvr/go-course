package main

import "fmt"

func factorial(numero int) int {

	if numero == 1 {
		return 1
	}

	return factorial(numero-1) * numero
}

type customFunction func(n int) int

func main() {

	//FORMA 1
	// var miFuncion = factorial

	//FORMA 2
	// var miFuncion func(n int) int // La declaro con su tipo, pero esta no es la mejor forma

	// if miFuncion == nil {
	// 	miFuncion = factorial // Valor por defecto
	// }

	//FORMA 3
	var miFuncion customFunction
	if miFuncion == nil {
		miFuncion = factorial // Valor por defecto
	}

	resultado := miFuncion(3)

	fmt.Println(resultado)

}
