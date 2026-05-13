package main

import(
	"fmt"
	"strconv"
)

func main(){
	
	x := "1234"
	y, err := strconv.ParseInt(x, 10, 0)
	fmt.Println(y,err)
}