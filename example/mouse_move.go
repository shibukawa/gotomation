package main

import (
	"fmt"
	"github.com/shibukawa/gotomation"
	"time"
)

func main() {
	screen, err := gotomation.GetMainScreen()
	if err != nil {
		panic(err)
	}
	mouse := screen.Mouse()
	fmt.Println("Move mouse to center")
	mouse.Move(0, 0, time.Millisecond*500)
	mouse.Move(screen.W(), screen.H(), time.Millisecond*500)
	mouse.Move(0, screen.H(), time.Millisecond*500)
	mouse.Move(screen.W(), 0, time.Millisecond*500)
	mouse.Move(screen.W()/2, screen.H()/2, time.Millisecond*500)
	fmt.Println("Click left button")
	mouse.Click()
	fmt.Println("Double click")
	mouse.DoubleClick()
	fmt.Println("Drag")
	mouse.Drag(screen.W()/2, screen.H()/2+100)
	fmt.Println("Scroll x")
	mouse.Scroll(50, 0, time.Second)
	time.Sleep(time.Second)
	fmt.Println("Scroll -x")
	mouse.Scroll(-50, 0, time.Second)
	time.Sleep(time.Second)
	fmt.Println("Scroll y")
	mouse.Scroll(0, 50, time.Second)
	time.Sleep(time.Second)
	fmt.Println("Scroll -y")
	mouse.Scroll(0, -50, time.Second)
	time.Sleep(time.Second)
}
