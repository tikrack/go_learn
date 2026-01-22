package main

import "fmt"

type Person struct {
	Name   string
	Family string
	Age    int
}

type PersonBuilder struct {
	Person
}

func main() {
	builder := PersonBuilder{}

	builder.setName("John")
	builder.setFamily("Smith")

	person := builder.Build()
	fmt.Println(person)
}

func (b *PersonBuilder) setName(name string) {
	b.Name = name
}

func (b *PersonBuilder) setFamily(family string) {
	b.Family = family
}

func (b PersonBuilder) Build() Person {
	return b.Person
}
