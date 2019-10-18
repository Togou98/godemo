package main

import "fmt"

type hello int


type bag struct{
	X hello
	Y hello
}


func main(){

	var foo bag = bag{Y:10 ,X:11}
	fmt.Println(foo.X ,foo.Y)

}

