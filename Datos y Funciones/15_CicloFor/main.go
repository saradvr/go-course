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
	fmt.Println("\n\n FOR COMO WHILE")
	numero := 123
	contador := 0

	for numero > 0 {

		numero = numero / 10
		contador++

	}

	fmt.Println("La cantidad de dígitos es:", contador)

	// CICLO FOREACH
	fmt.Println("\n\n FOREACH")
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
	fmt.Println("\n\n BREAK - CONTINUE")

	for i := 1; i <= 10; i++ {

		if i == 5 {
			fmt.Println("continue")
			continue // finaliza la iteración actual para pasar a la siguiente, entonces no alcanza a imprimirse en consola
		}

		if i == 8 {
			fmt.Println("break")
			break // finaliza todo el ciclo for ahí
		}

		fmt.Println(i)
	}

	// FOR COMO DO WHILE
	fmt.Println("\n\n FOR COMO DO WHILE")

	var num = 10

	for num < 10 { // While -> (no funciona porque de entrada no se cumpliría la condición)
		fmt.Println(num)
		num++
	}

	for ok := true; ok; ok = num < 10 { // DO while: cuando necesito que se ejecute al menos una vez
		fmt.Println(num)
		num++
	}

	// FOR COMO CICLO INFINITO
	fmt.Println("\n\n FOR COMO CICLO INFINITO")

	numero2 := 1

	for {
		fmt.Println(numero2)
		numero2++

		if numero2 == 100 {
			// break // rompe el ciclo infinito
			panic("Fin del ciclo for") // Detiene abruptamente el programa con un error.
		}
	}

}
