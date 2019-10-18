package main

import("time"
"fmt"
)


func main(){
		timer1:= time.NewTimer(3 * time.Second)
		<-timer1.C
		fmt.Println("timer 1 expired")
}

