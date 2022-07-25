package main

import(
	"fmt"
)

var name string

func main() {
	greet()
	playing()
}

func greet() {
	fmt.Print("Please enter your name: ")
	fmt.Scan(&name)
}

func playing() {
	fmt.Print("Would you like to play ", name, "?\n")
	var playing string
	fmt.Scan(&playing)
	if playing != "yes"{
		fmt.Print("Maybe next time then.\n")
		return
	} else {
		fmt.Print("Thank you for playing.\n")
	}
	}