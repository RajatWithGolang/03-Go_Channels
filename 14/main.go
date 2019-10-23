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
	fmt.Println("service1() started", time.Since(start))
	ch <- "Hello from Service 1"
}

func Service2(ch chan string){
	fmt.Println("service2() started", time.Since(start))
	ch <- "Hello from Service 2"
}

func main(){
	 fmt.Println("Main started at", time.Since(start))

	  chan1 := make(chan string)
	  chan2 := make(chan string)
	  
	  go Service1(chan1)
	  go Service2(chan2)

	  select { // No-blocking because of Default case
	  case res := <- chan1 : 
		fmt.Println("Response from Service 1",res,time.Since(start))
	  case res := <- chan2 :
		fmt.Println("Response from Service 2",res,time.Since(start))
	  default:
		fmt.Println("No Response Received",time.Since(start))
	  }	  
	fmt.Println("Main Stopped at",time.Since(start))
}