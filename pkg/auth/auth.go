package auth

import (
	"fmt"

	"github.com/zalando/go-keyring"
)

var (
	service  = "asars"
	username = "user"
)

func SetToken(token string) error {
	// Set token for Asana
	err := keyring.Set(service, username, token)
	if err != nil {
		return fmt.Errorf("failed to set token: %v", err)
	}
	fmt.Println("Token set successfully.")

	return nil
}

func GetToken() (string, error) {
	// Get token for Asana
	token, err := keyring.Get(service, username)
	if err != nil {
		return "", fmt.Errorf("failed to get token: %v", err)
	}

	return token, nil
}

func DeleteToken() error {
	// Delete token for Asana
	err := keyring.Delete(service, username)
	if err != nil {
		return fmt.Errorf("failed to delete token: %v", err)
	}
	fmt.Println("Token deleted successfully.")

	return nil
}

func HasToken() bool {
	// Check if token exists
	token, _ := GetToken()
	return token != ""
}
