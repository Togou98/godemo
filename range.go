package main

import "fmt"
func main(){
		val :=1
arr := make([]int,100,100)
for i,_ := range arr{
					arr[i] = val
		fmt.Printf("%d : %d  |",i,arr[i])

			val++
	}
	sum := 0
	for  _,val:= range arr{
			sum += val
	}
	fmt.Printf("1-100 total =:%d \n",sum)
	}

