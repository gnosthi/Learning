package main

import (
	"fmt"
	"log"
)

/*
This has been written purposely obtuse.
The purpose is to make one think about variable assignment
and variable scope.
 */


func getInput() int {
	var i int
	fmt.Printf("Enter a number: ")
	if _, err := fmt.Scan(&i); err != nil {
		log.Print(" Scan for i failed due to ", err)
		return i
	}
	return i
}

func isTrue(a int) {
	fmt.Printf("%v is true\n", a)
}

func isNotTrue(a int, b int, c int) {
	fmt.Printf("%v is not true\n", a)
	fmt.Printf("%v is the result of %v * %v\n", a, c, b)
}

func mVal(a int, b int) (c int, d int) {
	return a % b, a / b
}

func main() {
	a := getInput()
	for i := a; i > 0; i-- {
		if a <= 3 {
			isTrue(a)
			break
		}

		if i == a {
			i--
		}

		b, c := mVal(a, i)

		if i > 1 && b == 0 {
			isNotTrue(a, i, c)
			break
		} else if i <= 2 && b != 0 {
			isTrue(a)
			break
		}
	}

}
