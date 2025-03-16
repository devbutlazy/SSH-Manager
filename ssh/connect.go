package ssh

import (
	"fmt"
	"ssh-manager/config"
	"ssh-manager/utils"
)

func ConnectToSSH() {
	configs, err := config.LoadConfigs()
	if err != nil {
		fmt.Println("Error reading config file:", err)
		utils.WaitForEnter()
	}

	if len(configs) == 0 {
		fmt.Println("No configurations available.")
		utils.WaitForEnter()
	}

	utils.ClearScreen()

	var selected int

	fmt.Print("=== Available SSH Configurations ===\n\n")
	for index, cfg := range configs {
		fmt.Printf("[ %d ] %s (%s@%s)\n", index+1, cfg.UniqueName, cfg.Username, cfg.IPAddress)
	}

	fmt.Print(">>> ")
	fmt.Scan(&selected)

	if selected > 0 && selected <= len(configs) {
		ExecuteSSH(configs[selected-1])
	} else {
		fmt.Println("Invalid selection.")
		utils.WaitForEnter()
	}
}
