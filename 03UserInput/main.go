package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Print(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name: ")

	// comma ok
	// err ok

	input, _ := reader.ReadString('\n')
	fmt.Println("Hi, ", input)

	//when there are chances of error in code, we can give another variable in place of
	//underscore(_). It will execute that variable if error occurs.

}
