package main

import "fmt"

func calculateThing(a int, b int) int {
	return a * b / 2
}

func main() {

	a := 666
	b := "Do what thou wilt shall be the whole of the Law"
	c := 93
	d := "Love is the law, love under will."
	e := calculateThing(a,c)

	fmt.Printf("Decimal: %d, Binary: %b, Hex: %x\n", a, a, a)
	fmt.Printf("String: %s, \n Hex: %x\n",b, b)
	fmt.Printf("Decimal: %d, Binary: %b, Hex: %x\n", c, c, c)
	fmt.Printf("String: %s, \n Hex: %x\n", d, d)
	fmt.Printf("Decimal: %d, Binary: %b, Hex: %x\n", e, e, e)
}
