package main

import ("fmt"
		"log"
		"github.com/gorilla/mux"
		"gopkg.in/mgo.v2"
		"gopkg.in/mgo.v2/bson"
		"net/http")


type User struct {
	Username string 
	Email string
	Password string
}

func main() {

	session, err := mgo.Dial("127.0.0.1")
        if err != nil {
                panic(err)
        }
	defer session.Close()

	c := session.DB("usersDB").C("users")

	test := User{
		Username: "Theduardds",
		Email: "Theduardds@gmail.com",
		Password: "123456",
	}

	err = c.Insert(test)

    if err != nil {
    	log.Fatal(err)
    }


	result := User{}
    
    err = c.Find(bson.M{"username": "Theduardds"}).One(&result)
    if err != nil {
    	log.Fatal(err)
    }

    fmt.Println("User:", result.Username)


	r := mux.NewRouter()
	r.HandleFunc("/user", AddNewUser).Methods("POST")

	if err := http.ListenAndServe("172.22.51.161:8080", r); err != nil {
    	panic(err)
	}	
}

func AddNewUser(w http.ResponseWriter, r *http.Request) {
	
	return
}

// buscar pela biblioteca MGO 
-> https://labix.org/mgo
-> https://godoc.org/gopkg.in/mgo.v2
-> https://docs.mongodb.com/manual/

// função de conexão(garantir que só tem uma)
// inserir um json e buscar json
