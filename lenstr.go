package main
import"fmt"
var str string = "世界"

func main(){
	n := 0
	for range str{
	n++
	}
	fmt.Println(str ,len(str),n,"\ufffd")
}
