package main

import (
  "fmt"
)

type SuperHero struct {

   Name string
   Age  int
   Address Address
}

type Address struct {
    Number int
    Street string
    City   string
}

func main() {
   e := SuperHero {
     Name: "Batman",
     Age:  32,
     Address: Address{
        Number: 1007,
        Street:  "Mountain Drive",
        City:   "Gotham",
     },
   }
   fmt.Printf("%+v\n", e)
}
