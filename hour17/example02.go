package main

import (
  "flag"
  "fmt"
)

func main() {
   s := flag.String("s","Hello World", "String help text")
   m := flag.String("m", "Message to the world", "Type your message to the world!")
   flag.Parse()
   fmt.Println("value of s: ", *s)
   fmt.Println("value of m: ", *m)
}
