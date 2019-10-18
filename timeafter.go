package main
import (
		"fmt"
		"time"
)
func main(){

		ch := make(chan string , 1)
		go func(){
				time.Sleep(2 *time.Second)
				ch <- "I'm done"
		}()
		select {
		case  ok := <-ch :
				fmt.Println("that ok",ok)
		case <- time.After(3 * time.Second) :
				fmt.Println("that timeout")
		}

}
