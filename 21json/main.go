package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`              //it will not show this in json
	Tags     []string `json:"tags,omitempty"` //if field in nil, it will not show it
}

func main() {
	//EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	Mycourses := []course{
		{"ReactJS Bootcamp", 299, "go.dev", "abc123", []string{"web-dev", "JS"}},
		{"MERN Bootcamp", 499, "go.dev", "bu33d", []string{"web-dev", "MERN"}},
		{"Angular Bootcamp", 699, "go.dev", "h2bws", nil},
	}

	//package this data as JSON data

	finalJson, err := json.MarshalIndent(Mycourses, "", "\t")

	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDatafromWeb := []byte(`
	
	{
		"coursename": "ReactJS Bootcamp",
		"Price": 299,
		"website": "go.dev",
		"tags": ["web-dev","JS"]
    }

	`)

	var anotherCourse course

	checkValid := json.Valid(jsonDatafromWeb)

	if checkValid {
		fmt.Println("JSON is valid!")
		json.Unmarshal(jsonDatafromWeb, &anotherCourse)
		fmt.Printf("%#v\n", anotherCourse)
	} else {
		fmt.Println("JSON is not valid!")
	}

	//some cases where you just want to add data to key value

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDatafromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and Type is: %T\n", k, v, v)
	}

}
