package main

import (
	"fmt"
	"time"
)

func exec() {
	fmt.Println("Hello, World!")
}

func fork() {
	go exec()
	exec()
	time.Sleep(time.Second)
}
