package main

import "fmt"

type Person struct {
	name *string
	age  int
}

func (p *Person) setName(name string) {
	*p.name = name
	fmt.Println(p.name)
}

func callPerson() {
	p := new(Person)
	p.setName("Tom")
}
