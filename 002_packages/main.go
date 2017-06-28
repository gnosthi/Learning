package main

import (
	"fmt"
	"github.com/gnosthi/Learning/002_packages/stringutil"
)

func main() {
	fmt.Println("Hello", stringutil.MyName)
	fmt.Println("Reverse!", stringutil.Reverse(stringutil.MyName))
}
