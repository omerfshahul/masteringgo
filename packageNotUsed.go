package main

import (
	"fmt"
	_ "os"
       "github.com/matsouk/go/simpleGitHub"
)

func main() {

	fmt.Println("Hello there!")
        fmt.Println(simpleGitHub.AddTwo(5,6))

