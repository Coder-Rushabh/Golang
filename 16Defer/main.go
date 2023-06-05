package main

import "fmt"

func main() {

	fmt.Println("Hello")
	defer fmt.Println("World!")
	fmt.Println("Hello World!")
	main2()
}

/* OP-------------- */
/* Hello
Hello World!
World! */

//Defer will be executed at last.
// It works in Last in first out maner

func main2() {
	defer fmt.Println("World!")
	defer fmt.Println("1!")
	defer fmt.Println("2!")
	fmt.Println("Hello")

}

/* OP-------------- */
/* Hello
2!
1!
World! */
