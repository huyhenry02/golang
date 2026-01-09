package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct{}

func (c Cat) Speak() string {
	return "Meow!"
}

func main() {
	var a Animal

	a = Dog{}
	fmt.Println(a.Speak())

	a = Cat{}
	fmt.Println(a.Speak())
}
