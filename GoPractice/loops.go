package main

import "fmt"

func main(){
	for i:= 0;i<5;i++{
       fmt.Println(i)
       j := 0
       for j<5 {
       	    fmt.Println(j)
       	    j++
       }
	}
	
}