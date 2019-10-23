package main

import (
	"fmt"
	"time"
	"sync"
)
func sqrWorker(wg *sync.WaitGroup,tasks <- chan int,results chan <- int, instance int){
  for num := range tasks {
		time.Sleep(time.Millisecond) // simulating blocking task
		fmt.Printf("[worker %v] Sending result by worker %v\n", instance, instance)
		results <- num * num
	}
	wg.Done()
}

func main(){
	var wg sync.WaitGroup
	tasks := make(chan int,10)
	results := make(chan int,10)

	for i:=1; i<=3; i++{
		wg.Add(1)
       go sqrWorker(&wg,tasks,results,i)
	}

	for i:=0;i<5;i++{
       tasks <- i * 2
	}

	fmt.Println("main wrote 5 tasks")

	close(tasks)
	wg.Wait()
    for i := 0; i < 5; i++ {
		result := <-results // blocking because buffer is empty
		fmt.Println("[main] Result", i, ":", result)
	}
}