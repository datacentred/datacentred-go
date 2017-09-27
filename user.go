package datacentred

import (
	"encoding/json"
	"fmt"
	"time"
)

// User is a member of this account's team.
// Users can log into the DataCentred dashboard and be
// assigned to OpenStack projects and granted system
// permissions via role assignment.
type User struct {
	Id        string    `json:"id"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type usersResponse struct {
	Users []User `json:"users"`
}
type userResponse struct {
	User User `json:"user"`
}

// Users is the collection of users belonging to the currently authenticated user's account.
func Users() []User {
	data, err := Request("GET", "users", nil)
	if err != nil {
		fmt.Errorf("Request failed: %s", err)
		return nil
	}
	var res usersResponse
	json.Unmarshal(data, &res)
	return res.Users
}

// FindUser locates a specific user by its unique ID.
func FindUser(Id string) User {
	data, err := Request("GET", "users/"+Id, nil)
	if err != nil {
		fmt.Errorf("Request failed: %s", err)
	}
	var res userResponse
	json.Unmarshal(data, &res)
	return res.User
}

// CreateUser creates a new user on the currently authenticated user's account.
// A password must be specified to create a new user.
func CreateUser(Params interface{}) User {
	user := map[string]interface{}{
		"user": Params,
	}
	jsonStr, _ := json.Marshal(user)
	data, err := Request("POST", "users", jsonStr)
	if err != nil {
		fmt.Errorf("Request failed: %s", err)
	}
	var res userResponse
	json.Unmarshal(data, &res)
	return res.User
}

// Save commits any changes to this user's fields.
func (u User) Save() User {
	user := map[string]interface{}{
		"user": u,
	}
	jsonStr, _ := json.Marshal(user)
	data, err := Request("PUT", "users/"+u.Id, jsonStr)
	if err != nil {
		fmt.Errorf("Request failed: %s", err)
	}
	var res userResponse
	json.Unmarshal(data, &res)
	return res.User
}

// Destroy permanently removes this user.
func (u User) Destroy() bool {
	_, err := Request("DELETE", "users/"+u.Id, nil)
	if err != nil {
		fmt.Errorf("Request failed: %s", err)
	}
	return true
}

// ChangePassword allows a new password to be set for this user.
func (u User) ChangePassword(Password string) bool {
	user := map[string]interface{}{
		"password": Password,
	}
	jsonStr, _ := json.Marshal(user)
	_, err := Request("PUT", "users/"+u.Id, jsonStr)
	if err != nil {
		fmt.Errorf("Request failed: %s", err)
	}
	return true
}
