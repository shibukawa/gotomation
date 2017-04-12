// build! darwin

package gotomation

/*
#cgo LDFLAGS: -framework Carbon
#include <ApplicationServices/ApplicationServices.h>

// to avoid this: https://github.com/golang/go/issues/975
CGEventRef createScrollEvent(int32_t x, int32_t y) {
	return CGEventCreateScrollWheelEvent(NULL, kCGScrollEventUnitLine, 2, x, y);
}
*/
import "C"
import "time"

func rawMousePos() C.CGPoint {
	event := C.CGEventCreate(nil)
	defer C.CFRelease(C.CFTypeRef(event))
	return C.CGEventGetLocation(event)
}

func (m mouse) GetPosition() (x, y int) {
	point := rawMousePos()
	x = int(point.x)
	y = int(point.y)
	return
}

func calculateDeltas(event *C.CGEventRef, x, y int) {
	pos := rawMousePos()

	C.CGEventSetIntegerValueField(*event, C.kCGMouseEventDeltaX, C.int64_t(x)-C.int64_t(pos.x))
	C.CGEventSetIntegerValueField(*event, C.kCGMouseEventDeltaY, C.int64_t(y)-C.int64_t(pos.y))
}

func (m mouse) MoveQuickly(x, y int) error {
	move := C.CGEventCreateMouseEvent(nil, C.kCGEventMouseMoved,
		C.CGPointMake((C.CGFloat)(x), (C.CGFloat)(y)),
		C.kCGMouseButtonLeft)
	defer C.CFRelease(C.CFTypeRef(move))

	calculateDeltas(&move, x, y)

	C.CGEventPost(C.kCGSessionEventTap, move)
	return nil
}

func mouseType(down bool, button MouseButton) (mouseType C.CGEventType) {
	if down {
		switch button {
		case MouseLeft:
			mouseType = C.kCGEventLeftMouseDown
		case MouseRight:
			mouseType = C.kCGEventRightMouseDown
		case MouseCenter:
			mouseType = C.kCGEventOtherMouseDown
		}
	} else {
		switch button {
		case MouseLeft:
			mouseType = C.kCGEventLeftMouseUp
		case MouseRight:
			mouseType = C.kCGEventRightMouseUp
		case MouseCenter:
			mouseType = C.kCGEventOtherMouseUp
		}
	}
	return
}

func mouseToggleButton(down bool, button MouseButton) error {
	event := C.CGEventCreateMouseEvent(nil, mouseType(down, button), rawMousePos(), (C.CGMouseButton)(button))
	defer C.CFRelease(C.CFTypeRef(event))
	C.CGEventPost(C.kCGSessionEventTap, event)
	return nil
}

func (m mouse) ClickWith(button MouseButton) {
	mouseToggleButton(true, button)
	time.Sleep(time.Millisecond * 10)
	mouseToggleButton(false, button)
	time.Sleep(time.Millisecond * 10)
}

func (m mouse) DoubleClick() {
	event := C.CGEventCreateMouseEvent(nil, mouseType(true, MouseLeft), rawMousePos(), C.kCGMouseButtonLeft)
	defer C.CFRelease(C.CFTypeRef(event))

	C.CGEventSetIntegerValueField(event, C.kCGMouseEventClickState, 2)

	C.CGEventPost(C.kCGHIDEventTap, event)

	C.CGEventSetType(event, mouseType(false, MouseLeft))
	C.CGEventPost(C.kCGHIDEventTap, event)
	time.Sleep(time.Millisecond * 100)
}

func (m mouse) ScrollQuickly(x, y int) error {
	event := C.createScrollEvent(C.int32_t(x), C.int32_t(y))
	defer C.CFRelease(C.CFTypeRef(event))
	C.CGEventPost(C.kCGHIDEventTap, event)
	return nil
}

func (m mouse) DragWith(button MouseButton, x, y int) error {
	var dragType C.CGEventType
	switch button {
	case MouseLeft:
		dragType = C.kCGEventLeftMouseDragged
	case MouseRight:
		dragType = C.kCGEventRightMouseDragged
	case MouseCenter:
		dragType = C.kCGEventOtherMouseDragged
	}
	point := C.CGPointMake((C.CGFloat)(x), (C.CGFloat)(y))
	drag := C.CGEventCreateMouseEvent(nil, dragType, point, (C.CGMouseButton)(button))
	defer C.CFRelease(C.CFTypeRef(drag))
	calculateDeltas(&drag, x, y)

	C.CGEventPost(C.kCGSessionEventTap, drag)
	time.Sleep(time.Millisecond * 100)
	return nil
}
