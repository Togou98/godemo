package main
import(
		"fmt"
		"time"
)

func main(){
		str := "\|/"
		for {
				-,s := range str{
				fmt.Printf("%v",s)
				time.Sleep(time.Second)
				fmt.Printf("\b")
		}
}
}





