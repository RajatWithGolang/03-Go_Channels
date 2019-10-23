package main

import "fmt"

func main(){
	c:= make(chan int)
	fmt.Printf("%T\n",c)
	fmt.Printf("%v\n",c)
}