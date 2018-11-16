// Package example05 shows how to use the godoc tool.	
package example05

import (
  "errors"
)


//Animal specifies an animal

type Animal struct {
  Name string // Name hold the name of the animal

  //Age holds the age of the animal
  Age int
}
//ErrNotAnAnimal is returned if the name field of the Animal is human

var ErrNotAnAnimal = errors.New("Name is not an animal")

//Hello sends a greeting to the animal

func(a Animal) Hello() (string, error) {
    if a.Name == "Human" {
         return "", ErrNotAnAnimal
    }
    s := "Hello " + a.Name
    return s, nil
}
