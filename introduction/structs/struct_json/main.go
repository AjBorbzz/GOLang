package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func NewUserFromJSon(data string) (User, error) {
	var user User
	err := json.Unmarshal([]byte(data), &user)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func main() {
	jsonData := `{"name":"Alice", "email":"alice@example.com", "age":30}`

	user, err := NewUserFromJSon(jsonData)
	if err != nil {
		log.Fatalf("Error creating user: %v", err)
	}
	fmt.Println("User: %+v\n", user)
}
