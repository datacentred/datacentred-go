package datacentred

import (
	"encoding/json"
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
func Users() ([]User, error) {
	data, err := Request("GET", "users", nil)
	if err != nil {
		return nil, err
	}
	var res usersResponse
	json.Unmarshal(data, &res)
	return res.Users, nil
}

// FindUser locates a specific user by its unique ID.
func FindUser(Id string) (*User, error) {
	data, err := Request("GET", "users/"+Id, nil)
	if err != nil {
		return nil, err
	}
	var res userResponse
	json.Unmarshal(data, &res)
	return &res.User, nil
}

// CreateUser creates a new user on the currently authenticated user's account.
// A password must be specified to create a new user.
func CreateUser(Params interface{}) (*User, error) {
	user := map[string]interface{}{
		"user": Params,
	}
	jsonStr, _ := json.Marshal(user)
	data, err := Request("POST", "users", jsonStr)
	if err != nil {
		return nil, err
	}
	var res userResponse
	json.Unmarshal(data, &res)
	return &res.User, nil
}

// Save commits any changes to this user's fields.
func (u User) Save() (User, error) {
	user := map[string]interface{}{
		"user": u,
	}
	jsonStr, _ := json.Marshal(user)
	data, err := Request("PUT", "users/"+u.Id, jsonStr)
	if err != nil {
		return u, err
	}
	var res userResponse
	json.Unmarshal(data, &res)
	return res.User, nil
}

// Destroy permanently removes this user.
func (u User) Destroy() (bool, error) {
	_, err := Request("DELETE", "users/"+u.Id, nil)
	if err != nil {
		return false, err
	}
	return true, nil
}

// ChangePassword allows a new password to be set for this user.
func (u User) ChangePassword(Password string) (bool, error) {
	user := map[string]interface{}{
		"user": map[string]string{
			"password": Password,
		},
	}
	jsonStr, _ := json.Marshal(user)
	_, err := Request("PUT", "users/"+u.Id, jsonStr)
	if err != nil {
		return false, err
	}
	return true, nil
}
