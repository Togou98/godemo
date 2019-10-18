package main

import "fmt"


func main(){
		ch := make(chan string)
		go func(){
				ch <-"helloworld"
		}()
		x := <-ch
		fmt.Println("go channel :",x)
}

