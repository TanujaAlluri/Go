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
- when run `jim := person("jim","party")`; Jim is actually a reference to the struct in memory

### Pointers:

- Go is a <b>PASS BY VALUE</b> language

If we run `jim := person("jim","party")`; 
jim is a pointer to the value at address 101

| address | val                   |
|---------|-----------------------|
| 101     | person("jim","party") | <---- jim 

When we run `jim.updateName("jimmy") ... func (p person) updateName(newFirstName string)`
GO makes a copy of jim and populates it in variable p. and the copy is available in updateName function

| address | val                   |
|---------|-----------------------|
| 101     | person("jim","party") | <---- jim 
|...      |...                    |
| 105     | person("jim","party") | <---- p

And after running `p.firstName="jimmy"` inside the receiver function

| address | val                     |
|---------|-------------------------|
| 101     | person("jim","party")   | <---- jim 
|...      |...                      |
| 105     | person("jimmy","party") | <---- p

#### Pointer operators

Pointer is a variable that stores an address of a variable
- '&variable' => give me the memory address of this variable
- '*pointer' => give me the value this memory address is pointing at
- '*type' => Ex: `pp *person` => pp should be a pointer for a person type


| address | val                     |
|---------|-------------------------|
| 101     | person("jim","party")   | 
    ^             ^
    |             |
    |             |
    |             |
  address       value
  Ex: jimpointer  Ex: jim

address to value: *address
value to address: &value

#### GO Shortcut

If we declare a receiver on pointer to a type,
Go allows us to call the function on:
- pointer to the type
- the type itself

Ex: Receiver: `func (pp *person) updateName() {}`
We can call using:
- personPointer := &person{} personPointer.updateName()
- person{}.updateName()



### Reference types vs Value types

Arrays: primitive data structure, cant be resized, rarely used directly

Slice: Can grow or shrink, Used 99% of the time for lists of elements, Fancy Array

When we declare a slice:
`mySlice := []string{"Hi", "There", "How", "Are", "You"}`

2 data structures are created:

- slice: properties - pointer to head, capacity, length
- array: | "Hi" | "There" | .....| "You" |

|   Address   |      Value                                  |
| ----------- | ---------------------------------------     |
|   000       |                                             |
|   001       |      length, capacity, pointer to head --   |   <----- mySlice
|   002       |       []string{"Hi", "There", ...}  <----|  |

pointer to head points to the actual array at 002 address

Now when we pass a slice to a function: `func updateSlice(s []string) {}`
Go still makes a copy of the slice datastructure.

But then, the new copy of slice still points to the same array

Hence, when we try to modify the underlying array, since both are pointing to the same array, the implementation might not feel like 'Pass by value'

`mySlice := []string{"Hi", "There", "How", "Are", "You"}
updateFirstElement(mySlice) // returns => {"Bye", "There", "How", "Are", "You"}

func updateSlice(s []string) {s[0]="Bye"}`


| Value types           |     Reference types     |
| --------------------  |  ---------------------- |
| int                   |     Slices              |
| float                 |     Maps                |
| string                |     Channels            |
| bool                  |     Pointers            |
| structs               |     Functions           |
|                       |                         |
| 'Use pointers         |     'Dont worry         |
| to change             |     about pointers'     |
| these in a func'      |                         |


### Maps

- Similar to object in javascript
- collection of key value pairs
- keys and values are statically typed
- all keys should be same type
- all values should be same type


#### Differences between maps and structs

Map:
- All keys are of same type
- All values are of same type
- Keys are indexed, we can iterate over them
- Reference type :=> we are passing by references via functions
- dont need to know all keys at compile time
- used to represent a collection of similar or related proerties


Struct:
- Values can be of different type, keys are not statically typed
- Keys dont support indexing
- Value type 
- need to know all keys at compile time
- used to represent one "thing" with lot of properties

### Interfaces

declare an interface:

`type bot interface { getGreeting(<list-ofargs>) <list-of-return-types>}`

