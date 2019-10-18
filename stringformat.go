package main

import "fmt"
//import "os"

type point struct {
		x,y int
}

func main(){

		p:= point{x:1,y:2}

		fmt.Printf("%v\n",p)
		fmt.Printf("%+v \n",p)
		fmt.Printf("%#v \n",p)
		fmt.Printf("%T\n",p)
		fmt.Printf("%t \n",true)
		fmt.Printf("%d \n",123)
}

