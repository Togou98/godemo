package main

import "fmt"
import "time"

func main(){
	a := []int{10,3,46,87,546,432,654,3333}
	for _ , i := range a{
		go sleepsort(i)
	}
	//fmt.Printf("%T",time.Millisecond)
	time.Sleep(1 * time.Minute)
}



func sleepsort(i int){
	time.Sleep(time.Duration(i)* time.Millisecond)
	fmt.Printf("%d  ",i)
}

