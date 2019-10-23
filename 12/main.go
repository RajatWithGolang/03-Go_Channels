package main

import "fmt"

func main(){
	fmt.Println("Main Started")
	c:= make(chan string)

	go func(c chan string){
		fmt.Println("Hello " + <-c + "!")
	}(c)

   c <- "Rajat"

   fmt.Println("Main Terminated")
}

