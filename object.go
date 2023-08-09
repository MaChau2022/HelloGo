package main

import "fmt"

func useArray() {
	list_1 := [5]int{1, 2, 3, 4, 5}
	list_2 := [...]int{1, 2, 3, 4, 5}
	list_3 := [5]int{2: 3, 4: 5}

	slice_1 := list_1[2:]
	slice_2 := list_2[:3]

	fmt.Println(list_1)
	fmt.Println(list_2)
	fmt.Println(list_3)
	fmt.Println(slice_1)
	fmt.Println(slice_2)
}

type Iduck struct {
	name   string
	age    int
	color  string
	weight float32
}

func useDict() {
	duck_1 := Iduck{name: "Iduck", age: 15, color: "yellow"}
	duck_2 := Iduck{"Iduck", 15, "yellow", 5.6}

	fmt.Println(duck_1)
	fmt.Println(duck_2)
}
