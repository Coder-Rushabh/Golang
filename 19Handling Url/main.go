package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=gfne83nd"

func main() {

	fmt.Println(myurl) //https://lco.dev:3000/learn?coursename=reactjs&paymentid=gfne83nd

	//parsing

	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)   //https
	fmt.Println(result.Host)     //lco.dev:3000
	fmt.Println(result.Path)     ///learn
	fmt.Println(result.Port())   //3000
	fmt.Println(result.RawQuery) //coursename=reactjs&paymentid=gfne83nd

	//extracting query parameters from URL

	qparams := result.Query()
	fmt.Printf("The type of query params are: %T\n", qparams) //qparams - query parameters
	//OP - url.Values

	fmt.Println(qparams["coursename"]) //[reactjs]

	//iterating through query

	for _, val := range qparams {
		fmt.Println("Param is: ", val)
	}

	/* OP - Param is:  [reactjs]
	Param is:  [gfne83nd] */

	//making query through parts of URL

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=rushabh",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL) //https://lco.dev/tutcss

}
