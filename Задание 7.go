package main
import (
 "fmt"
 "time"
)

func main() {
 setChan := make(chan string) 
 getChan := make(chan string) 

 go func() {
  state := "начальное состояние" 
  
  for {
   select {
   case newState := <-setChan: 
    state = newState     
    fmt.Printf("Установлено: %s\n", state)
    
   case getChan <- state: 
   }
  }
 }()

 set := func(newState string) {
  setChan <- newState
 }

 get := func() string {
  return <-getChan
 }
 go func() {

  set("состояние 1")
  time.Sleep(100 * time.Millisecond)
  
  set("состояние 2") 
  time.Sleep(100 * time.Millisecond)
  
  set("состояние 3")
 }()

 go func() {
  for i := 0; i < 3; i++ {
   state := get()
   fmt.Printf(" Прочитано: %s\n", state)
   time.Sleep(150 * time.Millisecond)
  }
 }()
 time.Sleep(1 * time.Second)
 fmt.Println("Готово")
}
