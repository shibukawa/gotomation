package gotomation

import (
	"time"
	"unsafe"
)

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

func (m mouse) ClickWith(button MouseButton) {
	input := wMOUSEINPUT{
		typeCode: wINPUT_MOUSE,
	}
	input.mi.dwFlags = mouseType(true, button) | mouseType(false, button)
	procSendInput.Call(1, uintptr(unsafe.Pointer(&input)), unsafe.Sizeof(input))
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

func (m mouse) DoubleClickWith(button MouseButton) {
	m.ClickWith(button)
	time.Sleep(200 * time.Millisecond)
	m.ClickWith(button)
}

func (m mouse) ScrollQuickly(x, y int) {
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
}
