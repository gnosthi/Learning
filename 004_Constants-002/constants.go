package main

import "fmt"

const (
	/*
	iota increments by 1 each time it is called.
	A = iota = 0
	B = iota = 1
	C = iota = 2 etc.
	Therefore what is occurring below is assigning the 0 place iota to a blank identifier. Essentially throwing it away
	then we declare that KB has a value of 1 shifted 1 * 10 places in binary. 10th place in binary = 1024 which also is
	the same size as a kilobyte.
	The next decleration states MB has a value of 1 shifted 2 * 10 places  in binary. and so one. Since iota increments.
	 */

	_ = iota // 0
	KB = 1 << (iota * 10) //1 << (1 * 10)
	MB = 1 << (iota * 10) //1 << (2 * 10)
	GB = 1 << (iota * 10) //1 << (3 * 10)
)

func main() {
	fmt.Println("Binary\t\tdecimal")
	fmt.Printf("%b\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t", MB)
	fmt.Printf("%d\n", MB)
	fmt.Printf("%b\t", GB)
	fmt.Printf("%d\n", GB)
}