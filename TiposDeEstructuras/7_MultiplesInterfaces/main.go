package main

import "fmt"

type Animal interface {
	Dormir()
}

type Mascota interface {
	Comer()
}

type Gato struct {
	Nombre string
}

func (self Gato) Dormir() {
	fmt.Println("El gato duerme")
}

func (self Gato) Comer() {
	fmt.Println("El gato come")
}

func funcionParaUnAnimal(animal Animal) {
	fmt.Println("El objeto es un animal")
}

func funcionParaUnaMascota(mascota Mascota) {
	fmt.Println("El objeto es una mascota")
}

func main() {

	gato := Gato{Nombre: "Lupe"}

	gato.Comer()
	gato.Dormir()

	funcionParaUnAnimal(gato)
	funcionParaUnaMascota(gato)

}
