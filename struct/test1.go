package main

import "fmt"

func main() {

	type Person struct {
		name string
		age  int
	}

	var person Person
	fmt.Println(person)

	person.name = "penxiao"
	person.age = 31
	fmt.Println(person)

	person1 := new(Person)
	fmt.Println(person1)
	person1.name = "penxiao"
	fmt.Println(person1)

	person2 := Person{name: "penxiao", age: 31}
	fmt.Println(person2)

}
