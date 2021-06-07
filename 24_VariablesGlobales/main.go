package main

import "fmt"

var username string

func funcion1() {
	username = "Función 1"
	fmt.Println("Función 1:", username)
}

func funcion2() {
	username = "Función 2"
	fmt.Println("Función 2:", username)
}

func main() {
	username = "Cody"
	funcion1()
	funcion2()

	fmt.Println(username) //Toma el último valor según lo haya cambiado
}
