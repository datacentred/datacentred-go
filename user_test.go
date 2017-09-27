package datacentred

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFullUserLifeCycle(t *testing.T) {
	r1 := initRecorder("fixtures/user_lifecycle1")
	defer r1.Stop()

	users, _ := Users()
	usersCount := len(users)

	firstUser := users[0]

  user, _ := FindUser(firstUser.Id)
	assert.Equal(t, firstUser.Id, user.Id, "they should be equal")

	params := map[string]string{
		"email":      "bill.s.preston@esquire.com",
		"password":   "Non Heinous",
		"first_name": "Bill S.",
		"last_name":  "Preston",
	}

	newUser, _ := CreateUser(params)

	r2 := initRecorder("fixtures/user_lifecycle2")
	defer r2.Stop()

  user, _ = FindUser(newUser.Id)
	assert.Equal(t, newUser.Id, user.Id, "they should be equal")

	assert.Equal(t, "Preston", newUser.LastName, "they should be equal")

	r3 := initRecorder("fixtures/user_lifecycle3")
	defer r3.Stop()

  users, _ = Users()
	assert.Equal(t, usersCount+1, len(users), "they should be equal")

	newUser.LastName = "Preston Esq."
	newUser.Save()

  user, _ = FindUser(newUser.Id)
	assert.Equal(t, "Preston Esq.", user.LastName, "they should be equal")

  result, _ := newUser.ChangePassword("Excellent")
	assert.Equal(t, true, result, "they should be equal")

	result, _ = newUser.Destroy()
	assert.Equal(t, true, result, "they should be equal")
}
