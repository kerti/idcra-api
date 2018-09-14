package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func getNewValidUser() *User {
	return &User{
		ID:        "fakeUserID",
		Email:     "fakeUserEmail@gmail.com",
		Password:  "fakeUserPassword",
		IPAddress: "fakeIpAddress",
	}
}

func TestUser(t *testing.T) {

	t.Run("HashedPassword", func(t *testing.T) {

		t.Run("NoErrors", func(t *testing.T) {
			defaultUser := getNewValidUser()
			user := getNewValidUser()
			err := user.HashedPassword()

			assert.Nil(t, err)
			assert.NotEqual(t, user.Password, defaultUser.Password)
			assert.True(t, len(user.Password) > len(defaultUser.Password))
		})

		// TODO: mock bcrypt to make error handling test

	})

	t.Run("ComparePassword", func(t *testing.T) {

		t.Run("Match", func(t *testing.T) {
			defaultUser := getNewValidUser()
			user := getNewValidUser()
			err := user.HashedPassword()

			assert.Nil(t, err)

			passwordsMatch := user.ComparePassword(defaultUser.Password)

			assert.True(t, passwordsMatch)
		})

		t.Run("NoMatch", func(t *testing.T) {
			user := getNewValidUser()
			err := user.HashedPassword()

			assert.Nil(t, err)

			passwordsMatch := user.ComparePassword("someRandomString")

			assert.False(t, passwordsMatch)
		})

	})

}
