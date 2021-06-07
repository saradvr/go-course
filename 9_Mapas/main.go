package main

import "fmt"

func main() {

	// El mapa es una colección desordenada de elementos

	dias := make(map[int]string) // Para el map, dentro del paréntesis el tipo de dato de la llave, y afuera el tipo del valor

	dias[0] = "Domingo" // En este caso, como es un mapa, el valor entre paréntesis no es un índice sino una llave
	dias[1] = "Lunes"   // Si la llave no existe la crea, si ya existe, la modifica
	dias[2] = "Martes"
	dias[3] = "Miércoles"
	dias[4] = "Jueves"
	dias[5] = "Viernes"
	dias[6] = "Sábado"

	// fmt.Println(dias) // map[0:Domingo 1:Lunes 2:Martes 3:Miércoles 4:Jueves 5:Viernes 6:Sábado]
	// fmt.Println(dias[4]) // Jueves

	dias[4] = "JUEVES"
	// fmt.Println(dias[4]) //JUEVES

	delete(dias, 4)
	fmt.Println(dias[4]) // No imprime nada porque se eliminó
	fmt.Println(dias)    // map[0:Domingo 1:Lunes 2:Martes 3:Miércoles 5:Viernes 6:Sábado]

	// MAPAS MÁS COMPLEJOS

	usuarios := make(map[string][]int)

	usuarios["Eduardo"] = []int{10, 9, 8, 10, 5}
	usuarios["Sara"] = []int{10, 10, 7, 9, 8}

	fmt.Println(usuarios)

	// ITERACIONES

	users := map[int]string{} // Otra forma para crear mapas
	users[1] = "Usuario 1"
	users[2] = "Usuario 2"
	users[3] = "Usuario 3"
	users[4] = "Usuario 4"

	for llave, valor := range users {
		fmt.Println(llave, valor)
	}

	for id, username := range users { // pueden cambiarse los nombres de llave, valor
		fmt.Println(id, username)
	}

}
