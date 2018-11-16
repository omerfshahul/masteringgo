package main

import (
 "fmt"
)

func getPrice() (int, string) {

   i := 2
   s := "goldfish"

   return i, s
}

func main() {
   quantity, prize  := getPrice()
   
   fmt.Printf("You won %v %v \n", quantity, prize)

}
