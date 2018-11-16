package main

import (

  "fmt"
)

type Drink struct {
    Name string
    ice  bool
}

func main() {

    a := Drink{
      Name: "Lemonade",
      ice: true,
    }

    b := &a

    b.ice = false
    fmt.Printf("%+v\n", *b)
    fmt.Printf("%+v\n", a)
    fmt.Printf("%p\n", b)
    fmt.Printf("%p\n", &a)
}
