package main

import (

  "fmt"
)

func main() {

   s := "Hello World"
   t := "Goodbye Cruel World"
   echo(s)
   echo(t)
}
func echo(s string) {
   fmt.Println(s)
   return
}
