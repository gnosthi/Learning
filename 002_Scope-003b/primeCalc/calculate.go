package primeCalc

import "os"

func Calculate(num int) {
	var primeBool bool = isThreeOrLess(num)
	if primeBool == true {
		confirmPrime(num)
		os.Exit(0)
	}
	for i := num; i > 0; i-- {
		if i == num {
			i--
		}
		modNum, divNum := calculateModAndDiv(num, i)
		if i > 1 && modNum == 0 {
			confirmNotPrime(num, divNum, i)
			break
		} else if i <= 2 && modNum != 0 {
			confirmPrime(num)
			break
		}
	}
}