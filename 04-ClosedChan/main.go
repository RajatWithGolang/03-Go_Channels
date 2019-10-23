package main


import (
	"fmt"
)
func greet(c chan string){
	  <- c   //step 2 reading data from the channel which is available
	  <- c    // step 3 blocking the goroutine untill it gets data to read
	          // go scheduler launch main goroutine again
}
func main(){
	fmt.Println("Main Started")
	c:= make(chan string) // creating a channel
	go greet(c)  
	c <- "Rajat" //step 1 passing data to channel
	close(c) // step 4 main goroutine close the channel
	c <- "Rawat" // trying to pass the data to a closed channel but it'll panic at runtime
	fmt.Println("Main Stopped")
}