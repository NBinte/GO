package main

import "fmt"

func main(){
	mp := map[string][]int{"a": {1, 2, 3}}
	mp["b"] = []int{4, 5, 6}
	delete(mp, "b")

	value, check := mp["b"]
	fmt.Println(value, check)
}