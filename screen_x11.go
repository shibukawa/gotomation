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
	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgb/xtest"
	"image"
	"time"
)

type screen struct {
	conn    *xgb.Conn
	screen  *xproto.ScreenInfo
	segment uint32
	w, h    int
	id      int
}

func (s screen) getID() int {
	return s.id
}

func newScreen(conn *xgb.Conn, screenInfo *xproto.ScreenInfo, id int) (*Screen, error) {
	segment, err := conn.NewId()
	if err != nil {
		return nil, err
	}
	result := &screen{
		conn:    conn,
		screen:  screenInfo,
		segment: segment,
		id:      id,
	}
	result.w = int(result.screen.WidthInPixels)
	result.h = int(result.screen.HeightInPixels)
	err = xtest.Init(conn)
	if err != nil {
		return nil, err
	}
	return &Screen{
		screen: result,
		mouse: &mouse{
			screen: result,
		},
		keyboard: &keyboard{
			waitBetweenChars: 50 * time.Millisecond,
			screen:           result,
		},
	}, nil
}

func GetMainScreen() (*Screen, error) {
	conn, err := xgb.NewConn()
	if err != nil {
		return nil, err
	}
	setup := xproto.Setup(conn)

	defaultScreen := setup.DefaultScreen(conn)
	id := -1
	for i, screen := range setup.Roots {
		if defaultScreen.Root == screen.Root {
			id = i
			break
		}
	}
	if id == 0 {
		return nil, errors.New("Can't find default screen id")
	}
	return newScreen(conn, defaultScreen, id)
}

func GetScreen(id int) (*Screen, error) {
	conn, err := xgb.NewConn()
	if err != nil {
		return nil, err
	}
	setup := xproto.Setup(conn)
	return newScreen(conn, &setup.Roots[id], id)
}

func (s *screen) capture(rect image.Rectangle) (image.Image, error) {
	w, h := rect.Dx(), rect.Dy()
	xImg, err := xproto.GetImage(s.conn, xproto.ImageFormatZPixmap, xproto.Drawable(s.screen.Root), int16(rect.Min.X), int16(rect.Min.Y), uint16(w), uint16(h), 0xffffffff).Reply()
	if err != nil {
		return nil, err
	}

	data := xImg.Data
	for i := 0; i < len(data); i += 4 {
		data[i], data[i+2], data[i+3] = data[i+2], data[i], 255
	}

	return &image.RGBA{
		Pix:    data,
		Stride: 4 * w,
		Rect:   image.Rect(0, 0, w, h),
	}, nil
}

func (s screen) close() {
	s.conn.Close()
	s.conn = nil
}
