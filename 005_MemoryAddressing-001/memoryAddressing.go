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
		fmt.Printf("Memory Address: %v\n", &i)
	}

	var meters float64
	fmt.Print("Enter meters swam: ")
	fmt.Scan(&meters)
	yards := meters * metersToYards
	fmt.Println(meters, " meters is ", yards, " yards.")
	fmt.Printf("Number: %f\t\t", yards)
	fmt.Printf("Binary: %b											QQQQQQW3E45\t\t", yards)
	fmt.Printf("Hexadecimal: %x\t\t", int(yards))
	fmt.Printf("Memory Address: %v\n", &yards)
}
