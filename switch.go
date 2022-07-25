package main

import "fmt"

func main() {
	var x int = 0
	switch true {
	case x == 0:
		fmt.Print("x == 0")
	case x != 0:
		fmt.Print("x != 0")
	}
	fmt.Print("\n")
}