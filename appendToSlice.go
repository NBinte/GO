package main

import "fmt"

func main(){
	sl := []string{"hello", "world"}
	
	for x:= 0; x < 10; x++{
		sl = append(sl, "withRo")
		fmt.Println(sl, len(sl), cap(sl))
	}

	fmt.Println(sl)
}