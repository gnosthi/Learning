package main

import "fmt"

func testBones(x int, y int) int {
	return x * y
}

func main() {
	x := "Hello, World"
	a := 3
	b := 4
	y := testBones(a, b)
	fmt.Println(x)

	if y == 4 {
		fmt.Println("y = 10")
	} else {
		fmt.Println("y =", y)
	}
}