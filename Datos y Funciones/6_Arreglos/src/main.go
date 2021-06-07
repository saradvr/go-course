package main

import "fmt"

func main() {
	// var numeros [5]int // el 5 es la longitud y no recibe ningún tipo diferente

	// numeros[0] = 100
	// numeros[4] = 300

	// Inicializar el arreglo con valores
	// numeros := [5]int{100, 200, 300, 400, 500}
	// numeros := [...]int{100, 200, 300, 400, 500} //No se define la cantidad de elementos

	// fmt.Println(numeros[4])

	// Arreglos de String
	// monedas := [...]string{"Peso Mexicano", "Dólar", "Euro"} //Quedan automáticamente en esos índices
	// monedas := [...]string{0: "Peso Mexicano", 1: "Dólar", 2: "Euro"} //Queda en orden
	monedas := [...]string{0: "Peso Mexicano", 2: "Dólar", 1: "Euro"} //Puedo cambiar el orden de los índices, si no pongo un índice le pone un string vacío. Si quiero varios indices vacíos, asigno un valor al último índice que quiero tener.

	subMonedas := monedas[0:2] // Esto es un slice

	fmt.Println(monedas, subMonedas)
}
