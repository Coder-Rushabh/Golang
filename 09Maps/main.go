package main

import "fmt"

func main() {
	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["Go"] = "GoLang"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages ", languages)
	fmt.Println("JS is short for ", languages["JS"])

	//loops

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

}
