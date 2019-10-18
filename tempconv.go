package main

import( 
	"tmpconv"
	"fmt"

)
func main(){

	var ff tmpconv.Ftmp =135.2

	fconv := tmpconv.Ftoc(ff)
	fmt.Println(fconv.String())
}