`type bot interface { 
  getGreeting(int, string) (string, error)
  getBotVersion() float64
  ...
}`

implies, we have a new type bot. And any types that have a getGreeting() function that returns a string are members or implementations of this type/ interface.

And all functions on type bot, are accessible to the implementations/ members of the interface type

- interfaces are not generic types : other languages have generic types, Go doesnt
- interfaces are implicit: no manual code to link together bot and english bot, no explicit link like ' type englishBot implelements Bot '
- interfaces are a contract to help us manage types and help us reuse code

#### Concrete type and interface types

- concrete type
  - we can create a value out of this type
  - ex: map, struct, int, string

- interface type
  - we cannot create a value out of this type

Forming an interface from other interfaces:

specifies that if u want to satisfy ReadCloser interface,
satisfy Reader and Closer interfaces.

`type ReadCloser interface {
	Reader
	Closer
}`

### Channels and Go Routines

Channels and GO routines are structures in Go, for concurrent programming

#### Go Routines

Every running program has one Go Routine by default

One Go Routine executes our code line-by-line

Go routine :=> process that executes the code

Go has a scheduler that schedules the Go routines on a one core CPU

By default go runs on 1 core of the CPU

Concurrency => technique by which the program deals with multiple processes. But they are not exactly run at the same time.

Scheduler : Runs one routine until it is finished or makes a blocking call. 

After that routine is blocked or finished, another one is picked. (By default, since one CPU core)

We can configure Go to run on multiple CPUs

If there are 3 CPU cores and 3 Go routines, Scheduler will simply assign each Go routine to each CPU core.

Parallelism => when processes are run on 2 different CPU cores, such that there is a process running exactly at the same time. Doing multiple processes at the same time is parallelism. Dealing with multiple processing is concurrency.

Our running program consists of:
- Main routine : main routine created when we launch our program
- Child Go routines : Child routines created by the 'go' keyword

Use 'go' keyword only before function calls


time ------>

------------------------------------------------------------->
|program started

|-----main routine-----------------|(exits)

        |------child routine--------------------|

        |---------child routine---------------|

        |-------child routine---------------------|


Main routine doesnt care if there is some code execution left for the child routines. Hence, it might exit even before the child routine code actually finishes. And the program quits, and we cant see the result.

#### Channels

Channels are used for communicating between Go routines.

If we send something into a channel, that will be sent to any other routine that has access to that channel

Channels are similar to values that we create, like an int, a string, a struct

Channels are also typed. Meaning we can pass values of a single data type in a given channel

We create a channel that is meant for sharing strings between say A and B Go routines. We can only send strings through this channel.

Sending data through channels:

channel<-5  : sending the value '5' to a channel

myNumber<-channel : wait for a value to be sent from the channel. When we get one, assign that value to 'myNumber'

fmt.printLn(<-channel) : wait for a value to be sent from the channel. When we get one, log it out immediately.


'<-channel' or receiving message from a channel is a blocking call just as http request is a blocking call. Whenever main routine sees a '<-channel' call, it puts the main routine to sleep, and another go routine is picked up.



----------> time

---------------------------------------------------------------------------------->

      |program started

      |----main go routine-----|(for loop execution complete)     |---main go routine----|(main go routine prints the values received in channel, and exists since no more code to execute)

                        |-----google go routine-------------------|(google routine finished, sends message to channel, main routine wakes up and receives message from channel)

                            |----facebook go routine-----------------------------------------------|

                                |----stackoverflow go routine-------------------------------------|



'Receiving a value from a channel' is a blocking call. If no one is sending any value into the channel, the routine sleeps, and waiting to receive a value from the channel. If this is main routine, this will cause the program to hang.


Using 'range' keyword with slice := take element from the slice, assign it to a variable, that we can access inside the for loop

Using 'range' keyword with channel := wait for the channel to return a value. After the channel has returned a value, assign it to a variable, that we can access inside the for loop
















