package main

import "fmt"

func input() {
	fmt.Println("Please give your name")
	var name string
	fmt.Scanln(&name)
	fmt.Println("Your name is", name)
}
