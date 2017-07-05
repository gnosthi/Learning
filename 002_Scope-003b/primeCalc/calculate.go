package primeCalc

func Calculate(num int) {
	var primeBool bool
	for i := num; i > 0; i-- {
		primeBool = isThreeOrLess(num)
		if primeBool == true {
			confirmPrime(num)
			break
		}

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