package main

import "fmt"

func add(num1 int, num2 int)(int, string){
	return num1 + num2, "hello"
}

func main(){
	value, str := add(1, 2)
	fmt.Println(value, str)
}