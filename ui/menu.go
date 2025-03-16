package ui

import (
	"fmt"
	"os"
	"ssh-manager/config"
	"ssh-manager/ssh"
	"ssh-manager/utils"
)

func MainMenu() {
	utils.ClearScreen()

	var choice int

	fmt.Println("=== SSH Configuration Manager ===")
	fmt.Println("[ 1 ] Connect to SSH")
	fmt.Println("[ 2 ] Add SSH connection")
	fmt.Println("[ 3 ] Remove SSH connection")
	fmt.Println("[ 4 ] Exit")
	fmt.Print(">>> ")

	_, err := fmt.Scan(&choice)
	if err != nil {
		fmt.Println("Invalid input. Please enter a number.")
		utils.WaitForEnter()
		return
	}

	switch choice {
	case 1:
		ssh.ConnectToSSH()
	case 2:
		addSSH()
	case 3:
		ssh.RemoveSSH()
	case 4:
		utils.ClearScreen()
		fmt.Println("=== Goodbye ===")
		os.Exit(0)
	default:
		fmt.Println("Invalid choice")
		utils.WaitForEnter()
	}
}

func addSSH() {
	utils.ClearScreen()
	var newConfig config.SSHConfig
	fmt.Println("=== Add New SSH Configuration ===")

	newConfig.IPAddress = utils.ReadInput("Enter IP address: ")
	newConfig.Username = utils.ReadInput("Enter username: ")
	newConfig.PrivateKey = utils.ReadInput("Enter private key: ")
	newConfig.UniqueName = utils.ReadInput("Enter unique name: ")

	configs, err := config.LoadConfigs()
	if err != nil {
		fmt.Println("Error loading configs:", err)
		utils.WaitForEnter()
	}

	configs = append(configs, newConfig)
	if err := config.SaveConfigs(configs); err != nil {
		fmt.Println("Error saving config:", err)
		utils.WaitForEnter()
	}
}
