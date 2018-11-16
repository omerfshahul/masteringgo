package main

import (
 "fmt"
)

func sayHi() (x, y string) {

	x = "Hi "
        y = "There"
	return
}

func main() {

	fmt.Println(sayHi())
}
