package main

import "fmt"

func main(){
	
	arr := [...][2]int{{1, 2}, {2, 2}, {3, 2}}
	
	for i, value := range arr{
		fmt.Println(i, value)
	}
}