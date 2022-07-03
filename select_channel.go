package main

import (
	"fmt"
	"time"
)

func main(){
	out1 := make(chan int)
	out2 := make(chan int)
	
	go slow(3,out1)
	go fast(4,out2)
	
	select{
		case res := <- out1:
			fmt.Println("slow is finished first, result : ",res)
		case res := <- out2:
			fmt.Println("fast is finished first, result : ",res)
	}
}

func slow(result int, out1 chan int){
	result = result*2
	time.Sleep(15*time.Millisecond)
	out1 <- result
	
}

func fast(result int, out2 chan int){
	result = result*2
	time.Sleep(5*time.Millisecond)
	out2 <- result
}