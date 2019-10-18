package main

import(
		"fmt"
		"os/exec"
)
func main(){

		datecmd := exec.Command("date")

		datecmdoutput , err := datecmd.Output()
		if err != nil {
				panic(err)
		}

		fmt.Println(" > date")
		fmt.Println(string(datecmdoutput))
}

