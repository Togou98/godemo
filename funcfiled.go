package main

import "fmt"

func f()int {

	return 1
}

func g( data int) int {
	return data +f()
}

var x int = 21
func main(){
	if x := f();  x == 0 {
		fmt.Println(x)
	}else if y := g(x) ; x == y{
		fmt.Println(x,y)
	}else {
		fmt.Println(x,y)
	}
	fmt.Println(x,y)
}

