package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	// exercise0()
	// exercise1()
	// exercise2()
	exercise3()
}

// documentation
func exercise0() {
	res, err := http.Get("http://www.google.com/")
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)
}

func exercise1() {
	response, err := http.Get("https://www.google.com/")

	if err != nil {
		fmt.Println("ERROR :- ", err)
		os.Exit(1)
	}

	fmt.Printf("%+v", response)

	fmt.Println("\n\nSPEW OUTPUT")
	scs := spew.ConfigState{MaxDepth: 2}
	scs.Dump(response)
}

// 1st way, using Reader interface
func exercise2() {
	response, err := http.Get("https://www.google.com/")

	if err != nil {
		fmt.Println("ERROR :- ", err)
		os.Exit(1)
	}
	bs := make([]byte, 99999)
	response.Body.Read(bs)
	fmt.Println(string(bs))
}

// 2nd way, using io.Copy(Writer, Reader)
func exercise3() {
	response, err := http.Get("https://www.google.com/")

	if err != nil {
		fmt.Println("ERROR :- ", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, response.Body)
}

// creating type with Writer interface
type logWriter struct{}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:- ", len(bs))
	return len(bs), nil
}
