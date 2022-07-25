package main

import "fmt"

func main() {
	fmt.Print("Please enter your name: ")
	var name string
	fmt.Scan(&name)
	fmt.Print("Lovely to meet you " + name + "\n")
}