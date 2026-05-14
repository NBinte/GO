package main

import "fmt"

func main(){
	
	a := "ro"
	
	switch {
	case a == "yes", a == "no", a == "ro":
		fmt.Println("a is yes, no or ro")
	default:
		fmt.Println("default")
	}
}