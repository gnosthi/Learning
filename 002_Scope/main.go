package main

import "fmt"

//Package level declaration
//Available to whole package.
var x int = 42

func isPrime(a int) {
	fmt.Printf("%v is a prime number", a)
}

func isNotPrime(a int, b int, c int) {
	fmt.Printf("%v is not a prime number\n", a)
	fmt.Printf("It is the result of %v * %v \n", c, b)
}

func analysePrime(a int, b int) int {
	return a % b
}

func divNumbers(a int, b int) int {
	return a / b
}

func main() {
	//Keep your scope tight

	//x is available to any func in the package.
	fmt.Println(x)

	//p scope is only available to func main().
	// would not be available to any other part of the application.
	p := 67
	for i := p; i > 0; i-- {
		if p <= 3 {
			isPrime(p)
			break
		}

		if i == p {
			i--
		}

		//however we can use the value of p as an argument to a function thus allowing the function to be
		//aware of the value but not the variable.
		modValue := analysePrime(p, i)
		divValue := divNumbers(p, i)

		if i > 1 && modValue == 0 {
			isNotPrime(p, i, divValue)
			break
		} else if i <= 2 && modValue != 0 {
			isPrime(p)
			break
		}
	}
}

/*func foo() int {
	//p not available to func foo()
	//fmt.Println(p)

	// h only available to func foo
	h := 93
	return h

}*/
