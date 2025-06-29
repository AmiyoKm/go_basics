package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	r, err := os.Open("hello.txt")
	if err != nil {
		panic("something went wrong")
	}
	defer r.Close()

	buffer := make([]byte, 100)
	n, err := r.Read(buffer)

	fmt.Println(n)
	fmt.Println(string(buffer))

	w, err := os.Create("output.txt")
	if err != nil {
		panic(err)
	}
	defer w.Close()

	_, err = io.Copy(w, os.Stdin)
	if err != nil {
		panic(err)
	}

	fmt.Println("Data copied from stdin to file")
}
