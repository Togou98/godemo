package main

import(
		"fmt"
		//"bufio"
		"io"
		"io/ioutil"
		"os"
)

func check(e error){
		if e != nil{
				panic(e)
		}
}
func main(){
		dat ,err := ioutil.ReadFile("/home/togou/go/hello.txt")
		check(err)
		fmt.Println(string(dat))

		f , err := os.Open("/home/togou/go/hello.txt")
		check(err)

		b1 := make([]byte ,5 )
		n1 , err := f.Read(b1)
		check(err)
		fmt.Printf("%d bytes: %s \n",n1,string(b1[:n1]))

		o2 ,err := f.Seek(5,0)
		check(err)

		b2 := make([]byte ,2)
		n2 ,err := f.Read(b2)
		check(err)

		fmt.Printf("%d bytes @%d;",n2,o2)
		fmt.Printf("%v \n",string(b2[:n2]))

		o3, err := f.Seek(6,0)
		b3 := make([]byte, 2)
		    n3, err := io.ReadAtLeast(f, b3, 2)
			    check(err)
				    fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))
					f.Close()
			}


