package main

import (

   "fmt"
   "os"

)


func sayHello(s string) string {
     return "Hello " + s
}


func main() {
//    fmt.Println(len(os.Args))
  
    if(len(os.Args) >= 2) {
         fmt.Println(sayHello(os.Args[1]))
    }else{
    	fmt.Println(sayHello("Mangoose"))
    }
}
