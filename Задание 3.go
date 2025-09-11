package main
import (
	"fmt"
	"time"
)
func main() {
	tick := time.Tick(200 * time.Millisecond)
	request := make(chan int, 15)
	result := make(chan string, 15)
	
	for i := 1; i <= 15; i++ {
		request <- i 
	}
	close(request)
	go func() {
		for req := range request {
			<- tick 

            time.Sleep(50 * time.Millisecond)
		res := fmt.Sprintf("Запросик %d успешно обработан в %v", req, time.Now().Format("11:17:08.000"))
		result <- res
		}
		close(result)
	}()
		 fmt.Println("Результаты запросов:")
    for res := range result {
    fmt.Println(res)
	}
}
