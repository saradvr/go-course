package main

import "fmt"

func promedio(calificaciones ...int) int { //Recibe cualquier cantidad de elementos, y para eso se usan los ... el nombre es el del arreglo que tendrá todos los argumentos

	var promedio, suma int

	for _, calificacion := range calificaciones {
		suma += calificacion
	}

	promedio = suma / len(calificaciones)

	return promedio
}

func main() {
	// Variadic Function: n cantidad de elementos que deseemos
	fmt.Println("Hola", "Mundo", "Desde el curso", "Profesional")

	resultado := promedio(10)

	fmt.Println("Promedio:", resultado)
}
