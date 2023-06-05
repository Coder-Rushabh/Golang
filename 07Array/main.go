package main

import "fmt"

func main() {

	var fruitList [3]string

	fruitList[0] = "Apple"
	fruitList[1] = "Strawberry"
	fruitList[2] = "Mango"

	fmt.Println("Fruit List is", fruitList)
	fmt.Println("Fruit List is", len(fruitList)) //4

	var vegList = [3]string{"Potato", "Beans", "Mushroom"}

}
