package main

import "fmt"

func mostrarVariable(objeto interface{}) {
	fmt.Printf("El valor de la variable es: %v\n", objeto)
}

func suma(num1, num2 int) int {
	return num1 + num2
}

type User struct {
	Nombre string
}

func main() {

	//fmt.Println() // recibe cualquier número y tipo de objeto
	mostrarVariable(10)
	mostrarVariable("String")
	mostrarVariable(true)
	mostrarVariable(suma)

	usuario := User{Nombre: "Eduardo"}
	mostrarVariable(usuario)

}
