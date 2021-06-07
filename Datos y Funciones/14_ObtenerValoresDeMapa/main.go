package main

import "fmt"

func main() {

	usuarios := make(map[string]string)
	usuarios["Eduardo"] = "eduardo@codigofacilito.com"

	// correo := usuarios["Eduardo"] // eduardo@codigofacilito.com
	// // correo2 := usuarios["Cody"] // Imprime vacío

	// if correo != "" { // en los casos en que sí se haya almacenado una cadena vacía no debería retornar que no fue posible obtener el valor, porque ese es el valor
	// 	fmt.Println(correo)
	// } else {
	// 	fmt.Println("No fue posible obtener el valor.")
	// }

	// Seleccionar una llave de mapas devuelve dos valores:
	// correo, ok := usuarios["Eduardo"] // ok es un bool que indica si pudo obtenerse el valor

	// if ok {
	// 	fmt.Println(correo)
	// } else {
	// 	fmt.Println("No fue posible obtener el valor.")
	// }

	// ESTRUCTURA RECOMENDADA PARA OBTENER VALOR DE UN MAPA
	if correo, ok := usuarios["cody"]; ok {
		fmt.Println(correo)
	} else {
		fmt.Println("No fue posible obtener el valor.")
	}
}
