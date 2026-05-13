package main

import "fmt"

func main(){
	
	y := 3.312456879894510

	x := fmt.Sprintf("\"%10.2f%%", y)

	fmt.Println(x, x)
}