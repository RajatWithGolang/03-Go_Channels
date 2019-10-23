package main

import (
	"fmt"
	//"time"
	"sync"
	
)

var wg sync.WaitGroup

func foo(){
	for i:=0;i<5;i++{
		fmt.Println("Foo is Printed")
	}
	wg.Done()
}
func bar(){
	for i:=0;i<5;i++{
		fmt.Println("Bar is Printed")
	}
	
}

func main(){
	fmt.Println("Main Started")
	wg.Add(1)
	go foo()
	wg.Wait()
	bar()
	fmt.Println("Main Stopped")


}