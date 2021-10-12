# Go Udemy course - "Go: The Complete Developer's Guide (Golang)"

### How to run the file:

`go run hello-world.go `

`go run <filename>.go`

### GO CLI commands:

- go run : takes input of 1 or more files, compiles and executes them instantly

- go build: doesnt execute them instantly, creates an executable. compiles a bunch of src code files.

- go fmt: formats the code

- go install: compiles and installs a package

- go get: downloads the raw src code of someone's package

- go test: runs any tests

### Types of packages:

- Executable: 
  - generates runnable executable
  - package main => executable or runnable file
  - must always have a function called 'main'
- Reusable: 
  - code used as helpers
  - package blahblah => reusable package

### Data types:

Go, java, C++ => statically typed languages

Javascript, python, ruby => dynamically typed languages

basic go data types:
- bool
- string
- int
- float64

zero value or default value:
- string: ""
- int: 0
- float64: 0
- bool: false

Note: default value is not 'nil'

### basic Data structures:

- Array: fixed length
- Slice: 
  - array with variable length
  - slices are zero-indexed
  - accessing ex: 
      `fruits := []string {"grape", "apple", "orange", "banana"}
       
       fruits[0] // reutrns "grape"
       
       fruits[0:3] // returns "grape", "apple", "orange"
       
       fruits[:3] // returns "grape", "apple", "orange"
       
       fruits[2:]// returns "orange", "banana"`
  - 

### Receiver:

A function which takes the argument of a custom data type.
It sets up methods on a type so that the method is
accessible to all variables that we declare of that type.

By creating a new type with a function that has a receiver, we
are adding a 'method' to any value of that type.

Syntax: `func (<name> <type>) <name-of-fun>(zero or more args<arg-name arg-type>) <return-type> {}`

Ex: `func (d deck) toString() string {}`

Usage: `cards := deck{} cards.toString()`

### Type conversion with Go:

From string to slice byte - Ex: `[]byte("Hi There!")`

### Testing:

- Test file name ends with _test.go
- Command: `go test`
- No testing framework

### Structs:

- data structure in Go
- collection of properties that are related together
- similar to an object in javascript (very high level)






