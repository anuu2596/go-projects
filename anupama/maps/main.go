package main

import "fmt"

func main() { // map is a collection of key-value pair.

	//[string] key is of type string
	// string - value of type string
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "sdfer22",
	}

	// var colors map[string]string
	// colors := make(map[string]string)
	// using above syntax when we want to add extra key:value pair
	// colors["white"] = "#ff0010"
	// delete(colors, "white")
	fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println(key, value)
	}
}
