// +build !windows,!darwin,!cgo

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
	"errors"
	"github.com/BurntSushi/xgb/xproto"
	"math/rand"
	"time"
)

const (
	VK_NOT_A_KEY KeyCode = 9999

	VK_A            = 30
	VK_S            = 31
	VK_D            = 32
	VK_F            = 33
	VK_H            = 35
	VK_G            = 34
	VK_Z            = 44
	VK_X            = 45
	VK_C            = 46
	VK_V            = 47
	VK_B            = 48
	VK_Q            = 16
	VK_W            = 17
	VK_E            = 18
	VK_R            = 19
	VK_Y            = 21
	VK_T            = 20
	VK_1            = 2
	VK_2            = 3
	VK_3            = 4
	VK_4            = 5
	VK_6            = 7
	VK_5            = 6
	VK_EQUAL        = 13
	VK_9            = 10
	VK_7            = 8
	VK_MINUS        = 12
	VK_8            = 9
	VK_0            = 11
	VK_RIGHTBRACKET = 27
	VK_O            = 24
	VK_U            = 22
	VK_LEFTBRACKET  = 26
	VK_I            = 23
	VK_P            = 25
	VK_L            = 38
	VK_J            = 36
	VK_QUOTE        = 40
	VK_K            = 37
	VK_SEMICOLON    = 39
	VK_BACKSLASH    = 43
	VK_COMMA        = 51
	VK_SLASH        = 53
	VK_N            = 49
	VK_M            = 50
	VK_PERIOD       = 52
	VK_GRAVE        = 41

	VK_BACKSPACE = 14
	VK_DELETE    = 111
	VK_RETURN    = 28
	VK_TAB       = 15
	VK_ESCAPE    = 1
	VK_UP        = 103
	VK_DOWN      = 108
	VK_RIGHT     = 106
	VK_LEFT      = 105
	VK_HOME      = 102
	VK_END       = 107
	VK_PAGEUP    = 104
	VK_PAGEDOWN  = 109
	VK_F1        = 59 // VK_F1
	VK_F2        = 60 // VK_F2
	VK_F3        = 61 // VK_F3
	VK_F4        = 62 // VK_F4
	VK_F5        = 63 // VK_F5
	VK_F6        = 64 // VK_F6
	VK_F7        = 65 // VK_F7
	VK_F8        = 66 // VK_F8
	VK_F9        = 67 // VK_F9
	VK_F10       = 68 // VK_F10
	VK_F11       = 87 // VK_F11
	VK_F12       = 88 // VK_F12
	VK_F13       = VK_NOT_A_KEY
	VK_F14       = VK_NOT_A_KEY
	VK_F15       = VK_NOT_A_KEY
	VK_F16       = VK_NOT_A_KEY
	VK_F17       = VK_NOT_A_KEY
	VK_F18       = VK_NOT_A_KEY
	VK_F19       = VK_NOT_A_KEY
	VK_F20       = VK_NOT_A_KEY
	VK_ALT       = 56
	VK_LALT      = 56
	VK_RALT      = 100
	VK_CONTROL   = 29
	VK_LCONTROL  = 29
	VK_RCONTROL  = 97
	VK_SHIFT     = 42
	VK_LSHIFT    = 42
	VK_RSHIFT    = 54
	VK_LMENU     = VK_NOT_A_KEY
	VK_RMENU     = VK_NOT_A_KEY
	VK_LWIN      = VK_NOT_A_KEY
	VK_RWIN      = VK_NOT_A_KEY
	VK_META      = VK_NOT_A_KEY
	VK_LMETA     = VK_NOT_A_KEY
	VK_RMETA     = VK_NOT_A_KEY
	VK_LCOMMAND  = VK_NOT_A_KEY
	VK_RCOMMAND  = VK_NOT_A_KEY
	VK_CAPSLOCK  = 58
	VK_SPACE     = 57
	VK_INSERT    = 100
	VK_SNAPSHOT  = VK_NOT_A_KEY
	VK_NUMLOCK   = 69
	VK_SCROLL    = 70

	VK_NUMPAD_0       = VK_NOT_A_KEY
	VK_NUMPAD_1       = 79
	VK_NUMPAD_2       = 80
	VK_NUMPAD_3       = 81
	VK_NUMPAD_4       = 75
	VK_NUMPAD_5       = 76
	VK_NUMPAD_6       = 77
	VK_NUMPAD_7       = 71
	VK_NUMPAD_8       = 72
	VK_NUMPAD_9       = 73
	VK_NUMPAD_DECIMAL = VK_NOT_A_KEY
	VK_NUMPAD_PLUS    = 78
	VK_NUMPAD_MINUS   = 74
	VK_NUMPAD_MUL     = 55
	VK_NUMPAD_DIV     = 98
	VK_NUMPAD_CLEAR   = VK_NOT_A_KEY
	VK_NUMPAD_ENTER   = 96
	VK_NUMPAD_EQUAL   = VK_NOT_A_KEY

	VK_AUDIO_VOLUME_MUTE = VK_NOT_A_KEY
	VK_AUDIO_VOLUME_DOWN = VK_NOT_A_KEY
	VK_AUDIO_VOLUME_UP   = VK_NOT_A_KEY
	VK_AUDIO_PLAY        = VK_NOT_A_KEY
	VK_AUDIO_STOP        = VK_NOT_A_KEY
	VK_AUDIO_PREV        = VK_NOT_A_KEY
	VK_AUDIO_NEXT        = VK_NOT_A_KEY

	VK_LIGHTS_MON_UP     = VK_NOT_A_KEY
	VK_LIGHTS_MON_DOWN   = VK_NOT_A_KEY
	VK_LIGHTS_KBD_TOGGLE = VK_NOT_A_KEY
	VK_LIGHTS_KBD_UP     = VK_NOT_A_KEY
	VK_LIGHTS_KBD_DOWN   = VK_NOT_A_KEY

	VK_YEN          = VK_NOT_A_KEY
	VK_UNDERSCORE   = VK_NOT_A_KEY
	VK_KEYPAD_COMMA = VK_NOT_A_KEY
	VK_EISU         = VK_NOT_A_KEY
	VK_KANA         = VK_NOT_A_KEY
	VK_HANGUL       = VK_NOT_A_KEY
	VK_JUNJA        = VK_NOT_A_KEY
	VK_FINAL        = VK_NOT_A_KEY
)

type keyboard struct {
	waitBetweenChars time.Duration // delay
	screen           *screen
}

func (k keyboard) toggleKeyByCode(code KeyCode, down bool, modifiers []KeyModifier) error {
	var typ xXCB_TYPE
	if down {
		typ = xXCB_KEY_PRESS
	} else {
		typ = xXCB_KEY_RELEASE
	}
	for _, modifier := range modifiers {
		var err error
		switch modifier {
		case SHIFT:
			err = fakeInput(k.screen.conn, k.screen.screen, VK_SHIFT, typ)
		case ALT:
			err = fakeInput(k.screen.conn, k.screen.screen, VK_ALT, typ)
		case CONTROL:
			err = fakeInput(k.screen.conn, k.screen.screen, VK_CONTROL, typ)
		case META:
			err = fakeInput(k.screen.conn, k.screen.screen, VK_META, typ)
		}
		if err != nil {
			return err
		}
		time.Sleep(time.Duration(rand.Int31n(63)) * time.Millisecond)
	}
	return fakeInput(k.screen.conn, k.screen.screen, code, typ)
}

func (k keyboard) Type(str string) error {
	return errors.New("Type feature is not working without cgo")
}
