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

func (k keyboard) KeyPress(code KeyCode, modifiers ...KeyModifier) error {
	err := k.toggleKeyByCode(code, true, modifiers)
	if err != nil {
		return err
	}
	time.Sleep(10 * time.Millisecond)
	err = k.toggleKeyByCode(code, false, modifiers)
	if err != nil {
		return err
	}
	time.Sleep(10 * time.Millisecond)
	return nil
}

func (k keyboard) KeyDown(code KeyCode, modifiers ...KeyModifier) error {
	err := k.toggleKeyByCode(code, true, modifiers)
	if err != nil {
		return err
	}
	time.Sleep(10 * time.Millisecond)
	return nil
}

func (k keyboard) KeyUp(code KeyCode, modifiers ...KeyModifier) error {
	err := k.toggleKeyByCode(code, false, modifiers)
	if err != nil {
		return err
	}
	time.Sleep(10 * time.Millisecond)
	return nil
}
