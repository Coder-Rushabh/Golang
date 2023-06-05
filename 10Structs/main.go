package main

import "fmt"

func main() {

	rushabh := User{"Rushabh", "rushabh@gmail.com", true, 22}
	fmt.Println(rushabh)                      //{Rushabh rushabh@gmail.com true 22}
	fmt.Printf("Details are: %+v\n", rushabh) //Details are: {Name:Rushabh Email:rushabh@gmail.com Status:true Age:22}
	fmt.Printf("Name is %v and email is %v.", rushabh.Name, rushabh.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
