package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	res, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// bs := []byte{}  ==> returns empty byte slice

	//	so that when Read function fills this slice, there is space.
	// Read function cannot resize the slice that we provide, hence this appraoch.
	bs := make([]byte, 99999) // ==> returns byte slice with 99999 empty spaces,
	res.Body.Read(bs)
	fmt.Println("Byte slice:", string(bs))

	io.Copy(os.Stdout, res.Body)
}

/*
Exploring http package:
------------------------


type Response struct {
	status     int
	statusCode string
	body       io.ReadCloser
}

--- forming an interface from other interfaces ---
specifies that if u want to satisfy ReadCloser interface,
satisfy Reader and Closer interfaces
type ReadCloser interface {
	Reader
	Closer
}

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Closer interface {
	Close() (err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

func Copy(r Writer, r Reader) (written )

*/
