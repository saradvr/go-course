package main

import "fmt"

type User struct {
	Name  string
	Email string
}

func (self *User) setName(name string) { // Así se define un método, se usa un puntero
	self.Name = name
}

func (self *User) getName() string {
	return self.Name
}

func main() {
	cody := User{Name: "Cody", Email: "info@mail.com"}
	fmt.Println(cody)

	cody.setName("Sara")

	fmt.Println(cody)
	fmt.Println(cody.getName())
}
