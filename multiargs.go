package main
import "fmt"

func sum (val ...int) int{
	sum :=0
	for _,v := range val{
		sum += v
	}
	return sum
}
func main(){
	arr := make([]int,100,100)
	for i:=0;i<100;i++{
		arr[i]= i+1
	}

	fmt.Println(sum(arr...))

	a:= []int{0,1,2,3,4,5}
	b:= make ([]int,6)
	copy(b[1:6],a[:])
	fmt.Println(b)
	fmt.Println("b cap is :%v  len is:%v" ,cap(b),len(b))

}

