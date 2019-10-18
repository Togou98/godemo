package main

import (
		"fmt"
		//"io"
		"os"
		"bufio"
)


func main(){

		s := bufio.NewScanner(os.Stdin)
		for s.Scan(){
				fmt.Println(s.Text())
				fmt.Printf("%T\n,",s.Text())
		}
}

