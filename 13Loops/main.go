package main

import "fmt"

func main() {

	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}

	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	for i := range days {
		fmt.Println(days[i])
	}

	for index, day := range days {
		fmt.Printf("index is %v and value is %v", index, day)
	}

	value := 1

	for value < 10 {

		if value == 2 {
			goto vari
		}

		if value == 5 {
			break
		}

		fmt.Println("Value is: ", value)
		value++
	}

vari:
	fmt.Println("Program ends!")

}
