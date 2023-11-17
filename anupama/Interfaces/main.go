package main

import "fmt"

// import "fmt"

// // interfaces are used to define func set
// // any type insdie of our program that implments GetGreeting func
// // it will be of bot type
// // we cannot create direct value of interface

// type bot interface {
// 	GetGreeting() string
// }

// type englishBot struct{}
// type spanishBot struct{}

// func main() {
// 	eb := englishBot{}
// 	sb := spanishBot{}

// 	// fmt.Println(eb.GetGreeting())

// 	printGreeting(eb)
// 	printGreeting(sb)
// }

// func printGreeting(b bot) {
// 	fmt.Println(b.GetGreeting())
// }

// func (englishBot) GetGreeting() string {
// 	return "hi there"
// }

// func (spanishBot) GetGreeting() string {
// 	return "hola"
// }

// Square = a*2
// triangle = 1/2*b*h
// Rectangle = l*b
// Circle = 3.14*r*r

type area interface {
	calculateArea() float64
}

type Square struct {
	side    int
	getArea func() int
}

type Triangle struct {
	base   int
	height int
}

type Rectangle struct {
	length  int
	breadth int
}

type Circle struct {
	radius int
}

func main() {
	s := Square{
		side: 2,
	}
	r := Rectangle{
		length:  3,
		breadth: 4,
	}
	c := Circle{
		radius: 6,
	}
	t := Triangle{
		base:   7,
		height: 8,
	}

	// fmt.Println(s.calculateArea())
	// fmt.Println(t.calculateArea())
	// fmt.Println(r.calculateArea())
	// fmt.Println(c.calculateArea())

	printArea(s)
	printArea(r)
	printArea(c)
	printArea(t)

}
func (s Square) calculateArea() float64 {
	return float64(s.side * s.side)
}
func (r Rectangle) calculateArea() float64 {
	return float64(r.length * r.breadth)
}
func (c Circle) calculateArea() float64 {
	return float64(c.radius*c.radius) * 3.14
}
func (t Triangle) calculateArea() float64 {
	return float64(t.base * t.height)
}

func printArea(a area) {
	fmt.Println(a.calculateArea())
}
