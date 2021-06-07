package main

import "fmt"

func greet(username string) {
	fmt.Println("Hola", username)
}

func suma(num1, num2 int) (int, string) { // si ambos parámetros son del mismo tipo, sólo debo ponerlo al final. Después del () pongo el tipo de dato que voy a retornar, si retorna varios los pongo entre ()
	return num1 + num2, "Este es un mensaje desde la función suma"
}

// FUNCIONES COMO ARGUMENTO

type Operacion func(balance, cantidad int) int // Cómo crear nuevos tipos

func sum(num1, num2 int) int { // es de tipo operacion porque cumple las condiciones
	return num1 + num2
}

func subs(num1, num2 int) int { // es de tipo operacion porque cumple las condiciones
	return num1 - num2
}

func crearOperacion(tipo string) Operacion {
	if tipo == "suma" {
		return func(balance, cantidad int) int { return balance + cantidad }
	} else if tipo == "resta" {
		return func(balance, cantidad int) int { return balance - cantidad }
	} else {
		return func(balance, cantidad int) int { return balance * cantidad }
	}
}

func ejecutarOperacion(funcion Operacion, balance, cantidad int) {

	fmt.Println("Antes de la operacion")

	resultado := funcion(balance, cantidad)
	fmt.Println("El resultado es:", resultado)

	fmt.Println("Después de la operación")

}

func main() {
	greet("Sara")
	resultadoSuma, mensajeSuma := suma(1, 5) // si sólo quiero uno de los valores le pongo _ al otro
	fmt.Println(resultadoSuma, mensajeSuma)

	// FUNCIONES ANÓNIMAS
	fmt.Println("\n\n FUNCIONES ANÓNIMAS")
	func() {
		fmt.Println("Hola mundo, desde una función sin nombre")
	}()

	miFuncion := func(username string) { // Puedo almacenar la función en una variable
		fmt.Println("Hola,", username, "esta es miFuncion")
	}
	miFuncion("Sara")

	otraFuncion := func(username string) string {
		message := fmt.Sprintf("Hola %s, te saludamos desde una función sin nombre", username)
		return message
	}
	mensajeFinal := otraFuncion("Natalia")
	fmt.Println(mensajeFinal)

	// FUNCIONES COMO ARGUMENTO
	fmt.Println("\n\n FUNCIONES COMO ARGUMENTO")
	ejecutarOperacion(subs, 50, 25)
	ejecutarOperacion(sum, 50, 25)

	// RETORNAR FUNCIONES
	fmt.Println("\n\n RETORNAR FUNCIONES")
	functionSum := crearOperacion("suma")
	resultadoSum := functionSum(40, 50)
	fmt.Println("Suma", resultadoSum)
	functionResta := crearOperacion("resta")
	resultadoResta := functionResta(40, 50)
	fmt.Println("Resta:", resultadoResta)
	functionDefault := crearOperacion("fgewrgegew")
	resultadoDefault := functionDefault(40, 50)
	fmt.Println("Default:", resultadoDefault)

}
