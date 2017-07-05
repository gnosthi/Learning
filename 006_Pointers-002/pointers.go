package main

import "fmt"

func zero(x int) {
	x = 0
}

func zeroTest(z int) {
	fmt.Printf("%p\n", &z) //mem address of x in func zeroTest
	fmt.Println(&z) //address in func zeroTest
	z = 0
}

func main() {
	x := 5
	zero(x)
	fmt.Println(x)

	fmt.Printf("%p\n", &x) // address in main
	fmt.Println(&x) //address in main
	zeroTest(x)
	fmt.Println(x) //x is still 5.
}
