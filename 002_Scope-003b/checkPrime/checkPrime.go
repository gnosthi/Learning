package checkPrime

import (
	"fmt"
	"log"
	"github.com/gnosthi/Learning/002_Scope-003b/primeCalc"
)

func getInput() int {
	fmt.Printf("Enter a number: ")
	var i int
	if _, err := fmt.Scan(&i); err != nil {
		log.Print("Scan for i failed due to ", err)
		return i
	}
	return i
}

func CheckPrime() {
	numVal := getInput()
	primeCalc.Calculate(numVal)
}