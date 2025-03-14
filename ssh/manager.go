package ssh

import (
	"bufio"
	"fmt"
	"os"
	"ssh_manager/config"
	"ssh_manager/utils"
)

func AddConfig() {
	var newConfig config.SSHConfig
	reader := bufio.NewReader(os.Stdin)

	utils.ClearScreen()
	fmt.Println("=== Add New SSH Configuration ===")
	fmt.Println()

	fmt.Print("Enter SSH IP Address: ")
	newConfig.IPAddress, _ = reader.ReadString('\n')
	newConfig.IPAddress = config.CleanInput(newConfig.IPAddress)

	fmt.Print("Enter SSH Username: ")
	newConfig.Username, _ = reader.ReadString('\n')
	newConfig.Username = config.CleanInput(newConfig.Username)

	fmt.Print("Enter Private Key Path (optional, press Enter to skip): ")
	newConfig.PrivateKey, _ = reader.ReadString('\n')
	newConfig.PrivateKey = config.CleanInput(newConfig.PrivateKey)

	fmt.Print("Enter a Unique Name: ")
	newConfig.UniqueName, _ = reader.ReadString('\n')
	newConfig.UniqueName = config.CleanInput(newConfig.UniqueName)

	configs, err := config.LoadConfigs()
	if err != nil {
		fmt.Println("Error loading configs:", err)
		return
	}

	configs = append(configs, newConfig)
	if err := config.SaveConfigs(configs); err != nil {
		fmt.Println("Error saving config:", err)
		return
	}

	fmt.Println("\nSSH Configuration added successfully!")
	utils.WaitForEnter()
}

func RemoveConfig() {
	fmt.Println("Remove SSH functionality not implemented yet.")
	utils.WaitForEnter()
}