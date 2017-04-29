// +build !windows,!darwin

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
)

type xXCB_TYPE byte

const (
	xXCB_KEY_PRESS      xXCB_TYPE = 2
	xXCB_KEY_RELEASE              = 3
	xXCB_BUTTON_PRESS             = 4
	xXCB_BUTTON_RELEASE           = 5
	xXCB_MOTION_NOTIFY            = 6

	xSCROLL_UP    KeyCode = 4
	xSCROLL_DOWN          = 5
	xSCROLL_LEFT          = 6
	xSCROLL_RIGHT         = 7
)

const (
	xMouseScrollUp    MouseButton = 4
	xMouseScrollDown              = 5
	xMouseScrollLeft              = 6
	xMouseScrollRight             = 7
)

func fakeInput(conn *xgb.Conn, screen *xproto.ScreenInfo, code KeyCode, typ xXCB_TYPE) error {
	detail := byte(code)
	time := uint32(0)
	id := byte(0)
	cookie := xtest.FakeInputChecked(conn, byte(typ), detail, time, screen.Root, 0, 0, id)
	return cookie.Check()
}
