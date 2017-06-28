package main

import (
	"fmt"
	"time"
)

func doLoop(a int) {
	for i := 0; i < a; i++ {
		fmt.Printf("%d \t %b \t %x \n", i, i, i)
		time.Sleep(time.Second/3)
	}
}

func main() {
	interations := 100
	fmt.Printf("%d - %b %#x \n", 42, 42, 42)
	fmt.Printf("%d \t %b \t %#x \n", 42, 42, 42)
	time.Sleep(time.Second*2)
	doLoop(interations)
}