package gotomation

import (
	"math/rand"
	"time"
	"unsafe"
	"unicode/utf16"
)

const (
	VK_NOT_A_KEY KeyCode = 9999

	VK_A            = 0x41
	VK_S            = 0x53
	VK_D            = 0x44
	VK_F            = 0x46
	VK_H            = 0x48
	VK_G            = 0x47
	VK_Z            = 0x5a
	VK_X            = 0x58
	VK_C            = 0x43
	VK_V            = 0x56
	VK_B            = 0x42
	VK_Q            = 0x51
	VK_W            = 0x57
	VK_E            = 0x45
	VK_R            = 0x52
	VK_Y            = 0x59
	VK_T            = 0x54
	VK_1            = 0x31
	VK_2            = 0x32
	VK_3            = 0x33
	VK_4            = 0x34
	VK_6            = 0x36
	VK_5            = 0x35
	VK_EQUAL        = 0xbb // VK_OEM_PLUS
	VK_9            = 0x39
	VK_7            = 0x37
	VK_MINUS        = 0xbd // VK_OEM_MINUS
	VK_8            = 0x38
	VK_0            = 0x30
	VK_RIGHTBRACKET = 0xdd // VK_OEM_6
	VK_O            = 0x4f
	VK_U            = 0x55
	VK_LEFTBRACKET  = 0xdb // VK_OEM_4
	VK_I            = 0x49
	VK_P            = 0x50
	VK_L            = 0x4c
	VK_J            = 0x4a
	VK_QUOTE        = 0xde // VK_OEM_7
	VK_K            = 0x4b
	VK_SEMICOLON    = 0xdf // VK_OEM_8
	VK_BACKSLASH    = 0xdc // VK_OEM_5
	VK_COMMA        = 0xbc // VK_OEM_COMMA
	VK_SLASH        = 0xbf // C.VK_OEM_2
	VK_N            = 0x4e
	VK_M            = 0x4d
	VK_PERIOD       = 0xbe // C.VK_OEM_PERIOD
	VK_GRAVE        = 0xbf // VK_OEM_3

	VK_BACKSPACE   =  0x08 // VK_BACK
	VK_DELETE      = 0x2e // VK_DELETE
	VK_RETURN      = 0x0d // VK_RETURN
	VK_TAB         = 0x09 // VK_TAB
	VK_ESCAPE      = 0x1b // VK_ESCAPE
	VK_UP          = 0x26 // VK_UP
	VK_DOWN        = 0x28 // VK_DOWN
	VK_RIGHT       = 0x27 // VK_RIGHT
	VK_LEFT        = 0x25 // VK_LEFT
	VK_HOME        = 0x24 // VK_HOME
	VK_END         = 0x23 // VK_END
	VK_PAGEUP      = 0x21 // VK_PRIOR
	VK_PAGEDOWN    = 0x22 // VK_NEXT
	VK_F1          = 0x70 // VK_F1
	VK_F2          = 0x71 // VK_F2
	VK_F3          = 0x72 // VK_F3
	VK_F4          = 0x73 // VK_F4
	VK_F5          = 0x74 // VK_F5
	VK_F6          = 0x75 // VK_F6
	VK_F7          = 0x76 // VK_F7
	VK_F8          = 0x77 // VK_F8
	VK_F9          = 0x78 // VK_F9
	VK_F10         = 0x79 // VK_F10
	VK_F11         = 0x7a // VK_F11
	VK_F12         = 0x7b // VK_F12
	VK_F13         = 0x7c // VK_F13
	VK_F14         = 0x7d // VK_F14
	VK_F15         = 0x7e // VK_F15
	VK_F16         = 0x7f // VK_F16
	VK_F17         = 0x80 // VK_F17
	VK_F18         = 0x81 // VK_F18
	VK_F19         = 0x82 // VK_F19
	VK_F20         = 0x83 // VK_F20
	VK_ALT         = 0x12 // VK_MENU
	VK_CONTROL     = 0x11 // VK_CONTROL
	VK_LCONTROL     = 0xa2 // VK_LCONTROL
	VK_RCONTROL     = 0xa3 // VK_RCONTROL
	VK_SHIFT       = 0x10 // VK_SHIFT
	VK_LSHIFT       = 0xa0 // VK_LSHIFT
	VK_RSHIFT       = 0xa1 // VK_RSHIFT
	VK_LMENU        = 0xa4 // VK_LMENU
	VK_RMENU        = 0xa5 // VK_RMENU
	VK_LWIN        = 0x5b // VK_LWIN
	VK_RWIN        = 0x5c // VK_RWIN
	VK_META        = 0x5b // VK_LWIN
	VK_LMETA        = 0x5b // VK_LWIN
	VK_RMETA        = 0x5c // VK_RWIN
	VK_LCOMMAND        = 0x5b // VK_LWIN
	VK_RCOMMAND        = 0x5c // VK_RWIN
	VK_CAPSLOCK    = 0x14 // VK_CAPITAL
	VK_SPACE       = 0x20 // VK_SPACE
	VK_INSERT      = 0x2d // VK_INSERT
	VK_SNAPSHOT    = 0x2c // VK_SNAPSHOT

	VK_CAPITAL     = 0x14 // VK_CAPITAL
	VK_NUMLOCK     = 0x90 // VK_NUMLOCK
	VK_SCROLL      = 0x91 // VK_SCROLL

	VK_NUMPAD_0       = 0x60 // VK_NUMPAD0
	VK_NUMPAD_1       = 0x61 // VK_NUMPAD1
	VK_NUMPAD_2       = 0x62 // VK_NUMPAD2
	VK_NUMPAD_3       = 0x63 // VK_NUMPAD3
	VK_NUMPAD_4       = 0x64 // VK_NUMPAD4
	VK_NUMPAD_5       = 0x65 // VK_NUMPAD5
	VK_NUMPAD_6       = 0x66 // VK_NUMPAD6
	VK_NUMPAD_7       = 0x67 // VK_NUMPAD7
	VK_NUMPAD_8       = 0x68 // VK_NUMPAD8
	VK_NUMPAD_9       = 0x69 // VK_NUMPAD9
	VK_NUMPAD_DECIMAL = 0x6e // VK_DECIMAL
	VK_NUMPAD_PLUS    = 0x6b // VK_ADD
	VK_NUMPAD_MINUS   = 0x6d // VK_SUBTRACT
	VK_NUMPAD_MUL     = 0x6a // VK_MULTIPLY
	VK_NUMPAD_DIV     = 0x6f // VK_DIVIDE
	VK_NUMPAD_CLEAR   = VK_NOT_A_KEY
	VK_NUMPAD_ENTER   = VK_NOT_A_KEY
	VK_NUMPAD_EQUAL   = VK_NOT_A_KEY

	VK_AUDIO_VOLUME_MUTE = 0xad // VK_VOLUME_MUTE
	VK_AUDIO_VOLUME_DOWN = 0xae // VK_VOLUME_DOWN
	VK_AUDIO_VOLUME_UP   = 0xaf // VK_VOLUME_UP
	VK_AUDIO_PLAY        = 0xb3 // VK_MEDIA_PLAY_PAUSE
	VK_AUDIO_STOP        = 0xb2 // VK_MEDIA_STOP
	VK_AUDIO_PREV        = 0xb1 // VK_MEDIA_PREV_TRACK
	VK_AUDIO_NEXT        = 0xb0 // VK_MEDIA_NEXT_TRACK

	VK_LIGHTS_MON_UP     = 1002
	VK_LIGHTS_MON_DOWN   = 1003
	VK_LIGHTS_KBD_TOGGLE = 1023
	VK_LIGHTS_KBD_UP     = 1021
	VK_LIGHTS_KBD_DOWN   = 1022

	VK_YEN          = VK_NOT_A_KEY
	VK_UNDERSCORE   = 0xe2 // VK_OEM_102
	VK_KEYPAD_COMMA = 0x6c // VK_SEPARATOR
	VK_EISU         = 0xf6 // VK_ATTN
	VK_KANA         = 0x15 // VK_KANA
	VK_HANGUL       = 0x15 // VK_HANGUL
	VK_JUNJA        = 0x17 // VK_JUNJA
	VK_FINAL        = 0x18 // VK_FINAL
)

