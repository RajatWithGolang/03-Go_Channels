package main


import (
	"fmt"
	"sync"

)
var i int
func foo(wg *sync.WaitGroup,m *sync.Mutex){
	m.Lock()
	i = i + 1
	m.Unlock()
	wg.Done()
}
func main(){
  var wg sync.WaitGroup
  var m sync.Mutex

  for i:=1; i<=1000;i++{
	  wg.Add(1)
	  go foo(&wg,&m)
  }
  wg.Wait()
  fmt.Println("value of i after 1000 operations is", i)
}