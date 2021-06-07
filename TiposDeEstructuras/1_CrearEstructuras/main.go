package main

import "fmt"

type User struct {
	Name  string
	Email string
}

func main() {

	// FORMA 1

	// var cody User // Un objeto

	// cody.Name = "Cody"
	// cody.Email = "info@mail.com"

	// FORMA 2
	// cody := User{Name: "Cody", Email: "info@mail.com"}

	// FORMA 3
	Name := "Cody"
	Email := "info@mail.com"

	cody := User{Name, Email}

	fmt.Println(cody)

}
