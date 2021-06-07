package main

import "fmt"

func main() {
	if nombre, edad := "Cody", 7; nombre == "Cody" {
		fmt.Println("Hola", nombre, "te damos la bienvenida al curso de Go!")
	} else {
		fmt.Println("Los valores son: ", nombre, edad)
	}

	// fmt.Println(nombre) // Las variables quedan definidas dentro del if y no pueden usarse fuera. Sólo en el if o el else.
}
