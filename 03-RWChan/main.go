package main


import (
	"fmt"
)
func greet(c chan string){
       fmt.Println("Hello ",<-c) // reading data from the channel
}
func main(){
	fmt.Println("Main Started")
	c:= make(chan string) // creating a channel
	go greet(c) //passing channel to goroutine
	c <-"Rajat" // passing data to channes
	fmt.Println("Main Stopped")
}