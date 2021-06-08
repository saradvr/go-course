package main

import (
	"fmt"
	"paquetes/testPackage"
)

func main() {

	curso := testPackage.Curso{Titulo: "Curso Profesional de GO"}

	fmt.Println(curso)
	fmt.Println(curso.GetTitulo())

}
