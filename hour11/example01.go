package main

import (
  "fmt"
  "time"
)

func slowFunc() {
   time.Sleep(time.Second * 2)
   fmt.Println("slowFunc() finished")
}
func main() {
    slowFunc()
    fmt.Println("I am not shown until slowFunc() completes")
}
