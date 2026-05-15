package main

import "fmt"

func main(){
	mp := map[string]bool{}
	mp["b"] = false

	value, check := mp["b"]
	fmt.Println(value, check)
}