package datacentred

import (
	"encoding/json"
	"fmt"
	"time"
)

type Role struct {
	Id          string
	Name        string
	Admin       bool
	Permissions []string
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
	data, err := Request("GET", "roles")
	if err != nil {
		fmt.Errorf("Request failed: %s", err)
		return nil
	}
	var res RolesResponse
	json.Unmarshal(data, &res)
	return res.Roles
}

func FindRole(Id string) Role {
	data, err := Request("GET", "roles/"+Id)
	if err != nil {
		fmt.Errorf("Request failed: %s", err)
	}
	var res RoleResponse
	json.Unmarshal(data, &res)
	return res.Role
}
