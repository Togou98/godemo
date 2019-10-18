package main
import "fmt"

func HasPrefix (s ,prefix string) bool {
	return len(s) >= len(prefix) && s[ : len(prefix)] == prefix 
}
func Contains (s , substr string) bool {
	for i := 0 ;i < len(s) ; i++ {
		if HasPrefix(s[i:] ,substr) {
			return true 
		}
	}
	return false 
}

func main(){
	a := "helloworld"
	b := "world"
	fmt.Println( Contains( a , b))
}

