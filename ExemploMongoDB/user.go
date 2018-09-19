package main

import "net/http"

// Struct com o modelo de usuario que vai ser recebido do servidor
type User struct {
	Name     string `json:name`
	Pass     string `json:pass`
	Birthday string `json:birthday`
	Gender   string `json:gender`
}

// Função  que converte o corpo do link em um json para armazenar em um usuario
func RegisteUser(response http.ResponseWriter, request *http.Request) {

	body := request.Body
	user := User{}

	DecodeJson(body, &user)
	err := SaveUser(user)

	if err != nil {
		response.Write([]byte("400"))
		return
	}
	response.Write([]byte("200"))
}

// Função salva valores no banco de dados mongoDB
func SaveUser(user User) (err error) {

	conn, err := GetConnection()
	collection := conn.DB("reprova").C("users")

	collection.Insert(user)

	return
}

func SearchUser(response http.ResponseWriter, request *http.Request) {

}
