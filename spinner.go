package main
import( "fmt"
	"time"
)
func spin(delay time.Duration){
	for{
		for _,t :=range `-\|/`{
			fmt.Printf("\r%c",t)
			time.Sleep(delay)
		}
	}
}
func fib(i int) int {
	if i < 2 {
		return i
	}
	return fib(i-1) + fib (i-2)
}
func main(){
	go spin(100 * time.Millisecond)
	const n = 45
	fib45 := fib(n)
	fmt.Printf("\r Fibonacci(%d) = %d\n",n,fib45)
}

