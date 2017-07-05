package primeCalc

import "fmt"

func confirmPrime(a int) {
	fmt.Printf("%v is a prime number\n", a)
}

func confirmNotPrime(a int, b int, c int) {
	fmt.Printf("%v is not a prime number. It is the sum of %v * %v", a, b, c)
}