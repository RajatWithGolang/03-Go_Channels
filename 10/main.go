package main

import "fmt"

func square(c chan int){
	fmt.Println("Reading Square")
		num:= <-c
		c <- num*num

}
func cube(c chan int){
	fmt.Println("Reading Cube")
		num:= <- c
		c <- num*num*num

}


func main(){
	sqrchan := make(chan int)
	cubchan := make(chan int)
	
	go square(sqrchan)
	go cube(cubchan)
	fmt.Println("Sent number to sqrchan")
	sqrchan <- 3
	fmt.Println("Main Resuming")
	fmt.Println("Sent number to cubChan")
	cubchan <- 4
	fmt.Println("Main Resuming")
	fmt.Println("Main reading from channels")
	sqrVal,cubVal := <-sqrchan ,<-cubchan
	sum := sqrVal+cubVal
	fmt.Println(sum)




}