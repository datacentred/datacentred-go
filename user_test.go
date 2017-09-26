package datacentred

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFullUserLifeCycle(t *testing.T) {
	r1 := initRecorder("fixtures/user_lifecycle1")
	defer r1.Stop()

	users := Users()
	usersCount := len(users)

	firstUser := users[0]

	assert.Equal(t, firstUser.Id, FindUser(firstUser.Id).Id, "they should be equal")

	params := map[string]interface{}{
		"user": map[string]interface{}{
			"email":      "bill.s.preston@esquire.com",
			"password":   "Non Heinous",
			"first_name": "Bill S.",
			"last_name":  "Preston",
		},
	}

	newUser := CreateUser(params)

	r2 := initRecorder("fixtures/user_lifecycle2")
	defer r2.Stop()

	assert.Equal(t, newUser.Id, FindUser(newUser.Id).Id, "they should be equal")

	assert.Equal(t, "Preston", newUser.LastName, "they should be equal")

	r3 := initRecorder("fixtures/user_lifecycle3")
	defer r3.Stop()

	assert.Equal(t, usersCount+1, len(Users()), "they should be equal")

	newUser.LastName = "Preston Esq."
	newUser.Save()

	assert.Equal(t, "Preston Esq.", FindUser(newUser.Id).LastName, "they should be equal")

	result := newUser.Destroy()
	assert.Equal(t, true, result, "they should be equal")
}
