package main

import "fmt"

// create a new type of 'deck' which is a new slice of string
type deck []string

// receiver: setsup methods on a type,
// so that the method is accessible to all variables of that type
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
