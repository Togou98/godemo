package main
import "fmt"
const (
	x = iota
)


func returnfunc() func() int{
	y := x
	return func() int {
		y++
		return y + y
	}
}

func main(){
	fun := returnfunc()

	fmt.Println(fun(x+10))
}
