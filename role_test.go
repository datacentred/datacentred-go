package datacentred

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFullRoleLifeCycle(t *testing.T) {
	r := initRecorder("fixtures/role_lifecycle1")
	defer r.Stop()

	roles := Roles()
	rolesCount := len(roles)

	firstRole := roles[0]

	assert.Equal(t, firstRole.Id, FindRole(firstRole.Id).Id, "they should be equal")

	params := map[string]interface{}{
		"role": map[string]interface{}{
			"name": "Wyld Stallyns",
		},
	}

	newRole := CreateRole(params)

	r2 := initRecorder("fixtures/role_lifecycle2")
	defer r2.Stop()

	assert.Equal(t, newRole.Id, FindRole(newRole.Id).Id, "they should be equal")

	assert.Equal(t, "Wyld Stallyns", newRole.Name, "they should be equal")

	r3 := initRecorder("fixtures/role_lifecycle3")
	defer r3.Stop()

	assert.Equal(t, rolesCount+1, len(Roles()), "they should be equal")

	newRole.Name = "Wyld Stallyns Rock!"
	newRole.Save()

	assert.Equal(t, "Wyld Stallyns Rock!", FindRole(newRole.Id).Name, "they should be equal")

	usersCount := len(newRole.Users())
	assert.Equal(t, 0, usersCount, "they should be equal")

	firstUser := Users()[0]

	r4 := initRecorder("fixtures/role_lifecycle4")
	defer r4.Stop()

	result := newRole.AddUser(firstUser.Id)
	assert.Equal(t, true, result, "they should be equal")

	assert.Equal(t, usersCount+1, len(newRole.Users()), "they should be equal")

	result = newRole.RemoveUser(firstUser.Id)
	assert.Equal(t, true, result, "they should be equal")

	r5 := initRecorder("fixtures/role_lifecycle5")
	defer r5.Stop()

	assert.Equal(t, usersCount, len(newRole.Users()), "they should be equal")

	result = newRole.Destroy()
	assert.Equal(t, true, result, "they should be equal")
}
