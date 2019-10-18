package main
import(
		"fmt"
		"os"
//		"errors"
)
func main(){
		//str := errors.New("hello")
		fd , err := os.Open("/etc/passwd")
		if err != nil{
			fmt.Println(err)

	}else {
			fmt.Println("file open ok")
			fmt.Printf("%T\n",fd)
			fd.Close()
	}
}
