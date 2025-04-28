package game

import (
	"fmt"
	"strings"
)

const SplashArt = `

 ____  _     _      _____ _____ ____  _      ____    ____  _  ____  _____
/  _ \/ \ /\/ \  /|/  __//  __//  _ \/ \  /|/ ___\  /  _ \/ \/   _\/  __/
| | \|| | ||| |\ ||| |  _|  \  | / \|| |\ |||    \  | | \|| ||  /  |  \  
| |_/|| \_/|| | \||| |_//|  /_ | \_/|| | \||\___ |  | |_/|| ||  \_ |  /_ 
\____/\____/\_/  \|\____\\____\\____/\_/  \|\____/  \____/\_/\____/\____\
                                                                        
 
`

func CenterText(text string, width int) {
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

func ShowSplashScreen() {
	ClearScreen()
	CenterTextSmart(SplashArt)
	fmt.Println()
	fmt.Println("Press ENTER to start your adeventure and explore the depths of this eerie dungeon...")
	fmt.Scanln()
}
