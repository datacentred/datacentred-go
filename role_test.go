package datacentred

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFullRoleLifeCycle(t *testing.T) {
	r := initRecorder("fixtures/role_lifecycle1")
	defer r.Stop()

	roles, _ := Roles()
	rolesCount := len(roles)

	firstRole := roles[0]

	role, _ := FindRole(firstRole.Id)
	assert.Equal(t, firstRole.Id, role.Id, "they should be equal")

	params := map[string]string{
		"name": "Wyld Stallyns",
	}

	newRole, _ := CreateRole(params)

	r2 := initRecorder("fixtures/role_lifecycle2")
	defer r2.Stop()

	role, _ = FindRole(newRole.Id)
	assert.Equal(t, newRole.Id, role.Id, "they should be equal")

	assert.Equal(t, "Wyld Stallyns", newRole.Name, "they should be equal")

	r3 := initRecorder("fixtures/role_lifecycle3")
	defer r3.Stop()

	roles, _ = Roles()
	assert.Equal(t, rolesCount+1, len(roles), "they should be equal")

	newRole.Name = "Wyld Stallyns Rock!"
	newRole.Save()

	role, _ = FindRole(newRole.Id)
	assert.Equal(t, "Wyld Stallyns Rock!", role.Name, "they should be equal")

	users, _ := newRole.Users()
	usersCount := len(users)
	assert.Equal(t, 0, usersCount, "they should be equal")

	users, _ = Users()
	firstUser := users[0]

	r4 := initRecorder("fixtures/role_lifecycle4")
	defer r4.Stop()

	result, _ := newRole.AddUser(firstUser.Id)
	assert.Equal(t, true, result, "they should be equal")

	users, _ = newRole.Users()
	assert.Equal(t, usersCount+1, len(users), "they should be equal")

	result, _ = newRole.RemoveUser(firstUser.Id)
	assert.Equal(t, true, result, "they should be equal")

	r5 := initRecorder("fixtures/role_lifecycle5")
	defer r5.Stop()

	users, _ = newRole.Users()
	assert.Equal(t, usersCount, len(users), "they should be equal")

	result, _ = newRole.Destroy()
	assert.Equal(t, true, result, "they should be equal")
}

func TestRoleErrors(t *testing.T) {
	r1 := initRecorder("fixtures/role_errors1")
	defer r1.Stop()

	_, err := FindRole("bogus")
	assert.Equal(t, "404 Not Found.", err.Error(), "they should be equal")

	params := map[string]string{
		"name": "",
	}

	newRole, err := CreateRole(params)
	assert.Nil(t, newRole)
	assert.Equal(t, "422 Unprocessable Entity. Role name too short.", err.Error(), "they should be equal")

	r2 := initRecorder("fixtures/role_errors2")
	defer r2.Stop()

	params = map[string]string{
		"name": "SanDimasHigh",
	}

	newRole, _ = CreateRole(params)

	_, err = newRole.AddUser("Boom!")
	assert.Equal(t, "404 Not Found. No such user 'Boom!'", err.Error(), "they should be equal")

	_, err = newRole.RemoveUser("Boom!")
	assert.Equal(t, "404 Not Found. No such user 'Boom!'", err.Error(), "they should be equal")

	newRole.Destroy()
	newRole.Name = "Boom!"
	_, err = newRole.Save()
	assert.Equal(t, "404 Not Found.", err.Error(), "they should be equal")

	r3 := initRecorder("fixtures/role_errors3")
	defer r3.Stop()

	_, err = newRole.Destroy()
	assert.Equal(t, "404 Not Found.", err.Error(), "they should be equal")

	_, err = newRole.Users()
	assert.Equal(t, "404 Not Found. No such role 'bcb904ee-afb7-4f38-ac55-a9e29804dbf1'", err.Error(), "they should be equal")
}
