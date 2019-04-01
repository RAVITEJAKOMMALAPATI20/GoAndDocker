package main

import "fmt"


func main(){
	a,b := 5,6
	fmt.Println(sum(a,b))
	fmt.Println(factorinl(a))

}
func sum(a int ,b int) int{
	return a+b;
}


func factorinl(num int) int{
	if num==1{return 1}
	return num*factorinl(num-1)
}