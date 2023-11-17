package main

import "fmt"

// struct is like a collection of diff properties that are somehow related
// together or have a common purpose

type contactInfo struct { //  defining a struct
	email   string
	zipCode int
}
type person struct {
	firstName   string // field
	lastName    string
	contactInfo // OR contact : contactInfo, embedded struct
}

func main() {
	// declarationn, intializing and assigning values usig ":="

	man := person{
		firstName: "Divyank",
		lastName:  "Singh",
		contactInfo: contactInfo{
			email:   "anu@gmail.com",
			zipCode: 250110,
		},
	}
	// 2nd way = var man person
	// man.firstName = "divyank", man.lastName= "singh"
	// fmt.Println(man)
	fmt.Printf("%+v", man) // %+v will include the struct's field names.
	fmt.Println()

	man.updateName("anu")
	man.print()

}

// struct with pointers
func (P *person) updateName(newName string) {
	//man variable (of type person) is turn into pointer person by Go
	// a pointer p of type person
	P.firstName = newName // or (*P).firstName
	// actual pointer that takes this pointer & points to actual value
}

// struct with receiver
func (p person) print() {
	fmt.Printf("%+v", p)
}
