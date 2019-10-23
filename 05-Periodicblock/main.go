package main

import "fmt"

func square(c chan int){   
	 for i:=0 ;i<10; i++{
		 c <- i*i   //here we are writing data to channel and blocking the condition
	 }
	 close(c) // after writing all the data we are closing the channel
	          // If you don’t close the channel in for range loop, program will throw deadlock fetal error in runtime.
}

func main(){
	fmt.Println("Main Started")
	c:= make(chan int)
	go square(c)
    for value := range c {  // periodic block untill it reads data from goroutin
		fmt.Println(value)
	}
	fmt.Println("Main Terminated")

}