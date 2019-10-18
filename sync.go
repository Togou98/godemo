package main

import(
		"fmt"
		"sync"
		"time"
)


func count(n *int,condvar *sync.Cond){
		for{	condvar.L.Lock()
				(*n)++
				condvar.L.Unlock()
				time.Sleep(100 * time.Millisecond)
				if *n > 10000 {
						condvar.Signal()
				}
		}
}
func main(){
		var mu sync.Mutex
		cond := sync.NewCond(&mu)
		num := new(int)
				go count(num, cond)
		cond.Wait()
		fmt.Println("now num value is :",*num)
}



