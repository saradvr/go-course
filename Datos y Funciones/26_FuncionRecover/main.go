package main

import (
	"fmt"
	"os"
)

func main() {

	defer func() {
		if err := recover(); err != nil { // No va a mostrar el error sino que ejecuta otra cosa que yo desee
			fmt.Println("Ups, al parecer el programa no finalizó de forma correcta:", err)
		}
	}()

	if file, err := os.Open("example.txt"); err != nil {

		panic("No fue posible obtener el archivo")

	} else {

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

}
