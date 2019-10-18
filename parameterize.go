package main 
import "fmt"
import "strings"


func add1( r rune ) rune { return r+1 }



func main(){
	fmt.Println(strings.Map(add1,"Halo-666"))
}

