package main

import "fmt"

func main() {

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for _, n := range numbers {
		if n%2 == 0 {
			fmt.Println("number is even")
		} else {
			fmt.Println("num is odd")
		}
	}

	for _, n := range numbers {
		if n%2 == 0 {
			fmt.Println(n, "is even")
		} else {
			fmt.Println(n, "is odd")
		}
	}
}
