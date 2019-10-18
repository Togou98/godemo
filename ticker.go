package main

import "fmt"
import "time"


func main(){
		tik := time.NewTicker(200 * time.Millisecond)

		go func() {
				        for t := range tik.C {
								            fmt.Println("Tick at", t)
											        }
													    }()
		time.Sleep(1500 * time.Millisecond)
		tik.Stop()
		fmt.Println("ticker stopped")
}

