package asars

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

type WorkspaceConfig struct {
	DefaultWorkspace string `json:"default_workspace_gid"`
}

// getConfigFilePath returns the path to the configuration file.
// The configuration file is located in the user's configuration directory.
func getConfigFilePath() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	configPath := filepath.Join(configDir, "asars")
	if err := os.MkdirAll(configPath, 0755); err != nil {
		return "", err
	}

	return filepath.Join(configPath, "asars_default_workspace.json"), nil
}

// SaveDefaultWorkspace saves the default workspace GID to the configuration file.
func SaveDefaultWorkspace(workspaceGID string) error {
	path, err := getConfigFilePath()
	if err != nil {
		return err
	}

	config := WorkspaceConfig{DefaultWorkspace: workspaceGID}
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(config)
}

// LoadDefaultWorkspace loads the default workspace GID from the configuration file.
func LoadDefaultWorkspace() (string, error) {
	path, err := getConfigFilePath()
	if err != nil {
		return "", err
	}

	file, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return "", errors.New("no default workspace set")
		}
		return "", err
	}
	defer file.Close()

	var config WorkspaceConfig
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return "", err
	}

	return config.DefaultWorkspace, nil
}
