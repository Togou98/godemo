package main


import "fmt"

func sum(arr ...int) int {
		fmt.Println(arr," ")
	total := 0
	for _,val := range arr {
		total += val
		}
		return total
}

func main(){
		fmt.Println(sum(1,2,3,4,5,6,7,8,9,10))
}
