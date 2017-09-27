package datacentred

import (
	"encoding/json"
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
func Roles() ([]Role, error) {
	data, err := Request("GET", "roles", nil)
	if err != nil {
		return nil, err
	}
	var res rolesResponse
	json.Unmarshal(data, &res)
	return res.Roles, nil
}

// FindRole locates a specific role by its unique ID.
func FindRole(Id string) (*Role, error) {
	data, err := Request("GET", "roles/"+Id, nil)
	if err != nil {
		return nil, err
	}
	var res roleResponse
	json.Unmarshal(data, &res)
	return &res.Role, nil
}

// CreateRole creates a new role on the currently authenticated user's account.
func CreateRole(Params interface{}) (*Role, error) {
	role := map[string]interface{}{
		"role": Params,
	}
	jsonStr, _ := json.Marshal(role)
	data, err := Request("POST", "roles", jsonStr)
	if err != nil {
		return nil, err
	}
	var res roleResponse
	json.Unmarshal(data, &res)
	return &res.Role, nil
}

// Save commits any changes to this role's fields.
func (r Role) Save() (Role, error) {
	role := map[string]interface{}{
		"role": r,
	}
	jsonStr, _ := json.Marshal(role)
	data, err := Request("PUT", "roles/"+r.Id, jsonStr)
	if err != nil {
		return r, err
	}
	var res roleResponse
	json.Unmarshal(data, &res)
	return res.Role, nil
}

// Destroy permanently removes this role.
func (r Role) Destroy() (bool, error) {
	_, err := Request("DELETE", "roles/"+r.Id, nil)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Users is the collection of users assigned to this role.
func (r Role) Users() ([]User, error) {
	data, err := Request("GET", "roles/"+r.Id+"/users", nil)
	if err != nil {
		return nil, err
	}
	var res usersResponse
	json.Unmarshal(data, &res)
	return res.Users, nil
}

// AddUser assigns a user to this role, granting them the role's permissions.
func (r Role) AddUser(UserId string) (bool, error) {
	_, err := Request("PUT", "roles/"+r.Id+"/users/"+UserId, nil)
	if err != nil {
		return false, err
	}
	return true, nil
}

// RemoveUser removes a user from this role, revoking the role's permissions.
func (r Role) RemoveUser(UserId string) (bool, error) {
	_, err := Request("DELETE", "roles/"+r.Id+"/users/"+UserId, nil)
	if err != nil {
		return false, err
	}
	return true, nil
}
