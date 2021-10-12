package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	// contact   contactInfo
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		// contact: contactInfo{
		// 	email:   "jim@email.com",
		// 	zipCode: 987654,
		// },
		contactInfo: contactInfo{
			email:   "jim@email.com",
			zipCode: 987654,
		},
	}
	fmt.Printf("%+v", jim)
}
