package main

import(
		
		"time"
)

func main(){
		c1 := make(chan string)
		c2 := make(chan string)
		
		go func(){
				time.Sleep(3 *time.Second)
				c1 <- "I 'm routine 1 , I'm done"
		}()
		go func(){
				time.Sleep(2 *time.Second)
				c2 <-"I'm go routine 2, I'm done"
		}()

		
		select{

		}

}

