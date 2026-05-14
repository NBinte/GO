package main

import "fmt"

func main(){
	
	arr := [...][2]int{{1, 2}, {2, 2}, {3, 2}}
	arr[0] = [2]int{10, 11}
	fmt.Println(arr)
}