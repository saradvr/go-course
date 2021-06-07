package main

import "fmt"

func main() {

	// slice := make([]int, 0, 3) // Primero un arreglo que es la base de todo slice, número de elementos del slice, capacidad máxima
	//Imprime: []

	slice := make([]int, 3, 3) // imprime: [0 0 0] si no cambio sus valores

	slice[0] = 100
	slice[1] = 200
	slice[2] = 300

	slice = append(slice, 400) // Al agregar un elemento más por encima de la capacidad, crea un nuevo elemento en memoria con el doble de capacidad

	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
}
