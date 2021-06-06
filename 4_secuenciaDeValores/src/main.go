package main

import "fmt"

// const Domingo int = 0
// const Lunes int = 1
// const Martes int = 2
// const Miercoles int = 3
// const Jueves int = 4
// const Viernes int = 5
// const Sabado int = 6

// const (
// 	Domingo int = 0
// 	Lunes int = 1
// 	Martes int = 2
// 	Miercoles int = 3
// 	Jueves int = 4
// 	Viernes int = 5
// 	Sabado int = 6
// )

const (
	Domingo int = iota //empieza por defecto en 0, le puedo sumar el número en que quiero comenzar
	Lunes
	Martes
	Miercoles
	Jueves
	Viernes
	Sabado
)

func main () {
	fmt.Println(Domingo)
	fmt.Println(Lunes)
	fmt.Println(Martes)
	fmt.Println(Miercoles)
	fmt.Println(Jueves)
	fmt.Println(Viernes)
	fmt.Println(Sabado)
}