var extendedKeys = map[KeyCode]bool{
	VK_RCONTROL: true,
	VK_SNAPSHOT: true,
	VK_RMENU: true,
	//VK_PAUSE: true,
	VK_HOME: true,
	VK_UP: true,
	VK_PAGEUP: true,
	VK_LEFT: true,
	VK_RIGHT: true,
	VK_END: true,
	VK_DOWN: true,
	VK_PAGEDOWN: true,
	VK_INSERT: true,
	VK_DELETE: true,
	VK_LWIN: true,
	VK_RWIN: true,
	//VK_APPS: true,
	VK_AUDIO_VOLUME_MUTE: true,
	VK_AUDIO_VOLUME_DOWN: true,
	VK_AUDIO_VOLUME_UP: true,
	VK_AUDIO_PLAY: true,
	VK_AUDIO_STOP: true,
	VK_AUDIO_NEXT: true,
	VK_AUDIO_PREV: true,
	//VK_BROWSER_BACK:
	//VK_BROWSER_FORWARD: true,
	//VK_BROWSER_REFRESH: true,
	//VK_BROWSER_STOP: true,
	//VK_BROWSER_SEARCH: true,
	//VK_BROWSER_FAVORITES: true,
	//VK_BROWSER_HOME: true,
	//VK_LAUNCH_MAIL: true,
}

