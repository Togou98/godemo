package main

import ("fmt"
		"time"
		"sync"
)

func show(str string ,wg *sync.WaitGroup){
		wg.Add(1)
		fmt.Println(str)
		time.Sleep(time.Second)
		wg.Done()
}

func main(){
 var wg sync.WaitGroup
  for i := 0 ;i<5;i++{
		  go show("helloworld",&wg)
  }
  wg.Wait()
  }
