package main

import "fmt"

func pointer() {
	a := 10;
	p := &a;
	q := &p;
 
	fmt.Println(a, &a)
	fmt.Println(*p, p)
	fmt.Println(*q, q)
}