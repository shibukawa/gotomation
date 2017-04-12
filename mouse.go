package gotomation

import (
	"time"
)

type MouseButton int

const (
	MouseLeft   MouseButton = 1
	MouseCenter             = 2
	MouseRight              = 3
)

type mouse struct{}

var Mouse = mouse{}

func easeInOutCubic(values [][2]int, duration time.Duration, callback func([]int)) {
	count := int(duration / (time.Millisecond * 16))
	delta := make([]float64, len(values))
	finalValue := make([]int, len(values))
	lastValue := make([]int, len(values))
	for i, value := range values {
		delta[i] = float64(value[1] - value[0])
		finalValue[i] = value[1]
		lastValue[i] = value[0]
	}
	if count == 0 {
		callback(finalValue)
		return
	}

	for f := 0; f < count; f++ {
		time.Sleep(time.Millisecond * 16)

		if f+1 == count {
			callback(finalValue)
			break
		}

		t := float64(f) / float64(count) * 2.0
		var dt float64
		if t < 1.0 {
			dt = 0.5 * t * t * t
		} else {
			t -= 2
			dt = 0.5*t*t*t + 1.0
		}
		currentValues := make([]int, len(values))
		for i, value := range values {
			currentValues[i] = value[0] + int(delta[i]*dt)
		}
		callback(currentValues)
	}
}

/*
	MoveMouse moves mouse cursor with Robert Penner's Easing Function: easeInOutCubic

	http://robertpenner.com/easing/
*/
func (m mouse) Move(x, y int, duration time.Duration) {
	sx, sy := m.GetPosition()
	easeInOutCubic([][2]int{{sx, x}, {sy, y}}, duration, func(value []int) {
		m.MoveQuickly(value[0], value[1])
	})
}

func (m mouse) Click() {
	m.ClickWith(MouseLeft)
}

func (m mouse) Drag(x, y int) {
	m.DragWith(MouseLeft, x, y)
}

func (m mouse) Scroll(x, y int, duration time.Duration) {
	lastValues := []int{0, 0}
	easeInOutCubic([][2]int{{0, x}, {0, y}}, duration, func(values []int) {
		m.ScrollQuickly(values[0]-lastValues[0], values[1]-lastValues[1])
		copy(lastValues, values)
	})
}
