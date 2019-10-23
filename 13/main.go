package main

import (
	"fmt"
	"time"
)
var start time.Time
func init(){
	start = time.Now()
}

func Service1(ch chan string){ 
	
	ch <- "Hello from Service 1"
}
func Service2(ch chan string){
	
	ch <- "Hello from Service 2"
}

func main(){
	 fmt.Println("Main started at", time.Since(start))
	  chan1 := make(chan string)
	  chan2 := make(chan string)
	  go Service1(chan1)
	  go Service2(chan2)

	  select { // block the execution
	  case res := <- chan1 : 
		fmt.Println("Response from Service 1",res,time.Since(start))
	  case res := <- chan2 :
		fmt.Println("Response from Service 2",res,time.Since(start))
		  }
		  fmt.Println("Main Stopped at",time.Since(start))
}