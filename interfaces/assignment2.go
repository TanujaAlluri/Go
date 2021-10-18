package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fileName := os.Args[1]
	filePointer, err := os.Open(fileName)
	if err != nil {
		fmt.Errorf("Error: ", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, filePointer)
}
