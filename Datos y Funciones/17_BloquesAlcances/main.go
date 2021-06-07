package main

import "fmt"

func main() { // Bloque 1

	// Bloque: colección de sentencias agrupadas en un juego de llaves

	var curso = "Curso profesional de Go!" // puede usarse en todos los bloques inferiores
	var version = 10

	{ //Bloque 2
		fmt.Println(curso)
		fmt.Println(version)

		{ // Bloque 3
			var version = 5 // Sólo existe dentro del bloque 3, puede usarse en los superiores.
			fmt.Println("Bloque 3:", version)
			fmt.Println(curso)
		}
	}

	fmt.Println(curso)

}
