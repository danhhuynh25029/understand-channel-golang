package main

import "fmt"

func main(){
	// khoi tao channel
	value := make(chan int,3)

	defer close(value)
	//  tao ra 2 go routine
	go addTwo(1,value)
	go addTwo(2,value)
	
	// result := <- value
	// lay ra gia tri value
	fmt.Println(<-value)
}
func addTwo(number int,value chan int){
	// gan gia tri vao channel
	value <- number
}