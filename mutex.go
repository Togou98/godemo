package main

import (
		"fmt"
		"sync"
)


func main(){
		var  mtx sync.Mutex
		var wg sync.WaitGroup
		i := 0
		pi := &i
		go func(p *int){
				wg.Add(1)
				for  i:=0; i<100; i++{
						
						mtx.Lock()
						if *p >= 100 {break}
						*p++;
						mtx.Unlock()
						fmt.Println("thread 1 val is :",*p)
				}
				wg.Done()
		}(pi)
	go func(p *int){
			wg.Add(1)
				for  i:=0; i<100; i++{
						mtx.Lock()
						if *p >= 100 {break}
						*p++;
						mtx.Unlock()
						fmt.Println("thread 2 val is :",*p)
				}
		}(pi)

wg.Wait()
fmt.Println("???")
for{}

}

