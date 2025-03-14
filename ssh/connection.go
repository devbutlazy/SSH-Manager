package ssh

import (
	"fmt"
	"os"
	"os/exec"
	"ssh_manager/config"
	"ssh_manager/utils"

	"golang.org/x/term"
)

func Connect() {
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

	selected := selectConfig(configs)
	if selected == -1 {
		return
	}

	selectedConfig := configs[selected]
	sshCmd := buildSSHCommand(selectedConfig)

	utils.ClearScreen()
	fmt.Printf("Connecting to %s (%s@%s)...\n", selectedConfig.UniqueName, selectedConfig.Username, selectedConfig.IPAddress)
	cmd := exec.Command(sshCmd[0], sshCmd[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()

	if err != nil {
		fmt.Println("Error connecting to SSH:", err)
	}
	utils.WaitForEnter()
}

func selectConfig(configs []config.SSHConfig) int {
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
		fmt.Println("\nUse ↑↓ arrow keys to navigate, Enter to connect, Esc (twice) to go back")

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

func buildSSHCommand(config config.SSHConfig) []string {
	sshCmd := []string{"ssh"}
	if config.PrivateKey != "" {
		sshCmd = append(sshCmd, "-i", config.PrivateKey)
	}
	sshCmd = append(sshCmd, fmt.Sprintf("%s@%s", config.Username, config.IPAddress))
	return sshCmd
}
