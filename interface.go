package main

import "fmt"

type inter interface {
	Print()
}

type info struct {
	name string
	age int
}

func (i info) Print(){
	fmt.Printf("This person'name is: %s ,age is: %d\n", i.name, i.age)
}

type rootinfo struct{
	name string 
	age int
}

func (r rootinfo) Print(){
	fmt.Printf("This Person is root ,can't get info from normal users")
	panic("system carshed")
}

func main(){
var cxx info = info{"GJY",21}
var x inter;
x = cxx
	if t,ok := x.(info);ok{
		fmt.Printf("%T \n",t)
	}
useinterface(x)


}


func useinterface ( i  inter){
	i.Print()
}

