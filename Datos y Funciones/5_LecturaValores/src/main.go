package main

import "fmt"

func main() {
	var nombre string
	var edad int
	var altura float64

	fmt.Print("Ingresa tu nombre: ")
	fmt.Scanf("%s", &nombre)

	fmt.Print("Ingresa tu edad: ")
	fmt.Scanf("%d", &edad)

	fmt.Print("Ingresa tu altura: ")
	fmt.Scanf("%f", &altura)

	fmt.Printf("Hola %s con una edad %d y una altura %.2f\n", nombre, edad, altura)
}
