package main

import "fmt"

func main() {

	// ESTRUCTURA
	// for <variable>; <condicion> ; <incremento> {
	// 	// Sentencias
	// }

	// for i := 1; i <= 100; i++ {

	// 	if i%2 == 0 {
	// 		fmt.Println(i)
	// 	}

	// }

	// FOR COMO WHILE
	numero := 123
	contador := 0

	for numero > 0 {

		numero = numero / 10
		contador++

	}

	fmt.Println("La cantidad de dígitos es:", contador)

	// CICLO FOREACH
	animales := [...]string{"Perro", "Gato", "Pez", "Vaca", "Cerdo"}

	for indice := 0; indice < len(animales); indice++ {

		elemento := animales[indice]
		fmt.Println(elemento)

	}

	for indice, elemento := range animales {
		fmt.Println("El índice es:", indice, "el valor es", elemento)
	}

	for _, elemento := range animales { // Si no quiero usar alguno de los dos debo poner _
		fmt.Println("El valor es", elemento)
	}

	// BREAK - CONTINUE

	for i := 1; i <= 10; i++ {

		if i == 5 {
			continue // finaliza la iteración actual para pasar a la siguiente, entonces no alcanza a imprimirse en consola
		}

		if i == 8 {
			break // finaliza todo el ciclo for ahí
		}

		fmt.Println(i)
	}

}
