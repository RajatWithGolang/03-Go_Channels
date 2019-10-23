package main

import (
   "fmt"
)

func square(c chan int){
	for i:=0; i<=3;i++{
		num := <- c
		fmt.Println(num*num) 
	}
}

func main(){
  fmt.Println("Main Started")
  c:= make(chan int,3)
  go square(c)
  c <- 1
  c <- 2
  c <- 3
  fmt.Println("Main Terminated")
}