package main


func main(){
	var x  string = "hello"
	//var y string = "hello"
	var z  *string =&x
	if x ==*z {
			panic("ok")
	}

}

