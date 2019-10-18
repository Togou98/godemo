package main
import(
	"fmt"
	"sync"
)

type money struct {
	rmb int
	m sync.Mutex
}

func main(){

	r := money{}
test(r)
}



func(m money) add(i int){
	m.m.Lock()
	m.rmb += i
	m.m.Unlock()
}
func(m money) show(){
	fmt.Printf("Account money: %d\n",m.rmb)
}

func test(m money){
		for{
			go m.add(100)
			go m.show()
		}
	}
