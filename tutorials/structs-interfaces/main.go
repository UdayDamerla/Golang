package main

import (
	"fmt"
)

type PersonData struct {
	Name string
	Age  int
}

// A constructor-style helper keeps initialization tidy.
func NewPerson(name string, age int) PersonData {
	return PersonData{Name: name, Age: age}
}

// Pointer receivers can access and modify struct state.
func (e *PersonData) GetName() string {
	return e.Name
}

func (e *PersonData) GetAge() int {
	return e.Age
}

// Interfaces in Go are satisfied implicitly.
type Person interface {
	GetName() string
	GetAge() int
}

// Embedding promotes fields and methods from PersonData.
type Employee struct {
	PersonData
	Role string
}

func (e *Employee) GetName() string {
	return e.PersonData.Name + " (" + e.Role + ")"
}

// Interface values enable polymorphic behavior.
func printPersonDetails(p Person) {
	fmt.Printf("Name: %s | Age: %d\n", p.GetName(), p.GetAge())
}

func main() {
	myPerson := NewPerson("John", 30)
	myEmployee := Employee{
		PersonData: NewPerson("Alice", 28),
		Role:       "Developer",
	}

	fmt.Println("Direct struct access:", myPerson.Name, myPerson.Age)

	var myPersonInterface Person = &myPerson
	printPersonDetails(myPersonInterface)

	var myEmployeeInterface Person = &myEmployee
	printPersonDetails(myEmployeeInterface)
}
