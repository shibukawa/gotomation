package main

import (
	"fmt"
	"github.com/shibukawa/gotomation"
)

func main() {
	fmt.Println("Type Chars (300 chars/minutes)")
	gotomation.Keyboard.SetTypeSpeed(300)
	gotomation.Keyboard.Type("Hello World ハローワールド 🙆")
}
