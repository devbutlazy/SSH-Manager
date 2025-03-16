package config

import (
	"encoding/json"
	"os"
)

const ConfigFile = "ssh_config.json"

type SSHConfig struct {
	IPAddress  string `json:"ip_address"`
	Username   string `json:"username"`
	PrivateKey string `json:"private_key"`
	UniqueName string `json:"name"`
}

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
