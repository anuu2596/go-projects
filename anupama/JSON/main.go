package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type student struct { // JSON to struct, map etc
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Phone string `json:"phone"`
}

type animal struct { // struct, map etc to JSON string
	Category string `json:"category"`
	Color    string `json:"color"`
	Voice    string `json:"voice"`
	Legs     int    `json:"legs"`
}

func main() {

	s1 := student{}
	abc := `{ "name":"Anu", "age" :26, "phone": "8587998002" }`
	err := json.Unmarshal([]byte(abc), &s1)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Printf("%+v", s1)
	fmt.Println()

	//////////////////////////

	a1 := animal{
		Category: "dog",
		Color:    "brown",
		Voice:    "bark",
		Legs:     4,
	}
	str, err := json.Marshal(a1) // str is type of []byte
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println("=========")
	fmt.Println(string(str)) // converting str []byte to string
	fmt.Println("=========")
}
