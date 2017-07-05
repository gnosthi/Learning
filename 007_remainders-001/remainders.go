package main

import "fmt"

func main() {
	x := 13 % 3
	fmt.Println(x)
	if x == 1 {
		fmt.Println("ODD")
	} else {
		fmt.Println("EVEN")
	}

	for n := 0; n < 100; n++ {
		if n % 2 == 1 {
			fmt.Println( n, " is ODD")
		} else {
			fmt.Println( n, " is EVEN")
		}
	}
}
