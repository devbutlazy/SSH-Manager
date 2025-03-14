package ssh

import (
	"bufio"
	"fmt"
	"os"
	"ssh_manager/config"
	"ssh_manager/utils"
	"strings"

	"golang.org/x/term"
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

func selectConfigForRemoval(configs []config.SSHConfig) int {
	selected := 0
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println("Error setting terminal mode:", err)
		return -1
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	for {
		utils.ClearScreen()
		fmt.Println("=== Available SSH Configurations ===")
		fmt.Println()
		for i, config := range configs {
			if i == selected {
				fmt.Printf("> [%d] %s (%s@%s)\n", i+1, config.UniqueName, config.Username, config.IPAddress)
			} else {
				fmt.Printf("  [%d] %s (%s@%s)\n", i+1, config.UniqueName, config.Username, config.IPAddress)
			}
		}
		fmt.Println("\nUse ↑↓ arrow keys to navigate, Enter to select, Esc (twice) to go back")

		b := make([]byte, 1)
		_, err := os.Stdin.Read(b)
		if err != nil {
			fmt.Println("Error reading input:", err)
			return -1
		}

		if b[0] == 27 {
			b = make([]byte, 2)
			_, err := os.Stdin.Read(b)
			if err != nil {
				continue
			}
			if b[0] == 91 { // CSI (Control Sequence Introducer)
				switch b[1] {
				case 65:
					if selected > 0 {
						selected--
					}

				case 66:
					if selected < len(configs)-1 {
						selected++
					}
				}
			} else if b[0] == 0 || b[1] == 0 {
				return -1
			}
		} else if b[0] == 13 {
			return selected
		}
	}
}

func RemoveConfig() {
	configs, err := config.LoadConfigs()
	if err != nil {
		fmt.Println("Error loading configs:", err)
		utils.WaitForEnter()
		return
	}

	if len(configs) == 0 {
		fmt.Println("No configurations available.")
		utils.WaitForEnter()
		return
	}

	// Display available SSH configurations
	selected := selectConfigForRemoval(configs)
	if selected == -1 {
		return
	}

	// Confirm removal
	selectedConfig := configs[selected]
	fmt.Printf("Are you sure you want to remove the configuration: %s (%s@%s)? (y/n): ", selectedConfig.UniqueName, selectedConfig.Username, selectedConfig.IPAddress)
	var confirmation string
	fmt.Scanln(&confirmation)

	if strings.ToLower(confirmation) == "y" {
		// Remove the selected configuration from the list
		configs = append(configs[:selected], configs[selected+1:]...)

		// Save the updated configurations
		if err := config.SaveConfigs(configs); err != nil {
			fmt.Println("Error saving config:", err)
			utils.WaitForEnter()
			return
		}

		fmt.Println("\nSSH Configuration removed successfully!")
	} else {
		fmt.Println("\nOperation canceled.")
	}

	utils.WaitForEnter()
}
