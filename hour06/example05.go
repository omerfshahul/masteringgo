package main

import (
 "fmt"
)

func main() {
 
   var cheeses = make([]string, 3)
   cheeses[0] = "Mariolles"
   cheeses[1] = "Eposses de bourgogne"
   
   var smellyCheeses = make([]string,2)
   copy(smellyCheeses, cheeses)
   fmt.Println(smellyCheeses)
}
