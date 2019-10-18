package main

import "fmt"
import "time"


func main(){
		i := 2 
		fmt.Print("Write ",i," as ")
		switch i {
		case 1:
				fmt.Println("one")
		case 2:
				fmt.Println("two")
		case 3:
				fmt.Println("three")
		}

		switch time.Now().Weekday(){
		case time.Saturday , time.Sunday:
				fmt.Println("You can take a rest")
				default : 
				fmt.Println("Go to fucking company")
		}
		t := time.Now()
		switch {
		case t.Hour() > 12 :
				fmt.Println("Now is afternoon")
				default :
				fmt.Println(" Now is before noon")
		}
}
