package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter,r *http.Request){
	if r.URL.Path !="/form"{
		http.Error(w,"Wrong URL path",http.StatusNotFound)
		return
	}
	if r.Method !="POST"{
		http.Error(w,"Not a POST requset",http.StatusNotFound)
		return 
	}

	if err:=r.ParseForm();err!=nil{
        fmt.Fprintf(w,"Can't parse form data %v",err)
		return 
	}

	fmt.Fprint(w,"POST request recieved")

	name:=r.FormValue("name")
	age:=r.FormValue("age")
	address := r.FormValue("address")


	fmt.Fprintf(w,"Name is %v",name)
	fmt.Fprintf(w,"Age is %v",age)
	fmt.Fprintf(w,"Address is %v",address)
}


func helloHandler(w http.ResponseWriter,r *http.Request){
	if r.URL.Path !="/hello"{
		http.Error(w,"Wrong URL path",http.StatusNotFound)
		return
	}
	if r.Method !="GET"{
		http.Error(w,"Not a GET requset",http.StatusNotFound)
		return 
	}
	fmt.Fprintf(w,"Hello!!")
}



func main(){
	server := http.FileServer(http.Dir("./static"))
	http.Handle("/",server)
	http.HandleFunc("/form",formHandler)
	http.HandleFunc("/hello",helloHandler)


	fmt.Printf("Server is listening on PORT 8080\n")

	if err:= http.ListenAndServe(":8080",nil);err!=nil{
		log.Fatal(err)
	}

}
