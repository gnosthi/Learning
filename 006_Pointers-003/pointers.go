package main

import "fmt"

func zero(x *int) {
	fmt.Println(x)
	*x = 0
	fmt.Printf("%p\n", x)
}

func main() {
	x := 5
	fmt.Println(x)
	fmt.Println(&x)
	zero(&x)
	fmt.Println(x)
	fmt.Println(&x)
}
