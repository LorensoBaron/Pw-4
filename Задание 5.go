package main

import (
	"crypto/md5"
	"fmt"
	"sync"
)

func main() {
	data := []string{"aaaa", "bbbb", "vvvv", "cccc", "rrrr", "gggg"}
	limit := make(chan bool, 3)
	var ff sync.WaitGroup

	for _, text := range data {
		ff.Add(1)
		limit <- true
		go func(t string) {
			defer ff.Done()
			defer func() { <-limit }()
			hash := md5.Sum([]byte(t))
			fmt.Printf("Хеш '%s' : %x\n", t, hash)
		}(text)
	}
	ff.Wait()
}
