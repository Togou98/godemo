package main
import "fmt"

const boolingF  =212.1

func main(){
	var f =boolingF
	var c float64 = (f - 32 )* 5 / 9
	fmt.Println("boiling point = %g or %v ",f,c)
}

