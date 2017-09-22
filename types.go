package datacentred

type User struct {
	Id        string
	Email     string
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Password  string
}

type Project struct {
	Id       string
	Name     string
	QuotaSet struct {
		Compute struct {
			Cores    int
			Instance int
			Ram      int
		}
		Volume struct {
			Gigabytes int
			Snapshots int
			Volumes   int
		}
		Network struct {
			FloatingIp        int `json:"floating_ip"`
			Network           int
			Port              int
			Router            int
			SecurityGroup     int `json:"security_group"`
			SecurityGroupRule int `json:"security_group_rule"`
			Subnet            int
		}
	} `json:"quota_set"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type Role struct {
	Id          string
	Name        string
	Admin       bool
	Permissions []string
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type UserResponse struct {
	Users []User
}

type ProjectResponse struct {
	Projects []Project
}

type RoleResponse struct {
	Roles []Role
}
