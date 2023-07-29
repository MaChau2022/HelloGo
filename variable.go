package main

import "fmt"

func useVariable() {
	var name string = "John Doe"
	var age = 15
	height := 5.6
	const gender = true

	fmt.Println("name:", name)
	fmt.Println("age:", age)
	fmt.Println("height:", height)
	fmt.Println("gender:", gender)
}

func exchange() {
	a := 10
	b := 20

	fmt.Println(a, b)
	a, b = b, a
	fmt.Println(a, b)
}
