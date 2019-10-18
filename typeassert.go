package main

import "fmt"

func typeid(i interface{}){
	switch i.(type){
	case int :
		fmt.Println("Integer")
	case string :
		fmt.Println("String")
	case bool :
		fmt.Println("Boolean")
	default : 
		panic("Unknow type")
	}
}
func main(){
	x := 100000
	typeid(x)
}
