package main
import"fmt"
import"time"
var tick  int = 0
func t1(c chan int){
	 x := 1
	 
	go timer()
	for {
			 c <- x 
		      <-c
			tick++
	 }
	 

}
func timer(){
		for{	times := tick
			t := time.Now()
			if time.Since(t) == time.Second{
					fmt.Printf("1s :%d",tick-times)
					times = tick
			}
	}
}

func main(){
		ch := make(chan int)
		 x:=1
		go  t1(ch)
	for{
			<-ch
			ch <-x
	}
}
