package gotomation

import "syscall"

var (
	user32               = syscall.MustLoadDLL("user32.dll")
	procGetDesktopWindow = user32.MustFindProc("GetDesktopWindow")
	procGetDC            = user32.MustFindProc("GetDC")
	procReleaseDC        = user32.MustFindProc("ReleaseDC")
	procGetCursorPos     = user32.MustFindProc("GetCursorPos")
	procSetCursorPos     = user32.MustFindProc("SetCursorPos")
	procGetSystemMetrics = user32.MustFindProc("GetSystemMetrics")
	procSendInput        = user32.MustFindProc("SendInput")
	procMapVirtualKey    = user32.MustFindProc("MapVirtualKeyW")

	gdi32                  = syscall.MustLoadDLL("gdi32.dll")
	procGetDeviceCaps      = gdi32.MustFindProc("GetDeviceCaps")
	procCreateDIBSection   = gdi32.MustFindProc("CreateDIBSection")
	procCreateCompatibleDC = gdi32.MustFindProc("CreateCompatibleDC")
	procSelectObject       = gdi32.MustFindProc("SelectObject")
	procDeleteObject       = gdi32.MustFindProc("DeleteObject")
	procDeleteDC           = gdi32.MustFindProc("DeleteDC")
	procBitBlt             = gdi32.MustFindProc("BitBlt")
)

const (
	wHORZRES        = 8
	wVERTRES        = 10
	wBI_RGB         = 0
	wDIB_RGB_COLORS = 0
	wSRCCOPY        = 0xCC0020

	wSM_CXSCREEN    = 0
	wSM_CYSCREEN    = 1
	wINPUT_MOUSE    = 0
	wINPUT_KEYBOARD = 1

	wMOUSEEVENTF_MOVE        uint32 = 0x0001
	wMOUSEEVENTF_LEFTDOWN           = 0x0002
	wMOUSEEVENTF_LEFTUP             = 0x0004
	wMOUSEEVENTF_RIGHTDOWN          = 0x0008
	wMOUSEEVENTF_RIGHTUP            = 0x0010
	wMOUSEEVENTF_MIDDLEDOWN         = 0x0020
	wMOUSEEVENTF_MIDDLEUP           = 0x0040
	wMOUSEEVENTF_WHEEL              = 0x0800
	wMOUSEEVENTF_HWHEEL             = 0x1000
	wMOUSEEVENTF_VIRTUALDESK        = 0x4000
	wMOUSEEVENTF_ABSOLUTE           = 0x8000

	wKEYEVENTF_EXTENDEDKEY = 0x0001
	wKEYEVENTF_KEYUP       = 0x0002
	wKEYEVENTF_UNICODE     = 0x0004
)

type wPOINT struct {
	x int32
	y int32
}

type wBITMAPINFO struct {
	bmiHeader struct {
		biSize          uint32
		biWidth         int32
		biHeight        int32
		biPlanes        uint16
		biBitCount      uint16
		biCompression   uint32
		biSizeImage     uint32
		biXPelsPerMeter int32
		biYPelsPerMeter int32
		biClrUsed       uint32
		biClrImportant  uint32
	}
	bmiColors [1]struct {
		rgbBlue     uint8
		rgbGreen    uint8
		rgbRed      uint8
		rgbReserved uint8
	}
}

type wMOUSEINPUT struct {
	typeCode uint32
	mi       struct {
		dx          int32
		dy          int32
		mouseData   int32
		dwFlags     uint32
		time        uint32
		dwExtraInfo uintptr
	}
}

type wKEYBDINPUT struct {
	typeCode uint32
	ki       struct {
		wVk         uint16
		wScan       uint16
		dwFlags     uint32
		time        uint32
		dwExtraInfo uintptr
	}
}

func isError(err error) bool {
	eno, ok := err.(syscall.Errno)
	if !ok {
		panic(err)
	}
	return uintptr(eno) != 0
}
