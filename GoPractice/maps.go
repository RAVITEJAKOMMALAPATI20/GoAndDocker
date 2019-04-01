package main

import "fmt"

func main(){
	studentage := make(map[string] int)
	studentage["Ravi"]=35
	fmt.Println(studentage["Ravi"])
		fmt.Println(len(studentage))

}