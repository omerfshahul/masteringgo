package main

import (
  "fmt"
)

func main() {
 i := 5

 switch i {

   case 2:
	fmt.Println("Two")
   case 3:
        fmt.Println("Three")
   case 4:
        fmt.Println("Four")
   default:
        fmt.Println("Greater than Four")
  }
}
