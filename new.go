package main
import "fmt"
func main(){
	var p = new (int)
	*p = 100

	fmt.Println(*p)
}
