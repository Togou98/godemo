package main
import (
		"fmt"
	str	"strings"
)
func main(){

		s1 := "I am 郭金宇"
		fmt.Println("s1 have space :",str.Count(s1,"t"))
		s2 :=  str.Replace(s1," ","%20",-1)
		fmt.Println(s2)
		fmt.Printf("s1 contained 我的名字么？",str.Contains(s1,"郭金宇"))
		fmt.Println("/ ascii code is:","/"[0])
}


