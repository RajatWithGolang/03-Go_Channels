package main

import (
   "fmt"
)

func greet (roc <-chan string){
	fmt.Println("Hello",<- roc)
	
}

func main(){
  fmt.Println("Main Started")
  c:= make(chan string)
  go greet(c)
  c <- "Rajat"
  
  fmt.Println("Main Terminated")
}