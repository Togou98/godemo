package main

import(
		"fmt"
		"net"
)


func main(){

		ln , err := net.Listen("tcp","localhost:9999")
		if err != nil {
				panic("listen error")
		}for {
				conn , err := ln.Accept()
				if err != nil  return
				fmt.Printf("%T\n",conn)

}

