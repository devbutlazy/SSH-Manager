package config

import (
	"encoding/json"
	"os"
	"strings"
)

type SSHConfig struct {
	IPAddress  string `json:"ip_address"`
	Username   string `json:"username"`
	PrivateKey string `json:"private_key,omitempty"`
	UniqueName string `json:"unique_name"`
}

const ConfigFile = "ssh_config.json"

func LoadConfigs() ([]SSHConfig, error) {
	var configs []SSHConfig
	if _, err := os.Stat(ConfigFile); os.IsNotExist(err) {
		return configs, nil
	}

	file, err := os.ReadFile(ConfigFile)
	if err != nil {
		return nil, err
	}
	
	err = json.Unmarshal(file, &configs)
	return configs, err
}

func SaveConfigs(configs []SSHConfig) error {
	fileData, err := json.MarshalIndent(configs, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(ConfigFile, fileData, 0644)
}

func CleanInput(input string) string {
	return strings.TrimSpace(input)
}