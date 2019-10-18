package main
import"fmt"
//import"time"


const Pi = 3.14


type Cir struct {
	Radius float64
}
func (c *Cir) Area() float64 {
	return c.Radius * c.Radius * Pi
}

func main(){
	var c Cir = Cir{2.0}
	fmt.Println((&c).Area())
	fmt.Printf("%T\n",c.Area())

}

