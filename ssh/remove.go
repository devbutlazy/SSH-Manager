package ssh

import (
	"fmt"
	"ssh-manager/config"
	"ssh-manager/utils"
)

func RemoveSSH() {
	configs, err := config.LoadConfigs()
	if err != nil {
		fmt.Println("Error reading config file:", err)
		utils.WaitForEnter()
		return
	}

	if len(configs) == 0 {
		fmt.Println("No configurations available.")
		utils.WaitForEnter()
		return
	}

	utils.ClearScreen()

	var selected int

	fmt.Print("=== Remove SSH Configurations ===\n\n")
	for index, cfg := range configs {
		fmt.Printf("[ %d ] %s (%s@%s)\n", index, cfg.UniqueName, cfg.Username, cfg.IPAddress)
	}

	fmt.Print(">>> ")
	fmt.Scan(&selected)

	configs = append(configs[:selected], configs[selected:]...)

	if err := config.SaveConfigs(configs); err != nil {
		fmt.Println("Error saving config:", err)
		utils.WaitForEnter()
		return
	}

	fmt.Println("\nSSH Configuration removed successfully!")
	utils.WaitForEnter()
}
