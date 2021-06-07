package main

import (
	"fmt"
	"os"
)

func funcion1() {
	fmt.Println("Hola, soy funcion 1")
}

func funcion2() {
	fmt.Println("Hola, soy funcion 2")
}

func main() {

	// EJEMPLO BÁSICO
	// fmt.Println("Hola, estamos en la función Main")
	// defer funcion1() // se ejecuta cuando el bloque de la función main finalice. Sólo se ejecuta cuando la función que la llama termina
	// funcion2()

	file, err := os.Open("example.txt")

	if err != nil {
		panic("No fue posible obtener el archivo")
	}

	defer func() {
		fmt.Println("Cerramos el archivo!!")
		file.Close()
	}()

	contenido := make([]byte, 254)

	longitud, _ := file.Read(contenido) //el _ es err

	fmt.Println(longitud)

	fmt.Println(contenido[0:8])

	contenidoArchivo := string(contenido[0:longitud]) //Convierte a texto todo el contenido

	fmt.Println(contenido)

	fmt.Println(contenidoArchivo)

}
