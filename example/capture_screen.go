package main

import (
	"fmt"
	"github.com/shibukawa/gotomation"
	"image/png"
	"os"
)

func main() {
	fmt.Println("Capture Screen")
	screen, err := gotomation.GetMainScreen()
	if err != nil {
		panic(err)
	}
	fmt.Printf("id: %d\n", screen.ID())
	fmt.Printf("w: %d\n", screen.W())
	fmt.Printf("h: %d\n", screen.H())
	image, err := screen.Capture()
	if err != nil {
		panic(err)
	}
	file, err := os.Create("capture.png")
	if err != nil {
		panic(err)
	}
	err = png.Encode(file, image)
	if err != nil {
		panic(err)
	}
}
