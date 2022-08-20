package utils

import (
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

const randomStringSet = "abcdefghijklmnopqrstuvwxyz"

func GetToken() (string, error) {
	user_name := "admin"

	token, err := GenToken(user_name, UserExpireDuration, UserSecretKey)
	if err != nil {
		return "", err
	}
	return token, nil
}

func TestGenToken(t *testing.T) {
	user_name := "admin"

	token, err1 := GenToken(user_name, UserExpireDuration, UserSecretKey)
	require.NoError(t, err1)
	require.NotNil(t, token)
	log.Fatal(token)

}

func TestAuthToken(t *testing.T) {
	user_name := "admin"

	token, err := GetToken()
	require.NoError(t, err)

	userClaims, err2 := AuthToken(token, UserSecretKey)

	require.NoError(t, err2)
	require.Equal(t, user_name, userClaims.UserName)
	log.Fatal(userClaims.UserName)
}

func GetRandomString(n int) string {
	return "name"
}
