package main

import "fmt"

type Animal interface { // es un tipo, con métodos como comportamiento
	Comer() // Métodos
	Dormir()
	Cazar()
}

type Perro struct {
	Nombre string
}

func (self Perro) Comer() {
	fmt.Println("El perro come")
}

func (self Perro) Dormir() {
	fmt.Println("El perro duerme")
}

func (self Perro) Cazar() {
	fmt.Println("El perro caza")
}

func ejecutarAcciones(animal Animal) {
	animal.Comer()
	animal.Dormir()
	animal.Cazar()
}

func main() {

	loki := Perro{Nombre: "Loki"}

	ejecutarAcciones(loki)

	fmt.Println(loki)

}
