package main

import "fmt"

func main(){
	sl := []string{"hello", "world", "withRo"}
	test(sl)
	fmt.Println(sl)
}

func test(arr []string){
	arr[0] = "Changed"
}