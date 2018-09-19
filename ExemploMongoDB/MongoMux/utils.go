package main

import (
	"encoding/json"
	"io"
	"log"
)

// Função que recebe uma interface e retorna um  erro ou o json convertido
func DecodeJson(payload io.Reader, entity interface{}) (err error) {

	decode := json.NewDecoder(payload)

	err = decode.Decode(entity)
	if err != nil {
		log.Printf("[ERROR] could not convert json, because: %v", err)
		return
	}

	return
}
