package main

import "fmt"

// Person is a struct
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// GetName is a method
func (p Person) GetName() {
	fmt.Println(p.FirstName, p.LastName)
}

func main() {

	p1 := Person{FirstName: "Peng", LastName: "Xiao", Age: 31}
	p1.GetName()

}
