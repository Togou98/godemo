package main

import "fmt"

func send( ch chan<- string ,msg string){
		ch <- msg
}
func receive( get <-chan string , from  chan<- string){
		msg := <-get
		from <- msg
}

func main(){
		ch1 := make(chan string ,1)
		ch2 := make(chan string ,1)
		go send(ch1, "helloworld")
		go receive(ch1 , ch2)
		fmt.Println(<-ch2)
}

