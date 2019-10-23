package main

import (
   "fmt"
)

func main(){
	c := make(chan int,3)
	c <- 1
	c <- 2
	c <- 3
	
	close(c)

	for value := range c{
		fmt.Println(value)
	}
}