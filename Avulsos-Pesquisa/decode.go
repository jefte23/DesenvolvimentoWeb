package main

import ("fmt"
		"time"
		"io/ioutil"
		"net/http"
		"encoding/json"
		"github.com/gorilla/mux")

func main() {

	fmt.Println("Starting server...")
	
	r := mux.NewRouter()

	r.HandleFunc("/", Home)

	r.HandleFunc("/time", Time).Methods("GET")

	r.HandleFunc("/users", CreateUser).Methods("POST")

	http.ListenAndServe("localhost:8080", r)
}


type User struct {
	Name  string   "json:name"
	Age   int      "json:age"
	Phone []string "json:phones"
	Valid bool     "json:valid"
}

func DecodeJson(body io.Reader, entity interface{}){
	
	dec := json.NewDecoder(body)
	for{
		if err := dec.Decode(&entidy); err == io.EOF{
			break			
		}else if err != nil{
			log.Fatal(err)
		}
		fmt.Print("Request: %v", entity)
	}
}


func CreateUser(response http.ResponseWriter, request *http.Request){

	user := User{}

	DecodeJson(request.Body, user)
}
