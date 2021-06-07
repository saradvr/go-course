package main

import "fmt"

func main() {

	//SLICE PARTE 1
	// Un slice no puede existir sin un arreglo
	// numeros := []int{1, 2, 3, 4}

	// numeros = append(numeros, 5) // Append devuelve un nuevo objeto
	// numeros = append(numeros, 6)
	// numeros = append(numeros, 7)
	// numeros = append(numeros, 8)
	// numeros = append(numeros, 9)
	// numeros = append(numeros, 10)

	// nuevoSlice := numeros[0:5] // ES UNA REFERENCIA A UNA PORCIÓN DE UN ARREGLO
	// tercerSlice := numeros[0:3]

	// numeros[0] = 100 // Al cambiar esto también cambia en el nuevoSlice

	// fmt.Println(numeros, nuevoSlice, tercerSlice)

	// SLICE PARTE 2

	meses := []string{"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio", "Julio", "Agosto", "Septiembre"}

	// Un puntero
	// Una longitud
	// Una capacidad

	longitud := len(meses)
	capacidad := cap(meses)

	fmt.Printf("La longitud es: %v, La capacidad es %v %p\n", longitud, capacidad, meses) // La longitud es: 9, La capacidad es 9 0xc000138000

	meses = append(meses, "Octubre") //Si la estructura se encuentra en su capacidad máxima, se genera un nuevo arreglo con una posición en memoria diferente
	meses = append(meses, "Noviembre")
	meses = append(meses, "Diciembre")

	longitud = len(meses)
	capacidad = cap(meses)

	fmt.Printf("La longitud es: %v, La capacidad es %v %p\n", longitud, capacidad, meses) // La longitud es: 10, La capacidad es 18 0xc00013e000

}
