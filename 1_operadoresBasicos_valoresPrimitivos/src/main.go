package main

import "fmt"

func main() {

	// Declaración de constantes
	const pi float64 = 3.14
	const pi2 = 3.1415
	fmt.Println("pi", pi)
	fmt.Println("pi2", pi2)

	// Declaración de variables enteras
	base := 12 //: significan que no ha sido creada anteriormente, si no le pongo los : significa que ya la cree anteriormente
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	// var nombre string
	// var apellido string
	// var pais string

	// var nombre, apellido, pais string
	// var nombre, apellido, pais = "Eduardo", "Garcia", "México"
	nombre, apellido, pais := "Eduardo", "Garcia", "México"

	fmt.Println(nombre, apellido, pais)

	// Zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	// Area Cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area cuadrado: ", areaCuadrado)

	x := 10
	y := 50

	// Suma
	result := x + y
	fmt.Println("Suma: ", result)

	// Resta
	result = y - x
	fmt.Println("Resta: ", result)

	// Multiplicación
	result = x * y
	fmt.Println("Multiplicación: ", result)

	// División
	result = y / x
	fmt.Println("División: ", result)

	// Módulo
	result = y % x
	fmt.Println("Residuo: ", result)

	// Incremental
	x++
	fmt.Println("Incremental: ", x)

	// Decremental
	x--
	fmt.Println("Decremental: ", x)

	// Valores primitivos
	// int (enteros)
	// uint (enteros positivos)
	// float32 o float64 (decimales)
	// string (texto, siempre con comillas dobles)
	// bool (true o false)
	// Complex64 o Complex128 (números con parte real e imaginaria, c := 10 + 8i)
}
