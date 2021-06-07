package main

import "fmt"

func main() {

	//panic finaliza la ejecución del programa

	var dividendo, divisor int

	fmt.Print("Ingresa un valor para el dividendo: ")
	fmt.Scanf("%d", &dividendo)

	fmt.Print("Ingresa un valor para el divisor: ")
	fmt.Scanf("%d", &divisor)

	if divisor == 0 {
		panic("El divisor no puede ser 0. No es posible dividir sobre 0.")
	}

	resultado := dividendo / divisor

	fmt.Println(resultado)

}
