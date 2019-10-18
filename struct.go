package main

import "fmt"
import "time"

type t struct{
		name string
		age int
}
func main(){
		GJY := t{"郭金宇",21}
		pp := &GJY
		fmt.Println(pp.name)
		fmt.Println(time.Now())

}
