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
	fmt.Println("[ ~ ] Press Enter to continue...")
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
}

func ReadInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)

	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
