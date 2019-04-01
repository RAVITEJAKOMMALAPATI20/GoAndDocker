package main

import "fmt"


func main(){
	var num[3] int
	num[0]=4
	num[1]=3
	num[2] =5
	for _,value:=range num{
		fmt.Println(value)
	}

	for i,value:=range num {
		fmt.Println(value,i)
	}
}