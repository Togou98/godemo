package main

import(
		"os"
		"strings"
		"fmt"
)
func main(){

		os.Setenv("FOO","1")
		fmt.Println("FOO:",os.Getenv("FOO"))
		fmt.Println("bin:",os.Getenv("PATH"))
		for _ ,e := range os.Environ(){
				pair := strings.Split(e,":")
				fmt.Println(pair[0])
		}
}

