package main

import("net/http"
		"fmt"
		"bufio"
)

func getdata(w  http.ResponseWriter , req *http.Request){
		resp , err := http.Get("www.h1z1.com")
		if err != nil{
				panic("http get error")
		}
		defer resp.Body.Close()
		scan := bufio.NewScanner(resp.Body)
		for i:= 0 ; scan.Scan() && i <10; i++{
				fmt.Println(scan.Text())
		}
}
func main(){
		http.ListenAndServe("localhost:9999",nil)
		http.HandleFunc("/nice",getdata)
}
