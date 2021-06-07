package main

import "fmt"

func main() {

	// ESTRUCTURA
	// for <variable>; <condicion> ; <incremento> {
	// 	// Sentencias
	// }

	for i := 1; i <= 100; i++ {

		if i%2 == 0 {
			fmt.Println(i)
		}

	}

}
