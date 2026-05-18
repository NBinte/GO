package main

import "fmt"

func callFunc(callable func(int) int) int{
	return callable(10)
}

func doubleNumber(num int) int{
	return 2 * num
}

func tripleNumber(num int) int{
	return 3 * num 
}

func main(){
	value := callFunc(doubleNumber)
	fmt.Println(value)
}