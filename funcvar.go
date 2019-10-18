package main
import "fmt"

func main(){
	f := add
	fmt.Println(f(5,3))

	f = product 
	fmt.Println(f(5,3))

}

func add(x , y int ) int {
	return x + y 
}
func product(x , y int ) int {
	return x * y 
}

