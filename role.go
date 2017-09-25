package datacentred

import (
	"encoding/json"
	"fmt"
	"time"
)

type Role struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Admin       bool      `json:"admin"`
	Permissions []string  `json:"permissions"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type RolesResponse struct {
	Roles []Role
}

type RoleResponse struct {
	Role Role
}

func ListRoles() []Role {
	data, err := Request("GET", "roles", nil)
	if err != nil {
		fmt.Errorf("Request failed: %s", err)
		return nil
	}
	var res RolesResponse
	json.Unmarshal(data, &res)
	return res.Roles
}

func FindRole(Id string) Role {
	data, err := Request("GET", "roles/"+Id, nil)
	if err != nil {
		fmt.Errorf("Request failed: %s", err)
	}
	var res RoleResponse
	json.Unmarshal(data, &res)
	return res.Role
}

func CreateRole(Params interface{}) Role {
	jsonStr, _ := json.Marshal(Params)
	data, err := Request("POST", "roles", jsonStr)
	if err != nil {
		fmt.Errorf("Request failed: %s", err)
	}
	var res RoleResponse
	json.Unmarshal(data, &res)
	return res.Role
}

func UpdateRole(Id string, Params interface{}) Role {
	jsonStr, _ := json.Marshal(Params)
	data, err := Request("PUT", "roles/"+Id, jsonStr)
	if err != nil {
		fmt.Errorf("Request failed: %s", err)
	}
	var res RoleResponse
	json.Unmarshal(data, &res)
	return res.Role
}

func DestroyRole(Id string) bool {
	_, err := Request("DELETE", "roles/"+Id, nil)
	if err != nil {
		fmt.Errorf("Request failed: %s", err)
	}
	return true
}

func (r Role) Users() []User {
	data, err := Request("GET", "roles/"+r.Id+"/users", nil)
	if err != nil {
		fmt.Errorf("Request failed: %s", err)
		return nil
	}
	var res UsersResponse
	json.Unmarshal(data, &res)
	return res.Users
}

func (r Role) AddUser(UserId string) bool {
	_, err := Request("PUT", "roles/"+r.Id+"/users/"+UserId, nil)
	if err != nil {
		fmt.Errorf("Request failed: %s", err)
	}
	return true
}

func (r Role) RemoveUser(UserId string) bool {
	_, err := Request("DELETE", "roles/"+r.Id+"/users/"+UserId, nil)
	if err != nil {
		fmt.Errorf("Request failed: %s", err)
	}
	return true
}
