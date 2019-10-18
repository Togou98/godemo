package main
import "fmt"


func main(){
	i := 100
	pi := &i
	z := 50
	pz := &z
	fmt.Println(*pi, *pz)
	pi = pz
	if pi == pz{

	fmt.Println(*pz)
	}
}
