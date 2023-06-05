package main

import (
	"fmt"
	"sort"
)

func main() {

	var fruitList = []string{"Apple", "Strawberry", "Peach"}
	fmt.Printf("Type of fruitList is %T\n", fruitList) //[]string

	fruitList = append(fruitList, "Mango", "Banana")

	fruitList = append(fruitList[1:3]) //it will add only elements from position 1 to 3 in slice and remove others

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 634
	highScores[2] = 982
	highScores[3] = 932

	highScores = append(highScores, 722, 234, 657) //initially we have declaired the size of slice as 4
	//but we can add more values to it with the help of append method

	sort.Ints(highScores)

	//how to remove a value from slices based on index

	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...) //["reactjs", "javascript", "python", "ruby"]

}
