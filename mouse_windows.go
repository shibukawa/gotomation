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

func mouseCoordToAbs(coord, total int) int32 {
	return int32((65536*coord)/total) + 1
}

func mouseMoveEvent(x, y int) (*wMOUSEINPUT, error) {
	input := &wMOUSEINPUT{
		typeCode: wINPUT_MOUSE,
	}

	absX, _, err := procGetSystemMetrics.Call(wSM_CXSCREEN)
	if err != nil {
		return nil, err
	}
	input.mi.dx = mouseCoordToAbs(x, int(absX))
	absY, _, err := procGetSystemMetrics.Call(wSM_CYSCREEN)
	if err != nil {
		return nil, err
	}
	input.mi.dy = mouseCoordToAbs(x, int(absY))
	input.mi.dwFlags = wMOUSEEVENTF_ABSOLUTE | wMOUSEEVENTF_MOVE
	return input, nil
}

func (m mouse) MoveQuickly(x, y int) error {
	input, err := mouseMoveEvent(x, y)
	if err != nil {
		return err
	}
	procSendInput.Call(1, uintptr(unsafe.Pointer(&input)), unsafe.Sizeof(input))
	return nil
}

func (m mouse) DragWith(button MouseButton, x, y int) error {
	input, err := mouseMoveEvent(x, y)
	if err != nil {
		return err
	}
	input.mi.dwFlags = wMOUSEEVENTF_ABSOLUTE | wMOUSEEVENTF_MOVE | mouseType(true, button)
	procSendInput.Call(1, uintptr(unsafe.Pointer(&input)), unsafe.Sizeof(input))
	time.Sleep(10 * time.Millisecond)
	input.mi.dwFlags = wMOUSEEVENTF_ABSOLUTE | wMOUSEEVENTF_MOVE | mouseType(false, button)
	procSendInput.Call(1, uintptr(unsafe.Pointer(&input)), unsafe.Sizeof(input))
	return nil
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
