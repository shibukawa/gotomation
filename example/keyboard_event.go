package main

import (
	"fmt"
	"github.com/shibukawa/gotomation"
	"runtime"
)

func main() {
	fmt.Println("Type 'HELLO'")
	gotomation.Keyboard.KeyDown(gotomation.VK_SHIFT)
	gotomation.Keyboard.KeyPress(gotomation.VK_H)
	gotomation.Keyboard.KeyPress(gotomation.VK_E)
	gotomation.Keyboard.KeyPress(gotomation.VK_L)
	gotomation.Keyboard.KeyPress(gotomation.VK_L)
	gotomation.Keyboard.KeyPress(gotomation.VK_O)
	gotomation.Keyboard.KeyUp(gotomation.VK_SHIFT)
	if runtime.GOOS == "darwin" {
		fmt.Println("Make keyboard light bright")
		for i := 0; i < 32; i++ {
			gotomation.Keyboard.KeyPress(gotomation.VK_LIGHTS_KBD_UP)
		}
	}
}
