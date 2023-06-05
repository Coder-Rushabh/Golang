//when we pass variable to any program, copy of it is being passed.
//but in case of pointer actual value will be passed, it is a direct reference to memory address

package main

import "fmt"

func main() {

	var pointer *int

	fmt.Println("Vaalue of Pointer is:", pointer) //OP <nil>

	myNumber := 23

	var pointer2 = &myNumber //it is referncing to myNumber thats why we used '&'

	fmt.Println("Value of Pointer 2 is", pointer2)  //it will return address
	fmt.Println("Value of Pointer 2 is", *pointer2) //it will return actual value: 23

	*pointer2 = *pointer2 + 2
	fmt.Println("New value is:", myNumber) //25

}
