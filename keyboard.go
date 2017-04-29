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

type Keyboard interface {
	SetTypeSpeed(charPerMin int)
	TypeSpeed() int
	KeyPress(code KeyCode, modifiers ...KeyModifier) error
	KeyDown(code KeyCode, modifiers ...KeyModifier) error
	KeyUp(code KeyCode, modifiers ...KeyModifier) error
	Type(str string) error
	TypeQuickly(str string) error
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
