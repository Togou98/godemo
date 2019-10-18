package main
import ("fmt" 
		"time" 
		"sync/atomic"
)

func main() {
	var cnt int32

	for i :=0; i < 50 ; i++ {
			go func(){
					for {
							atomic.AddInt32(&cnt,1)

							time.Sleep(time.Millisecond)
					}
			}()
	}
time.Sleep(time.Second * 3)

newcnt := atomic.LoadInt32(&cnt)
fmt.Println("ops:" ,cnt,"   ",newcnt)
}


