package main

import "fmt"

func printWords(a string) {
	fmt.Println(a)
}

func printNumbers(a int) {
	fmt.Println(a)
}

/*func printType(a ...interface()) {
	fmt.Printf("%T," a)
}*/

//Shpuld stay away from Globals where possible.
var Complete string = "ABRAHADABRA"

func main() {

	//shorthand
	a := "Do what thou wilt shall be the whole of the Law."
	b := "Love is the law, love under will."

	c := 93

	printWords(a)
	printWords(b)

	printNumbers(c)

	printWords(Complete)

	fmt.Printf("%v, is type: %T\n",a, a)
	fmt.Printf("%v, is type: %T\n",b, b)
	fmt.Printf("%v, is type: %T\n",c, c)
}
