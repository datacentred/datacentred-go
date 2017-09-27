package datacentred

import (
	"encoding/json"
	"fmt"
	"time"
)

// Role is a grouping structure for assigning account permissions to all
// users who are assigned as members of the role.
type Role struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Admin       bool      `json:"admin"`
	Permissions []string  `json:"permissions"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type rolesResponse struct {
	Roles []Role `json:"roles"`
}

type roleResponse struct {
	Role Role `json:"role"`
}

// Roles is the collection of roles belonging to the currently authenticated user's account.
func Roles() []Role {
	data, err := Request("GET", "roles", nil)
	if err != nil {
		fmt.Errorf("Request failed: %s", err)
		return nil
	}
	var res rolesResponse
	json.Unmarshal(data, &res)
	return res.Roles
}

// FindRole locates a specific role by its unique ID.
func FindRole(Id string) Role {
	data, err := Request("GET", "roles/"+Id, nil)
	if err != nil {
		fmt.Errorf("Request failed: %s", err)
	}
	var res roleResponse
	json.Unmarshal(data, &res)
	return res.Role
}

// CreateRole creates a new role on the currently authenticated user's account.
func CreateRole(Params interface{}) Role {
	role := map[string]interface{}{
		"role": Params,
	}
	jsonStr, _ := json.Marshal(role)
	data, err := Request("POST", "roles", jsonStr)
	if err != nil {
		fmt.Errorf("Request failed: %s", err)
	}
	var res roleResponse
	json.Unmarshal(data, &res)
	return res.Role
}

// Save commits any changes to this role's fields.
func (r Role) Save() Role {
	role := map[string]interface{}{
		"role": r,
	}
	jsonStr, _ := json.Marshal(role)
	data, err := Request("PUT", "roles/"+r.Id, jsonStr)
	if err != nil {
		fmt.Errorf("Request failed: %s", err)
	}
	var res roleResponse
	json.Unmarshal(data, &res)
	return res.Role
}

// Destroy permanently removes this role.
func (r Role) Destroy() bool {
	_, err := Request("DELETE", "roles/"+r.Id, nil)
	if err != nil {
		fmt.Errorf("Request failed: %s", err)
	}
	return true
}

// Users is the collection of users assigned to this role.
func (r Role) Users() []User {
	data, err := Request("GET", "roles/"+r.Id+"/users", nil)
	if err != nil {
		fmt.Errorf("Request failed: %s", err)
		return nil
	}
	var res usersResponse
	json.Unmarshal(data, &res)
	return res.Users
}

// AddUser assigns a user to this role, granting them the role's permissions.
func (r Role) AddUser(UserId string) bool {
	_, err := Request("PUT", "roles/"+r.Id+"/users/"+UserId, nil)
	if err != nil {
		fmt.Errorf("Request failed: %s", err)
	}
	return true
}

// RemoveUser removes a user from this role, revoking the role's permissions.
func (r Role) RemoveUser(UserId string) bool {
	_, err := Request("DELETE", "roles/"+r.Id+"/users/"+UserId, nil)
	if err != nil {
		fmt.Errorf("Request failed: %s", err)
	}
	return true
}
