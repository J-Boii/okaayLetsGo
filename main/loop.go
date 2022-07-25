package main

import(
	"fmt"
)

var count int

func up() {
	if count == 1000000 {
		return
	} else {
		fmt.Print(count, "\n")
		count += 1
		up()
	}
}

func main() {
	up()
}