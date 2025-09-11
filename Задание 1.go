package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    var ff sync.WaitGroup
    ff.Add(1)

    go func() {
        defer ff.Done()
        for i := 1; i <= 5; i++ {
            fmt.Println(i)
            time.Sleep(1 * time.Second)
        }
    }()

    ff.Wait()
}
