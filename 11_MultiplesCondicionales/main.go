package main

import "fmt"

func main() {

	// else if
	var calificacion int

	fmt.Print("Ingresa una calificación: ")
	fmt.Scanf("%d", &calificacion)

	// if calificacion == 10 {
	// 	fmt.Println("Felicidades, calificación perfecta")
	// } else {
	// 	if calificacion == 8 || calificacion == 9 {
	// 		fmt.Println("Aprobaste la materia")
	// 	} else {
	// 		if calificacion == 6 || calificacion == 7 {
	// 			fmt.Println("Aprobaste la materia, pero necesitas estudiar más")
	// 		} else {
	// 			if calificacion >= 0 && calificacion <= 5 {
	// 				fmt.Println("No aprobaste la materia")
	// 			} else {
	// 				fmt.Println("Calificación no válida")
	// 			}
	// 		}
	// 	}
	// }

	if calificacion == 10 {
		fmt.Println("Felicidades, calificación perfecta")
	} else if calificacion == 8 || calificacion == 9 {
		fmt.Println("Aprobaste la materia")
	} else if calificacion == 6 || calificacion == 7 {
		fmt.Println("Aprobaste la materia, pero necesitas estudiar más")
	} else if calificacion >= 0 && calificacion <= 5 {
		fmt.Println("No aprobaste la materia")
	} else {
		fmt.Println("Calificación no válida")
	}

}
