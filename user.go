package datacentred

import (
	"encoding/json"
	"fmt"
	"time"
)

type User struct {
	Id        string
	Email     string
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Password  string
}

type UsersResponse struct {
	Users []User
}
type UserResponse struct {
	User User
}

func ListUsers() []User {
	data, err := Request("GET", "users", nil)
	if err != nil {
		fmt.Errorf("Request failed: %s", err)
		return nil
	}
	var res UsersResponse
	json.Unmarshal(data, &res)
	return res.Users
}

func FindUser(Id string) User {
	data, err := Request("GET", "users/"+Id, nil)
	if err != nil {
		fmt.Errorf("Request failed: %s", err)
	}
	var res UserResponse
	json.Unmarshal(data, &res)
	return res.User
}
