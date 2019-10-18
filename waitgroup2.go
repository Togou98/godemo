package main
import(
		"fmt"
		"sync"
		"time"
)

func sleep(num int, wg *sync.WaitGroup){
		time.Sleep(1 * time.Second)
		fmt.Printf("coroutine %d \n",num)
		wg.Done()
}

func main(){
		var wg sync.WaitGroup
		for i := 0 ; i < 5 ; i++ {
				wg.Add(1)
				go sleep(i,&wg)
		}
		wg.Wait()
}


