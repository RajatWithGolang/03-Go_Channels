package main


import (
	"fmt"
	"sync"

)
var i int
func foo(wg *sync.WaitGroup){

	i = i + 1

	wg.Done()
}
func main(){
  var wg sync.WaitGroup
 

  for i:=1; i<=1000;i++{
	  wg.Add(1)
	  go foo(&wg)
  }
  wg.Wait()
  fmt.Println("value of i after 1000 operations is", i)
}