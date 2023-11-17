package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {

	// net pkg is used for network interface
	resp, err := http.Get("http://www.google.com")
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	lw := logWriter{}

	// i.o copy func is pulling out the data out of something that implements Reader interface
	// then passing all the data insdie of it off to Writer
	//first arg of something that implements Reader interface
	// second arg of something that implements Writer interface

	//io.Copy(os.Stdout, resp.Body)

	io.Copy(lw, resp.Body) // lw is a value implementing Writer interface
}
func (logWriter) Write(bs []byte) (int, error) {

	// int = no of byters just processed
	fmt.Println(string(bs))
	fmt.Println("just processed this many bytes :", len(bs))
	return len(bs), nil

}
