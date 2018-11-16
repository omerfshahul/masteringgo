package main


import (
    "fmt"
    "strconv"
    "reflect"
)

func main() {

  var b bool = true
  var s string = strconv.FormatBool(true)

  fmt.Println(reflect.TypeOf(b))
  fmt.Println(reflect.TypeOf(s))

  fmt.Println(s)
}
