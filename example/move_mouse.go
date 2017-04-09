package main

import (
	"fmt"
	"github.com/shibukawa/gotomation"
	"time"
)

func main() {
	x, y := gotomation.Mouse.GetPosition()
	fmt.Printf("Mouse Position: %d, %d\n", x, y)

	screen, err := gotomation.GetMainScreen()
	if err != nil {
		panic(err)
	}
	fmt.Println("Move mouse to center")
	gotomation.Mouse.Move(0, 0, time.Millisecond*500)
	gotomation.Mouse.Move(screen.W(), screen.H(), time.Millisecond*500)
	gotomation.Mouse.Move(0, screen.H(), time.Millisecond*500)
	gotomation.Mouse.Move(screen.W(), 0, time.Millisecond*500)
	gotomation.Mouse.Move(screen.W()/2, screen.H()/2, time.Millisecond*500)
	fmt.Println("Click left button")
	gotomation.Mouse.Click()
	fmt.Println("Double click")
	gotomation.Mouse.DoubleClick()
	fmt.Println("Drag")
	gotomation.Mouse.Drag(screen.W()/2, screen.H()/2+100)
	fmt.Println("Scroll x")
	gotomation.Mouse.Scroll(50, 0, time.Second)
	time.Sleep(time.Second)
	fmt.Println("Scroll -x")
	gotomation.Mouse.Scroll(-50, 0, time.Second)
	time.Sleep(time.Second)
	fmt.Println("Scroll y")
	gotomation.Mouse.Scroll(0, 50, time.Second)
	time.Sleep(time.Second)
	fmt.Println("Scroll -y")
	gotomation.Mouse.Scroll(0, -50, time.Second)
	time.Sleep(time.Second)
}
