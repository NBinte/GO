package main

import "fmt"

func add(num1 int, num2 int)int{
	return num1 + num2
}

func main(){
	value := add(1, 2)
	fmt.Println(value)
}