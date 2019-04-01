package main

import "fmt"


func main(){

fmt.Println(add(1,4,5,6,6,77,5))

}
func add(val ...int) int {
	sum :=0
	for _,value := range val {
		sum +=value
	}
	return sum
}