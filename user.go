package datacentred

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Id        string
	Email     string
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Password  string
}

type UserResponse struct {
	Users []User
}

func ListUsers() []User {
	data, err := Request("GET", "users")
	if err != nil {
		fmt.Errorf("Request failed: %s", err)
		return nil
	}
	var res UserResponse
	json.Unmarshal(data, &res)
	return res.Users
}
