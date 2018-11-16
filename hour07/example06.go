package main

import (
 "fmt"
)


type Alarm struct {
    Time string 
    Sound string
}


func newAlarm(time string) Alarm {
   a := Alarm {
     Time : time,
     Sound : "Klaxon",
   }
   return a
}

func main() {
   fmt.Printf("%+v\n", newAlarm("07:00"))
}
    
