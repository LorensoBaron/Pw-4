package main

import (
 "fmt"
 "sync"
 "time"
)

func main() {
 jobs := make(chan int, 10)
 

 var ff sync.WaitGroup
 for i := 1; i <= 3; i++ {
  ff.Add(1)
  go func(workerID int) {
   defer ff.Done()
   
   for job := range jobs {
    fmt.Printf("Воркер %d обработал задание %d\n", workerID, job)
    time.Sleep(100 * time.Millisecond)
   }
  }(i)
 }
 for i := 1; i <= 10; i++ {
  jobs <- i
 }
 close(jobs)

 ff.Wait()
 fmt.Println(" Все задания обработаны")
}
