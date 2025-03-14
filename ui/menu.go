package ui

import (
	"fmt"
	"os"
	"ssh_manager/ssh"
	"ssh_manager/utils"

	"golang.org/x/term"
)

func printMainMenu(selected int) {
	utils.ClearScreen()

	fmt.Println("=== SSH Configuration Manager ===")
	fmt.Println()
	options := []string{
		"Connect to SSH",
		"Add SSH Configuration",
		"Remove SSH Configuration",
		"Exit",
	}

	for i, option := range options {
		if i == selected {
			fmt.Printf("> [%d] %s\n", i+1, option)
		} else {
			fmt.Printf("  [%d] %s\n", i+1, option)
		}
	}
	fmt.Println("\nUse ↑↓ arrow keys to navigate, Enter to select")
}

func getMenuChoice() int {
	selected := 0
	maxOptions := 3

	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println("Error setting terminal mode:", err)
		return -1
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	for {
		printMainMenu(selected)

		b := make([]byte, 1)
		_, err := os.Stdin.Read(b)
		if err != nil {
			fmt.Println("Error reading input:", err)
			return -1
		}

		if b[0] == 27 {
			b = make([]byte, 2)
			os.Stdin.Read(b)
			
			if b[0] == 91 {
				switch b[1] {
				case 65: // Up arrow
					if selected > 0 {
						selected--
					}
				case 66: // Down arrow
					if selected < maxOptions {
						selected++
					}
				}
			}
		} else if b[0] == 13 { // Enter
			return selected + 1
		}
	}
}

func Run() {
	for {
		choice := getMenuChoice()
		switch choice {
		case 1:
			ssh.Connect()
		case 2:
			ssh.AddConfig()
		case 3:
			ssh.RemoveConfig()
		case 4:
			utils.ClearScreen()
			fmt.Println("=== Goodbye ===")
			os.Exit(0)
		}
	}
}