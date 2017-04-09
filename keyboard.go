package gotomation

import "time"

type KeyCode uint16

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

func (k keyboard) Type(str string) {
	for _, char := range str {
		k.tap(char)
		time.Sleep(k.waitBetweenChars)
	}
}

func (k keyboard) TypeQuickly(str string) {
	for _, char := range str {
		k.tap(char)
	}
}

func (k keyboard) tap(char rune) {
	k.toggleKeyByRune(char, true)
	k.toggleKeyByRune(char, false)
}

func (k keyboard) KeyPress(code KeyCode) {
	k.toggleKeyByCode(code, true)
	time.Sleep(10 * time.Millisecond)
	k.toggleKeyByCode(code, false)
	time.Sleep(10 * time.Millisecond)
}

func (k keyboard) KeyDown(code KeyCode) {
	k.toggleKeyByCode(code, true)
	time.Sleep(10 * time.Millisecond)
}

func (k keyboard) KeyUp(code KeyCode) {
	k.toggleKeyByCode(code, false)
	time.Sleep(10 * time.Millisecond)
}
