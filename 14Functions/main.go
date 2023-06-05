package main

import "fmt"

func main() {
	fmt.Println("Functions in Go")
	greeter()

	result := adder(3, 7)
	fmt.Println("Result is: ", result)

	ProResult := proAdder(4, 7, 2, 9)
	fmt.Println("Pr Result is: ", ProResult)

}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) int {
	total := 0

	for _, val := range values {
		total += val
	}
	return total
}

func greeter() {
	fmt.Println("Hello World!")
}
