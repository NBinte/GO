package main

import "fmt"

func main(){
	
	arr := [...][2]int{{1, 2}, {2, 2}, {3, 2}}
	
	for _, nested := range arr{
		for _, value := range nested{
			fmt.Println(value)
		}
	}
}