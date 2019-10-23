package main

import (
	"fmt"
	"sync"
	"time"
)
func foo(wg *sync.WaitGroup,i int){
    time.Sleep(2* time.Millisecond)
	fmt.Printf("Foo Printed at instance %d\n",i)
	wg.Done()
}

func main(){
	var wg sync.WaitGroup

	for i:=0;i<5;i++{
		wg.Add(1)
		foo(&wg,i)
	}
	
	wg.Wait()
}