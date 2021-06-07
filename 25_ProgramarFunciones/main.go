package main

import "fmt"

func funcion1() {
	fmt.Println("Hola, soy funcion 1")
}

func funcion2() {
	fmt.Println("Hola, soy funcion 2")
}

func main() {

	fmt.Println("Hola, estamos en la función Main")
	defer funcion1() // se ejecuta cuando el bloque de la función main finalice. Sólo se ejecuta cuando la función que la llama termina
	funcion2()
}
