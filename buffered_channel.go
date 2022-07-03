package main

import (
	"fmt"
	
)

func main(){
	
	out := make(chan int,3)
	out <- 1
	out <- 2
	close(out)
	

	for v := range out{
		fmt.Println(v)
	}
}