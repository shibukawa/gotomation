package main

import (
	"fmt"
	"github.com/shibukawa/gotomation"
)

func main() {
	fmt.Println("Type Chars (300 chars/minutes)")
	screen, _ := gotomation.GetMainScreen()
	keyboard := screen.Keyboard()
	keyboard.SetTypeSpeed(300)
	keyboard.Type("Hello World ハローワールド 🙆")
}
