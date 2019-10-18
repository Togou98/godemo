package main
import "fmt"
func main(){
	var i bool = false;
	fmt.Println(btoi(i))
	var j int = 1
	fmt.Println(itob(j))

}

func btoi( i bool) int {
	if i {
		return 1
	} else {
	return 0  }

}

func itob( i int) bool {
	return  i > 0
}
