package main

import (
	"fmt"
	"github.com/shibukawa/gotomation"
)

func main() {
	x, y := gotomation.Mouse.GetPosition()
	fmt.Printf("Mouse Position: %d, %d\n", x, y)
}
