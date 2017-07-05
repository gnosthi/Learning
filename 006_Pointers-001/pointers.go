package main

import "fmt"

func main() {
	a := 43

	fmt.Println("Print value of a: ", a)
	fmt.Println("Print memory address of a: ", &a)

	//reference mem address of a
	var b *int = &a

	fmt.Println("Print b as pointer: ", b)
	fmt.Println("Memory address of b should point to memory address of a")
	//Dereference mem address of a
	fmt.Println("Print b's value using *b, this called dereferencing", *b)

	//Assign new value to memory address.
	*b = 666
	//Print "a" which will have been updated to new value
	fmt.Println(a)

	/*
	The above code makes b a pointer to the memory address where a resides.
	b is of type "int pointer"
	*int -- the * is part of the type -- b is of type *int
	 */
}
