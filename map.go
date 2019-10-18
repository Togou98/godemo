package main
import "fmt"

func main(){
type key string
type value int
		m := make(map[key]value)
		m["GJY"] = 21
		m["SL"] = 20

		fmt.Println(m["SL"])

		v1 := m["GJY"]
		fmt.Println(v1)
		m["temp"] = 100
		fmt.Printf("delete m[temp]:%d\n",m["temp"])
		delete(m,"temp")
		fmt.Println(m)
		k,v := m["hello"]
		fmt.Println(k ,v)
}
