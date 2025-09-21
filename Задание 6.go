package main

import (
 "fmt"
 "sync"
 "time"
)

func main() {

 sources := map[string]time.Duration{
  "Локальная БД":    100 * time.Millisecond,
  "Облачная БД":     300 * time.Millisecond,
  "MySQL":       50 * time.Millisecond,
  "Redis":      500 * time.Millisecond,
  "Архивная БД":     800 * time.Millisecond,
 }
 
 fResult := make(chan string, 1)
 var once sync.Once
 var ff sync.WaitGroup
 
 for name, searchTime := range sources {
  ff.Add(1)
  
  go func(db string, duration time.Duration) {
   defer ff.Done()
   
   time.Sleep(duration)
   once.Do(func() {
    fResult <- fmt.Sprintf(" Найдено в %s (самый быстрый!)", db)
   })
  }(name, searchTime)
 }
 
 go func() {
  ff.Wait()
  close(fResult)
 }()
 
 fmt.Println(fResult)
}
