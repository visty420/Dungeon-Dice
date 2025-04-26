package game

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
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
