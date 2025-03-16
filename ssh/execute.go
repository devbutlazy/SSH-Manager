package ssh

import (
	"fmt"
	"os"
	"os/exec"
	"ssh-manager/config"
	"ssh-manager/utils"
)

func ExecuteSSH(cfg config.SSHConfig) {
	sshCmd := []string{"ssh"}
	if cfg.PrivateKey != "" {
		sshCmd = append(sshCmd, "-i", cfg.PrivateKey)
	}
	sshCmd = append(sshCmd, fmt.Sprintf("%s@%s", cfg.Username, cfg.IPAddress))

	utils.ClearScreen()
	fmt.Printf("Connecting to %s (%s@%s)...\n", cfg.UniqueName, cfg.Username, cfg.IPAddress)

	cmd := exec.Command(sshCmd[0], sshCmd[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("Error connecting to SSH:", err)
		utils.WaitForEnter()
	}
}
