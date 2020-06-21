package main

import (
	"fmt"
	"strconv"
)
// good for ref https://blog.wu-boy.com/2017/05/go-struct-method-pointer-or-value/
	// Define Person struct
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int

	firstName, lastName, city, gender string
	age                               int
}

func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}
// hasBirthday method (pointer reciever)
func (p *Person) hasBirthday() {
	p.age++
}
// getMarried (pointer reciever)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
		return
	}
}

func main() {
	// Init person using struct
	person1 := Person{firstName: "Samantha", lastName: "Smith", city: "Boston", gender: "f", age: 25}
	// Alternative
	person2 := Person{"Bob", "Johnson", "New York", "m", 30}

	// fmt.Println(person1.firstName)
	// person1.age++
	// fmt.Println(person1)

	person1.hasBirthday()
	person1.hasBirthday()
	person1.hasBirthday()
	person1.getMarried("Williams")

	person2.getMarried("Thompson")

	fmt.Println(person1.greet())
}