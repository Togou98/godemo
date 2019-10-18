package main

import "fmt"


func Sum(i int)float64{
		p , q , sum := 1.0,2.0,0.0
		for ;i > 0; i--{
			sum += q/p
			p ,q = q ,p+q
		} 
		return sum
}
func main(){
		fmt.Println(Sum(2))
}


