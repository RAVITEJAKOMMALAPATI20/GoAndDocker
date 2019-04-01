package main

import ("fmt"
"log"
"os"
"io/ioutil")


func main(){
	file,err := os.Create("sample.txt")

	if err !=nil{
		log.Fatal(err)
	}
	file.WriteString("Hi I am raviteja Kommalapati written code in Go")
	file.Close()
	stream,err := ioutil.ReadFile("sample.txt")
	if err !=nil{
		log.Fatal(err)
	}
	line :=string(stream)
	fmt.Println(line)

}