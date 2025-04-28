package game

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"golang.org/x/term"
)

func ClearScreen() {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Countdown(seconds int) {
	for i := seconds; i > 0; i-- {
		fmt.Printf("\rNext action in: %d seconds", i)
		time.Sleep(1 * time.Second)
	}
	fmt.Print("\n\r\r")
}

func CenterTextSmart(text string) {
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		width = 80
	}

	lines := strings.Split(text, "\n")
	for _, line := range lines {
		padding := (width - len(line)) / 2
		if padding > 0 {
			fmt.Println(strings.Repeat("", padding) + line)
		} else {
			fmt.Println(line)
		}
	}
}
