package main

import (
	"os"
	"io"
	_ "fmt"
)

func main() {

	myString := ""
	arguments := os.Args
    
	if len(arguments) == 1 {
		myString = "Please give me one argument"
	} else {
		myString = arguments[1]
	}

	io.WriteString(os.Stdout, myString)
	io.WriteString(os.Stdout, "\n")
}