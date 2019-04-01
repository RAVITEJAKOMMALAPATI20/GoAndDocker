package main


import (
	"fmt"
	"net/http"
)


func main(){
	http.HandleFunc("/",handle)
	http.HandleFunc("/Home",home)
	http.ListenAndServe(":6080",nil)
}
func handle(res http.ResponseWriter,req *http.Request){
fmt.Fprintf(res,"Welcome to my home page")
}
func home(res http.ResponseWriter,req *http.Request){
fmt.Fprintf(res,"Hello world")
}