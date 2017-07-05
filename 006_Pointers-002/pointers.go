package main

import "fmt"

func zero(x int) {
	x = 0
}

func zeroTest(x int) {
	fmt.Printf("%p\n", &x) //mem address of x in func zeroTest
	fmt.Println(&x) //address in func zeroTest
	x = 0
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
