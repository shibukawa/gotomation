/*
   Copyright 2017, Yoshiki Shibukawa

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

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

type Mouse interface {
	GetPosition() (x, y int)
	Move(x, y int, duration time.Duration) error
	MoveQuickly(x, y int) error
	Click() error
	ClickWith(button MouseButton) error
	DoubleClick() error
	DoubleClickWith(button MouseButton) error
	Drag(x, y int) error
	DragWith(button MouseButton, x, y int) error
	Scroll(x, y int, duration time.Duration) error
}

func easeInOutCubic(values [][2]int, duration time.Duration, callback func([]int) error) error {
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
		return nil
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
		err := callback(currentValues)
		if err != nil {
			return err
		}
	}
	return nil
}

/*
	MoveMouse moves mouse cursor with Robert Penner's Easing Function: easeInOutCubic

	http://robertpenner.com/easing/
*/
func (m mouse) Move(x, y int, duration time.Duration) error {
	sx, sy := m.GetPosition()
	return easeInOutCubic([][2]int{{sx, x}, {sy, y}}, duration, func(value []int) error {
		return m.MoveQuickly(value[0], value[1])
	})
}

func (m mouse) Click() error {
	return m.ClickWith(MouseLeft)
}

func (m mouse) DoubleClick() error {
	return m.DoubleClickWith(MouseLeft)
}

func (m mouse) Drag(x, y int) error {
	return m.DragWith(MouseLeft, x, y)
}

func (m mouse) Scroll(x, y int, duration time.Duration) error {
	lastValues := []int{0, 0}
	return easeInOutCubic([][2]int{{0, x}, {0, y}}, duration, func(values []int) error {
		err := m.ScrollQuickly(values[0]-lastValues[0], values[1]-lastValues[1])
		copy(lastValues, values)
		return err
	})
}
