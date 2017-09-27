package datacentred

import (
	"encoding/json"
	"fmt"
	"time"
)

type User struct {
	Id        string    `json:"id"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Password  string    `json:"password"`
}

type UsersResponse struct {
	Users []User `json:"users"`
}
type UserResponse struct {
	User User `json:"user"`
}

func Users() []User {
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

func CreateUser(Params interface{}) User {
	user := map[string]interface{}{
		"user": Params,
	}
	jsonStr, _ := json.Marshal(user)
	data, err := Request("POST", "users", jsonStr)
	if err != nil {
		fmt.Errorf("Request failed: %s", err)
	}
	var res UserResponse
	json.Unmarshal(data, &res)
	return res.User
}

func (u User) Save() User {
	user := map[string]interface{}{
		"user": u,
	}
	jsonStr, _ := json.Marshal(user)
	data, err := Request("PUT", "users/"+u.Id, jsonStr)
	if err != nil {
		fmt.Errorf("Request failed: %s", err)
	}
	var res UserResponse
	json.Unmarshal(data, &res)
	return res.User
}

func (u User) Destroy() bool {
	_, err := Request("DELETE", "users/"+u.Id, nil)
	if err != nil {
		fmt.Errorf("Request failed: %s", err)
	}
	return true
}
