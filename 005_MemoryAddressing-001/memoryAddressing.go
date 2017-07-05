package main

import "fmt"

//Currently Here!

const (
	metersToYards float64 = 1.09361
)

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Printf("Number: %d\t\t", i)
		fmt.Printf("Binary: %b\t\t", i)
		fmt.Printf("Hexadecimal: %x\t\t", i)
		//grab the memory address for "i"
		//memory address will be the same for each iteration.
		fmt.Printf("Memory Address: %v\n", &i)
	}

	//init var which will assign a mem address to var meters.
	var meters float64
	fmt.Print("Enter meters swam: ")
	//read value into memory address of "meters"
	fmt.Scan(&meters)
	yards := meters * metersToYards
	fmt.Println(meters, " meters is ", yards, " yards.")
	fmt.Printf("Number: %f\t\t", yards)
	fmt.Printf("Binary: %b\t\t", yards)
	fmt.Printf("Hexadecimal: %x\t\t", int(yards))
	//Grab mem address of yards.
	fmt.Printf("Memory Address: %v\n", &yards)
}
