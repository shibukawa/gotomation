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
	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgb/xtest"
	"time"
)

type mouse struct {
	screen *screen
}

func (m mouse) GetPosition() (x, y int) {
	r, err := xproto.QueryPointer(m.screen.conn, m.screen.screen.Root).Reply()
	if err != nil {
		return -1, -1
	}
	return int(r.RootX), int(r.RootY)
}

func (m mouse) MoveQuickly(x, y int) error {
	root := m.screen.screen.Root
	none := xproto.Window(0)
	cookie := xproto.WarpPointerChecked(m.screen.conn, none, root, 0, 0, 0, 0, int16(x), int16(y))
	return cookie.Check()
}

func mouseToggleButton(conn *xgb.Conn, screen *xproto.ScreenInfo, down bool, button MouseButton) error {
	var typ xXCB_TYPE
	if down {
		typ = xXCB_BUTTON_PRESS
	} else {
		typ = xXCB_BUTTON_RELEASE
	}

	detail := byte(button)
	time := uint32(0)
	id := byte(0)
	cookie := xtest.FakeInputChecked(conn, byte(typ), detail, time, screen.Root, 0, 0, id)
	return cookie.Check()
}

func (m mouse) ClickWith(button MouseButton) error {
	err := mouseToggleButton(m.screen.conn, m.screen.screen, true, button)
	if err != nil {
		return err
	}
	time.Sleep(time.Millisecond * 10)
	err = mouseToggleButton(m.screen.conn, m.screen.screen, false, button)
	if err != nil {
		return err
	}
	time.Sleep(time.Millisecond * 10)
	return nil
}

func (m mouse) DragWith(button MouseButton, x, y int) error {
	err := mouseToggleButton(m.screen.conn, m.screen.screen, true, button)
	if err != nil {
		return err
	}

	time.Sleep(10 * time.Millisecond)

	err = m.MoveQuickly(x, y)
	if err != nil {
		return err
	}

	time.Sleep(10 * time.Millisecond)

	return mouseToggleButton(m.screen.conn, m.screen.screen, false, button)
}

func (m mouse) DoubleClickWith(button MouseButton) error {
	err := m.ClickWith(button)
	if err != nil {
		return err
	}
	time.Sleep(200 * time.Millisecond)
	return m.ClickWith(button)
}

func (m mouse) ScrollQuickly(x, y int) error {
	if y < 0 {
		for ; y != 0; y++ {
			err := m.ClickWith(xMouseScrollUp)
			if err != nil {
				return err
			}
		}
	} else if y > 0 {
		for ; y != 0; y-- {
			err := m.ClickWith(xMouseScrollDown)
			if err != nil {
				return err
			}
		}
	}
	if x < 0 {
		for ; x != 0; x++ {
			err := m.ClickWith(xMouseScrollLeft)
			if err != nil {
				return err
			}
		}
	} else if x > 0 {
		for ; x != 0; x-- {
			err := m.ClickWith(xMouseScrollRight)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
