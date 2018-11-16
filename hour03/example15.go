package main

import (
  "fmt"
)

const greeting string = "Hello"

func main() {

  fmt.Println(greeting)
  
  greeting = "World"
 
  fmt.Println(greeting)
}
