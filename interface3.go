package main
import "fmt"


type iron struct {
	price float64
}

type gold struct {
	price float64
}
type silver struct {
	price float64
}



type materials interface {
	maxhave(m float64) float64
}

func (i iron) maxhave(m float64) float64{
	return m / i.price
}
func (g gold) maxhave(m float64)float64{
	return m / g.price 
}
func (s silver) maxhave(m float64) float64{
	return m / s.price
}



func main(){
var hav materials
var coins float64 = 100000.0
//var i iron = iron{price:5.3}
var g gold =gold{420.5 }
//var s silver = 28.9

hav = g
fmt.Printf("max can purchase%.3f material\n ",hav.maxhave(coins))
fmt.Printf("%T\n",hav)

}
