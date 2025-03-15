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
	executeSSH(selectedConfig)
}

func executeSSH(cfg config.SSHConfig) {
	sshCmd := buildSSHCommand(cfg)
	utils.ClearScreen()
	fmt.Printf("Connecting to %s (%s@%s)...\n", cfg.UniqueName, cfg.Username, cfg.IPAddress)

	cmd := exec.Command(sshCmd[0], sshCmd[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("Error connecting to SSH:", err)
	}

	utils.WaitForEnter()
}

func selectConfig(configs []config.SSHConfig) int {
	selected := 0
	if oldState, err := term.MakeRaw(int(os.Stdin.Fd())); err == nil {
		defer term.Restore(int(os.Stdin.Fd()), oldState)
	} else {
		fmt.Println("Error setting terminal mode:", err)
		return -1
	}

	for {
		printConfigs(configs, selected)
		input := readKeyPress()
		switch input {
		case "up":
			if selected > 0 {
				selected--
			}
		case "down":
			if selected < len(configs)-1 {
				selected++
			}
		case "enter":
			return selected
		case "esc":
			return -1
		}
	}
}

func printConfigs(configs []config.SSHConfig, selected int) {
	utils.ClearScreen()
	fmt.Print("=== Available SSH Configurations ===\n\n")
	for i, cfg := range configs {
		marker := "  "
		if i == selected {
			marker = "> "
		}
		fmt.Printf("%s[%d] %s (%s@%s)\n", marker, i+1, cfg.UniqueName, cfg.Username, cfg.IPAddress)
	}
	fmt.Println("\nUse ↑↓ arrow keys to navigate, Enter to connect, Esc to go back")
}

func readKeyPress() string {
	b := make([]byte, 3)
	_, err := os.Stdin.Read(b)
	if err != nil {
		return ""
	}

	switch {
	case b[0] == 27 && len(b) > 1 && b[1] == 91:
		switch b[2] {
		case 65:
			return "up"
		case 66:
			return "down"
		}
	case b[0] == 13:
		return "enter"
	case b[0] == 27:
		return "esc"
	}
	return ""
}

func buildSSHCommand(cfg config.SSHConfig) []string {
	cmd := []string{"ssh"}
	if cfg.PrivateKey != "" {
		cmd = append(cmd, "-i", cfg.PrivateKey)
	}
	cmd = append(cmd, fmt.Sprintf("%s@%s", cfg.Username, cfg.IPAddress))
	return cmd
}
