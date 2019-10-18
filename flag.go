package main

import(
		"flag"
		"fmt"
)

func main(){

		wordptr := flag.String("print","null", "a string")
		flag.Parse()
		fmt.Println(*wordptr)

}

