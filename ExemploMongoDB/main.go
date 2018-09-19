package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	StartServer()
}

// função que ativa o servidor na porta /users
func StartServer() {

	r := mux.NewRouter()

	r.HandleFunc("/users", RegisteUser).Methods("POST")

	r.HandleFunc("/users", SearchUser).Methods("POST")

	http.ListenAndServe(SERVER_HOST, r)

}
