package main
import (
	"fmt"
	"time"
)

var chc = make (chan int)
var chs = make (chan int)
func counter(c chan int){
	for i:=0; ;i++{
		c <- i
		time.Sleep( 100 * time.Millisecond)
		continue
	}
}

func square(a chan int,b chan int){
	for {
	i := <-a
	b <- (i * i)
	time.Sleep( 100 * time.Millisecond)
	}
}
func printer(c chan int){
	for{
		fmt.Println(<-c)
	time.Sleep( 50 * time.Millisecond)
	}
}

func main(){
go counter(chc)
go square(chc,chs)
go printer(chs)


for{}

}
