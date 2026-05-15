package main

import "fmt"

func main(){
	sl := []string{"hello", "world", "withRo"}

	for i, value := range sl{
		fmt.Println(i, value)
	}
}