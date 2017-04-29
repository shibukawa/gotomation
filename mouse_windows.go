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
	"unsafe"
)

type mouse struct{}

// https://pywinauto.github.io/
// todo: debug

func mouseType(down bool, button MouseButton) (mouseType uint32) {
	if down {
		switch button {
		case MouseLeft:
			mouseType = wMOUSEEVENTF_LEFTDOWN
		case MouseRight:
			mouseType = wMOUSEEVENTF_RIGHTDOWN
		case MouseCenter:
			mouseType = wMOUSEEVENTF_MIDDLEDOWN
		}
	} else {
		switch button {
		case MouseLeft:
			mouseType = wMOUSEEVENTF_LEFTUP
		case MouseRight:
			mouseType = wMOUSEEVENTF_RIGHTUP
		case MouseCenter:
			mouseType = wMOUSEEVENTF_MIDDLEUP
		}
	}
	return
}

func (m mouse) ClickWith(button MouseButton) error {
	input := wMOUSEINPUT{
		typeCode: wINPUT_MOUSE,
	}
	input.mi.dwFlags = mouseType(true, button) | mouseType(false, button)
	_, _, err := procSendInput.Call(1, uintptr(unsafe.Pointer(&input)), unsafe.Sizeof(input))
	if isError(err) {
		return err
	}
	return nil
}

func (m mouse) GetPosition() (x, y int) {
	point := wPOINT{}
	procGetCursorPos.Call(uintptr(unsafe.Pointer(&point)))
	x = int(point.x)
	y = int(point.y)
	return
}

func (m mouse) MoveQuickly(x, y int) error {
	sx, sy := m.GetPosition()
	input := wMOUSEINPUT{
		typeCode: wINPUT_MOUSE,
	}
	input.mi.dx = int32(x - sx)
	input.mi.dy = int32(y - sy)
	input.mi.time = 0
	input.mi.dwFlags = wMOUSEEVENTF_MOVE
	_, _, err := procSendInput.Call(1, uintptr(unsafe.Pointer(&input)), unsafe.Sizeof(input))
	if isError(err) {
		return err
	}
	return nil
}

func (m mouse) DragWith(button MouseButton, x, y int) error {
	input1 := wMOUSEINPUT{
		typeCode: wINPUT_MOUSE,
	}
	input1.mi.dwFlags = mouseType(true, button)
	_, _, err := procSendInput.Call(1, uintptr(unsafe.Pointer(&input1)), unsafe.Sizeof(input1))
	if isError(err) {
		return err
	}

	time.Sleep(10 * time.Millisecond)

	err = m.MoveQuickly(x, y)
	if err != nil {
		return err
	}

	time.Sleep(10 * time.Millisecond)

	input1.mi.dwFlags = mouseType(false, button)
	_, _, err = procSendInput.Call(1, uintptr(unsafe.Pointer(&input1)), unsafe.Sizeof(input1))
	if isError(err) {
		return err
	}
	return nil
}

func (m mouse) DoubleClickWith(button MouseButton) error {
	err := m.ClickWith(button)
	if err != nil {
		return err
	}
	time.Sleep(200 * time.Millisecond)
	return m.ClickWith(button)
}

func (m mouse) ScrollQuickly(x, y int) error {
	input := wMOUSEINPUT{
		typeCode: wINPUT_MOUSE,
	}
	if x != 0 {
		input.mi.dwFlags = wMOUSEEVENTF_HWHEEL
		input.mi.mouseData = int32(120 * x)
		procSendInput.Call(1, uintptr(unsafe.Pointer(&input)), unsafe.Sizeof(input))
		time.Sleep(10 * time.Millisecond)
	}
	if y != 0 {
		input.mi.dwFlags = wMOUSEEVENTF_WHEEL
		input.mi.mouseData = int32(120 * y)
		procSendInput.Call(1, uintptr(unsafe.Pointer(&input)), unsafe.Sizeof(input))
		time.Sleep(10 * time.Millisecond)
	}
	return nil
}
