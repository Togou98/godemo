package main
import "fmt"

func intseq() func()int{
		i := 0
		return func() int{
				i++
				return i
		}
}

func main(){
		nextint := intseq()
		fmt.Println(nextint())
		fmt.Println(nextint())
}
