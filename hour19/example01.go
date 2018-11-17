package main

import (
   "fmt"
   "io/ioutil"
   "log"
   "net/http"
   "os"
)

func main() {

    response, err := http.Get(os.Args[1])

    if err != nil {
      log.Fatal(err)
    }
    defer response.Body.Close()

    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
       log.Fatal(err)
    }
    fmt.Printf("%s", body)
}
