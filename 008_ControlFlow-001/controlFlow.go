package main

import (
	"fmt"
	"time"
	"os/exec"
	"os"
)

func main() {

	x := 50

	for i := 0; i <= 100; i++ {
		fmt.Println("i is: ", i)
		x--
		fmt.Println("x is: ", x)
		if x == 33 || x == 13 {
			time.Sleep(time.Second * 1)
		} else if x == 22 || x == 11 {
			time.Sleep(time.Second * 1)
		}

		if i == 13 {
			cmd := "/bin/bash"
			args := []string{"golangexec.sh"}
			if err := exec.Command(cmd, args...).Run(); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		}
	}
}