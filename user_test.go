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

func TestUserErrors(t *testing.T) {
	r1 := initRecorder("fixtures/user_errors1")
	defer r1.Stop()

	_, err := FindUser("bogus")
	assert.Equal(t, "Not found", err.Error(), "they should be equal")

	params := map[string]string{
		"email": "",
	}

	newUser, err := CreateUser(params)
	assert.Nil(t, newUser)
	assert.Equal(t, "User email can't be blank.", err.Error(), "they should be equal")

	r2 := initRecorder("fixtures/user_errors2")
	defer r2.Stop()

	params = map[string]string{
		"email": "bill.s.preston@esquire.com",
		"password": "Excellent",
	}

	newUser, _ = CreateUser(params)

  newUser.Destroy()
  newUser.FirstName = "Boom!"
  _, err = newUser.Save()
  assert.Equal(t, "Not found", err.Error(), "they should be equal")

  r3 := initRecorder("fixtures/user_errors3")
	defer r3.Stop()

  _, err = newUser.Destroy()
  assert.Equal(t, "Not found", err.Error(), "they should be equal")  
}
