package main
import "fmt"
import "time"

func main(){
	c := make(chan int )
	go chwrite(c)
	go chread(c)
	time.Sleep( 5 * time.Second)

}

func chwrite(c chan int){
	x := 1000
	c <-x
}
func chread(c chan int){

	 y := <-c
	fmt.Printf("GOCHA data form channel C :%d\n",y)
}

