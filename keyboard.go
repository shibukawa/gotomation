package gotomation

import "time"

type KeyCode uint16

type KeyModifier uint16

const (
	SHIFT   KeyModifier = 0x0001
	ALT                 = 0x0002
	CONTROL             = 0x0004
	META                = 0x0008
	WIN                 = META
	COMMAND             = META
)

type keyboard struct {
	waitBetweenChars time.Duration // delay
}

var Keyboard = keyboard{
	// default is 1200 chars per minute
	waitBetweenChars: 50 * time.Millisecond,
}

func (k *keyboard) SetTypeSpeed(charPerMin int) {
	k.waitBetweenChars = time.Minute / time.Duration(charPerMin)
}

func (k keyboard) TypeSpeed() int {
	return int(time.Minute / k.waitBetweenChars)
}

func (k keyboard) KeyPress(code KeyCode, modifiers ...KeyModifier) {
	k.toggleKeyByCode(code, true, modifiers)
	time.Sleep(10 * time.Millisecond)
	k.toggleKeyByCode(code, false, modifiers)
	time.Sleep(10 * time.Millisecond)
}

func (k keyboard) KeyDown(code KeyCode, modifiers ...KeyModifier) {
	k.toggleKeyByCode(code, true, modifiers)
	time.Sleep(10 * time.Millisecond)
}

func (k keyboard) KeyUp(code KeyCode, modifiers ...KeyModifier) {
	k.toggleKeyByCode(code, false, modifiers)
	time.Sleep(10 * time.Millisecond)
}
