package main

import(
		"fmt"
		"math/rand"
		"sync/atomic"
		"time"
)

type read struct {
		key int
		resp chan int
}
type wrtie struct {
		key int
		val int 
		resp chan bool
}
func main(){
		var readOps int32
		var writeOps int32

		reads := make(chan read)
		writes := make(chan write)

		go func(){
				var state = make(map[int]int)
				for {
						select {
						case read := <-reads:
								read.resp <- state[read.key]
						case write := <-writes:
								state[write.key] = write.val
								write.resp <- true
						}
				}
		}()

		for r := 0; r <100 ; r++{
				go func() {
						for { 
								rd := read{
										key : rand.Intn(5),
										resp : make(chan int)}
										reads <- read
										<-read.resp
										atomic.AddInt32(&readOps , 1)
										time.Sleep(time.Millisecond)
								}
						}()
						