func keyEvent(key KeyCode, flag uint32) {
	if extendedKeys[key] {
		flag |= wKEYEVENTF_EXTENDEDKEY
	}
	scan, _, _ := procMapVirtualKey.Call(uintptr(int(key) & 0xff), 0)
	if (flag & wKEYEVENTF_KEYUP) == wKEYEVENTF_KEYUP {
		scan |= 0x80
	}
	input := wKEYBDINPUT{
		typeCode: wINPUT_KEYBOARD,
	}
	input.ki.wVk = uint16(key)
	input.ki.wScan = uint16(scan)
	input.ki.dwFlags = flag
	procSendInput.Call(1, uintptr(unsafe.Pointer(&input)), unsafe.Sizeof(input))
	time.Sleep(10 * time.Millisecond)
}

func (k keyboard) toggleKeyByCode(code KeyCode, down bool, modifiers []KeyModifier) {
	var dwFlags uint32
	if !down {
		dwFlags = wKEYEVENTF_KEYUP
	}
	for _, modifier := range modifiers {
		switch modifier {
		case SHIFT:
			keyEvent(VK_SHIFT, dwFlags)
		case ALT:
			keyEvent(VK_ALT, dwFlags)
		case CONTROL:
			keyEvent(VK_CONTROL, dwFlags)
		case META:
			keyEvent(VK_META, dwFlags)
		}
		time.Sleep(time.Duration(rand.Int31n(63)) * time.Millisecond)
	}
	keyEvent(code, dwFlags)
}

func (k keyboard) Type(str string) {
	codes := utf16.Encode([]rune(str))
	if len(codes) == 0 {
		return
	}
	inputs := make([]wKEYBDINPUT, len(codes) * 2)
	for i, code := range codes {
		for j := 0; j < 2; j++ {
			offset := i * 2 + j
			inputs[offset].typeCode = wINPUT_KEYBOARD
			inputs[offset].ki.dwFlags = wKEYEVENTF_UNICODE
			inputs[offset].ki.wScan = code
		}
		inputs[i * 2 + 1].ki.dwFlags |= wKEYEVENTF_KEYUP
	}
	procSendInput.Call(uintptr(len(inputs)), uintptr(unsafe.Pointer(&inputs[0])), unsafe.Sizeof(inputs[0]))
}

func (k keyboard) TypeQuickly(str string) {
	k.Type(str)
}
