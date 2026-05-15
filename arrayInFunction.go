package main

import "fmt"

func main(){
	
	arr := [...][2]int{{1, 2}, {2, 2}, {3, 2}}
	test(arr)
	fmt.Println(arr)
}

func test(arr [3][2]int){
	arr[0] = [2]int{100, 100}
}