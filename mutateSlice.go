package main

import "fmt"

func main(){
	
	arr := [5]int{1, 2, 3, 4, 5}
	sl := arr[1:]
	sl[0] = 100
	fmt.Println(sl, arr)
}