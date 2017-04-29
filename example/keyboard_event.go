package main

import (
	"fmt"
	"github.com/shibukawa/gotomation"
	"runtime"
)

func main() {
	fmt.Println("Type 'HELLO'")
	screen, _ := gotomation.GetMainScreen()
	keyboard := screen.Keyboard()
	keyboard.KeyDown(gotomation.VK_SHIFT)
	keyboard.KeyPress(gotomation.VK_H)
	keyboard.KeyPress(gotomation.VK_E)
	keyboard.KeyPress(gotomation.VK_L)
	keyboard.KeyPress(gotomation.VK_L)
	keyboard.KeyPress(gotomation.VK_O)
	keyboard.KeyUp(gotomation.VK_SHIFT)
	if runtime.GOOS == "darwin" {
		fmt.Println("\n\nMake keyboard light bright")
		for i := 0; i < 32; i++ {
			keyboard.KeyPress(gotomation.VK_LIGHTS_KBD_UP)
		}
	}
}
