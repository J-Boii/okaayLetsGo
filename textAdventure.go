package main

import(
	"fmt"
	"os"
	"bufio"
)

func main() {
	isPlaying()
	getName()
	

}

func getName() {
	fmt.Print("Please enter your name: ")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Print("Lovely to meet you, ", name)
}

func isPlaying() {
	fmt.Print("Hi there. Would you like to play my game? (yes/no): ")
	reader := bufio.NewReader(os.Stdin)
	playing, _ := reader.ReadString('\n')
	if playing == "yes" {
		fmt.Print("You did not enter yes")
		break
	} 
	fmt.Print("playing = ", playing, "\n")
	fmt.Print("Maybe next time then...\n")
	quit()
}

func quit() {

}

