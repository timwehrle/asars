package auth

import (
	"fmt"

	"github.com/zalando/go-keyring"
)

var (
	service  = "asars"
	username = "user"
)

// Set token for Asana
func SetToken(token string) error {
	err := keyring.Set(service, username, token)
	if err != nil {
		return fmt.Errorf("failed to set token: %v", err)
	}
	fmt.Println("Token set successfully.")

	return nil
}

// Get token for Asana
func GetToken() (string, error) {
	token, err := keyring.Get(service, username)
	if err != nil {
		return "", fmt.Errorf("failed to get token: %v", err)
	}

	return token, nil
}

// Delete token for Asana
func DeleteToken() error {
	err := keyring.Delete(service, username)
	if err != nil {
		return fmt.Errorf("failed to delete token: %v", err)
	}
	fmt.Println("Token deleted successfully.")

	return nil
}

// Check if token exists
func HasToken() bool {
	token, err := GetToken()
	if err != nil {
		fmt.Println("Error checking token existence:", err)
		return false
	}
	
	return token != ""
}

func UpdateToken(token string) error {
	if !HasToken() {
		return fmt.Errorf("no token found")
	}

	err := keyring.Set(service, username, token)
	if err != nil {
		return fmt.Errorf("failed to update token: %v", err)
	}
	fmt.Println("Token updated successfully.")

	return nil
}
