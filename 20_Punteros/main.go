package main

import "fmt"

func modificarValor(parametro *string) { // modifica una referencia y no una copia
	*parametro = "Nuevo curso de Go"
}

func main() {
	var curso = "Curso profesional de Go"

	fmt.Println("Antes de la función", curso)

	modificarValor(&curso) //así se pasa una referencia, un espacio en memoria

	fmt.Println("Después de la función", curso)
}
