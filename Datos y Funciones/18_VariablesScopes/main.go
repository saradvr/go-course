package main

import "fmt"

func modificarVariable(parametro string) {
	fmt.Println("Valor actual:", parametro)
	parametro = "Nuevo curso de Go"
	fmt.Println("Nuevo valor:", parametro)

	fmt.Printf("%p\n", &parametro) //0xc000010250 // Posición en memoria
}

func main() {
	var curso = "Curso profesional de Go"

	modificarVariable(curso) // Hace una copia, pero son copias diferentes según el bloque

	fmt.Println(">>>", curso)
	fmt.Printf("%p\n", &curso) //0xc000010240
}
