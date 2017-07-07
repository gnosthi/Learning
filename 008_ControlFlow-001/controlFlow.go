package main

import (
	"fmt"
	"time"
	"os/exec"
	"os"
	"math/rand"
	"runtime"
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


		rand.NewSource(time.Time.UnixNano(time.Now()))
		randSeed := rand.Intn(1000000)
		rand.Seed(time.Now().UnixNano() * int64(randSeed))
		randNum := rand.Intn(100)
		fmt.Println("rand is: ", randNum)
		if runtime.GOOS == "linux" {
			if i == randNum {
				cmd := "/bin/bash"
				args := []string{"golangexec.sh"}
				if err := exec.Command(cmd, args...).Run(); err != nil {
					fmt.Fprintln(os.Stderr, err)
					os.Exit(1)
				}
			}
		} else if runtime.GOOS == "windows" {
			fmt.Println("This Windows... can't run shell script")
		} else {
			fmt.Println("Unable to determine OS type")
		}
	}


}