package main

import "fmt"

func main(){
	
	// pointer -> arr[1]
	// length -> 4
	// capacity -> 4

	arr := [5]int{1, 2, 3, 4, 5}
	sl := arr[1:3]
	sl = sl[:4]
	fmt.Println(sl, len(sl), cap(sl))
}