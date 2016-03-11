package models

import (
	"errors"
	"log"
)

type User struct {
	Id   int
	Name string
}

func Login(username string, password string) (*User, error) {

	//hashedPassword := sha512.Sum512([]byte(password))
	//b64Pass := base64.StdEncoding.EncodeToString(hashedPassword[:])

	if username == "yann" && password == "password" {
		result := User{Id: 1, Name: "Yann"}
		return &result, nil
	} else {
		log.Println("Login Failed")
		return nil, errors.New("Nope")
	}

}
