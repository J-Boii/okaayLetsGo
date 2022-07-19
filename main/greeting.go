package main

import (
	"bufio"
	"fmt"
	"os"
)


func greet() {
	reader := bufio.NewReader(os.Stdout)
	var name string
	print("Please enter your name: ")
	name, _ =  reader.ReadString('\n')
	fmt.Print("Lovely to meet you ", name)
}

func main() {
	greet()
}