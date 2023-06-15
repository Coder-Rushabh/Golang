package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	//PerformGetRequest()
	//PerformPostJsonRequest()
	PerformPostFormRequest()
}

// GET Request

func PerformGetRequest() {
	const myurl = "http://localhost:3000/get"

	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("status code: ", response.StatusCode)
	fmt.Println("Content length is: ", response.ContentLength)

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

	//Or we can write like this

	var responseString strings.Builder
	content2, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content2)

	fmt.Println("ByteCount is: ", byteCount)
	fmt.Println(responseString.String())
}

// POST Request

func PerformPostJsonRequest() {
	const myurl = "http://localhost:3000/post"

	//fake json payload

	requestBody := strings.NewReader(`
	{
		"coursename": "Go Lang",
		"price": 3000,
		"platform": "go.dev"
	}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}

//Send Form data

func PerformPostFormRequest() {
	const myurl = "http://localhost:3000/postform"

	//form data

	data := url.Values{}
	data.Add("firstname", "Rushabh")
	data.Add("lastname", "Dabhade")
	data.Add("email", "rushabh@go.dev")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

}
