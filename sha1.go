package main

import (
		"fmt"
		"crypto/sha1"
)

func main(){
		 s := "helloworld"
		 h := sha1.New()
		 h.Write([]byte(s))
		 bs:= h.Sum(nil)
		 fmt.Printf("%x \n",bs)
 }

