package main

import (
	"github.com/gnosthi/Learning/002_Scope-03/viz"
	"fmt"
)

func printName(a string) {
	fmt.Printf("Your name is %v\n", a)
}

func printAge(a int) {
	fmt.Printf("You are %v years old", a)
}

func main() {
	myName := viz.Name
	myAge := viz.Age

	printName(myName)
	printAge(myAge)
}
