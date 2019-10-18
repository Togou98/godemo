package main
import "fmt"


func remove (a []int ,pos int ) []int{
	copy (a[pos:] ,a[pos+1:])

	return a[:len(a)-1]
}
func main(){


	var a []int  = []int{1,3,5,7,9}
	fmt.Println(remove(a,3))
}




