package main

import (
 "fmt"
)

func main() {

    s := 'b'

    switch s {
     case 'a':
       fmt.Println("The letter a!")
     case 'b':
       fmt.Println("The letter b!")
     default:
       fmt.Println("I don't know that letter!")
    }
}
