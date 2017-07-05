package main

import "fmt"

func main() {

	x := 50

	for i := 0; i <= 100; i++ {
		fmt.Println(i)
		x--
		fmt.Println(x)
	}
}
