package main

import (
   "fmt"
)

func sender (c chan int){
	c <- 1
	c <- 2
	c <- 3
	c <- 4
}

func main(){
  fmt.Println("Main Started")
  c:= make(chan int,3)
  go sender(c)
  
  fmt.Printf("length of channel c is %v and capacity is %v\n",len(c),cap(c))

  for value := range c{
     fmt.Printf("length of channel c after  %v read is %v\n",value,len(c))
  }
  fmt.Println("Main Terminated")
}