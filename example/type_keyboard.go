package main

import (
	"fmt"
	"github.com/shibukawa/gotomation"
)

func main() {
	fmt.Println("Type Chars (300 chars/minutes)")
	gotomation.Keyboard.SetTypeSpeed(300)
	gotomation.Keyboard.Type("Hello World ハローワールド 🙆")
	fmt.Println("Press by key code")
	gotomation.Keyboard.KeyDown(gotomation.VK_SHIFT)
	gotomation.Keyboard.KeyPress(gotomation.VK_H)
	gotomation.Keyboard.KeyPress(gotomation.VK_E)
	gotomation.Keyboard.KeyPress(gotomation.VK_L)
	gotomation.Keyboard.KeyPress(gotomation.VK_L)
	gotomation.Keyboard.KeyPress(gotomation.VK_O)
	gotomation.Keyboard.KeyUp(gotomation.VK_SHIFT)

}
