package main

import (
	"fmt"
	"github.com/shibukawa/gotomation"
)

func main() {
	screen, _ := gotomation.GetMainScreen()
	x, y := screen.Mouse().GetPosition()
	fmt.Printf("Mouse Position: %d, %d\n", x, y)
}
