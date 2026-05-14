package main

import "fmt"

func main(){
	
	str := "whatever"

	for idx := 0; idx < len(str); idx++{
		fmt.Printf("%c", str[idx])
	}
}