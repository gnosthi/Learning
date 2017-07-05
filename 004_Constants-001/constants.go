package main

import "fmt"

const (
	A = iota
	B = iota
	C = iota
)

func printIota() {
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
}

func main() {
	const p string = "death and taxes"

	const q = 42

	fmt.Println(p)
	fmt.Println(q)

	printIota()
}