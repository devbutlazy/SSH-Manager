package utils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func ClearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func WaitForEnter() {
	fmt.Println("\n[ ~ ] Press enter to return to main menu...")
	bufio.NewScanner(os.Stdin).Scan()
}

func ReadInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
