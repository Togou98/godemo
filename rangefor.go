package main
import(
	"fmt"
	"os"
)

func main(){
	var s , str string 

	for _ , arg := range os.Args[1:]{
		s +=  str + arg
		str = " "
	}
	fmt.Println(s)
}
