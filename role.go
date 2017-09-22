package datacentred

import (
	"encoding/json"
	"fmt"
)

type Role struct {
	Id          string
	Name        string
	Admin       bool
	Permissions []string
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type RoleResponse struct {
	Roles []Role
}

func ListRoles() []Role {
	data, err := Request("GET", "roles")
	if err != nil {
		fmt.Errorf("Request failed: %s", err)
		return nil
	}
	var res RoleResponse
	json.Unmarshal(data, &res)
	return res.Roles
}
