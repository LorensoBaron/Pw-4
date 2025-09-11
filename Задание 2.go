package main

import (
    "fmt"
    "sync"
    
)
func main(){
	job := make(chan int, 10)
	result := make(chan int, 10)
	var ff sync.WaitGroup
	for i := 1 ; i <= 3; i ++ {
		ff.Add(1)
		go wor(i, job, result, &ff)
	}
	for j:=1; j <= 10; j++ {
		job<-j
	}
	close(job)
	ff.Wait()
	close(result)

	fmt.Println("result:")
	for res := range result {
		fmt.Println(res)
	}
}

func wor(id int, job <- chan int, result chan <- int, ff *sync.WaitGroup) {
defer ff.Done()
for job2:=range job{
	res := job2*job2
	result<-res
}
}
