package main
import "fmt"

func reverse(arr []int) {
	for i ,j := 0, len(arr)-1 ;i < j ; i , j  = i+1 ,j-1{
		arr[i], arr[j] = arr[j] , arr[i]
	}
}


func main(){
	a :=  []int {1,3,5,7,9,11}
	reverse(a)
	fmt.Println(a[:])
}

