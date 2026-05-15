package main

import "fmt"

func main(){
	// pointer
	// length
	// capacity
	
	arr := [5]int{1, 2, 3, 4, 5}
	sl := arr[:3]
	fmt.Println(sl, len(sl), cap(sl))
}