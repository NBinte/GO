package main

import "fmt"

func main(){
	
	a := -2
	
	switch {
	case a < -1: 
		fmt.Println("a is less than -1")
		fallthrough
	case a < 0: 
		fmt.Println("a is less than 0")
		fallthrough
	case a < 1: 
		fmt.Println("a is less than 1")
	default:
		fmt.Println("default")
	}
}