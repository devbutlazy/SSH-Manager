package ssh

import (
	"bufio"
	"fmt"
	"os"
	"ssh-manager/config"
	"ssh-manager/utils"
)

func ConnectToSSH() {
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

	fmt.Print("=== Available SSH Configurations ===\n\n")
	for index, cfg := range configs {
		fmt.Printf("[ %d ] %s (%s@%s)\n", index, cfg.UniqueName, cfg.Username, cfg.IPAddress)
	}

	fmt.Print(">>> ")
	fmt.Scan(&selected)

	bufio.NewReader(os.Stdin).ReadString('\n')

	if selected > 0 && selected <= len(configs) {
		ExecuteSSH(configs[selected])
	} else {
		fmt.Println("Invalid selection.")
		utils.WaitForEnter()
		return
	}
}
