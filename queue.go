package main
import "fmt"
import "time"
func main(){
		queue := make(chan int , 3)
	

		go func(){
				
				for  goods := range queue{
						fmt.Printf("get goods : %d \n",goods+1)
						time.Sleep(300 * time.Millisecond)


				}
				
		}()


		for i:=0; i < 20 ; i++{
			queue <- i
			fmt.Printf("Send goods:%d \n"  ,i+1)
	}
	
	close(queue)
	
}

