package main


import "fmt"

func main(){

	person := Person{firstname:"ravi",lastname:"teja"}
fmt.Println("Full name of a person is :"+person.fullname())
}
type Person struct{
 firstname string
 lastname string
}




func (person *Person) fullname() string{
	return  person.firstname+person.lastname
}