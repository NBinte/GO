package main

import(
	"fmt"
	"strconv"
)

func main(){
	
	x := "1234"
	y, err := strconv.Atoi(x)
	fmt.Println(y,err)
}