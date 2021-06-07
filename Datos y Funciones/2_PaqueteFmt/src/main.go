package main

import "fmt"

func main() {
	// Declaración de variables
	helloMessage := "Hello"
	worldMessage := "world"

	//Println -> Agrega el salto de línea de forma automática
	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)

	// Printf -> Interpolar datos
	nombre := "Plazti"
	cursos := 500
	fmt.Printf("%s tiene más de %d cursos\n", nombre, cursos) //si se sabe el tipo de dato se le agrega la letra del tipo
	fmt.Printf("%v tiene más de %v cursos\n", nombre, cursos) // cuando no sé cómo será el tipo de dato

	// Sprintf
	message := fmt.Sprintf("%s tiene más de %d cursos", nombre, cursos) //Todo lo que genere ese String lo guarda en message
	fmt.Println(message)

	// Tipo de datos
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("cursos: %T", cursos)

}
