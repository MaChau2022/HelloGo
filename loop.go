package main

import "fmt"

func useLoop() {
	list := [4]string{"Alice", "Bob", "Cassandra", "Dove"}
	dict := map[string]string{
		"Alice":     "A+",
		"Bob":       "B",
		"Cassandra": "F",
		"Dove":      "A-",
	}

	for i := 0; i < len(list); i++ {
		fmt.Print(list[i] + ", ")
	}
	fmt.Print("\n")

	for key, _ := range dict {
		fmt.Print(key + ", ")
	}
	fmt.Print("\n")
}
