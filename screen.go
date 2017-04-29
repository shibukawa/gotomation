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

/*
Package gotomation is cross-platform system automation library.
*/
package gotomation

import (
	"github.com/bamiaux/rez"
	"image"
)

type Screen struct {
	screen   *screen
	mouse    Mouse
	keyboard Keyboard
}

func (s Screen) ID() int {
	return s.screen.getID()
}

func (s Screen) X() int {
	return 0
}

func (s Screen) Y() int {
	return 0
}

func (s Screen) W() int {
	return s.screen.w
}

func (s Screen) H() int {
	return s.screen.h
}

func (s Screen) RawCapture() (image.Image, error) {
	return s.screen.capture(image.Rect(0, 0, s.screen.w, s.screen.h))
}

func (s Screen) Capture() (image.Image, error) {
	rawImage, err := s.RawCapture()
	if err != nil {
		return nil, err
	}
	size := rawImage.Bounds().Size()
	if size.X == s.screen.w && size.Y == s.screen.h {
		return rawImage, err
	}
	result := image.NewRGBA(image.Rect(0, 0, s.screen.w, s.screen.h))
	rez.Convert(result, rawImage, rez.NewBilinearFilter())
	return result, nil
}

func (s Screen) fixRegion(rect image.Rectangle) image.Rectangle {
	if rect.Min.X < 0 {
		rect.Min.X = 0
	}
	if rect.Min.Y < 0 {
		rect.Min.Y = 0
	}
	if s.screen.w < rect.Max.X {
		rect.Max.X = s.screen.w
	}
	if s.screen.h < rect.Max.Y {
		rect.Max.Y = s.screen.h
	}
	return rect
}

func (s Screen) RawCaptureRegion(rect image.Rectangle) (image.Image, error) {
	return s.screen.capture(s.fixRegion(rect))
}

func (s Screen) CaptureRegion(rect image.Rectangle) (image.Image, error) {
	rawImage, err := s.RawCaptureRegion(s.fixRegion(rect))
	if err != nil {
		return nil, err
	}
	size := rawImage.Bounds().Size()
	w := rect.Dx()
	h := rect.Dy()
	if size.X == w && size.Y == h {
		return rawImage, err
	}
	result := image.NewRGBA(image.Rect(0, 0, w, h))
	rez.Convert(result, rawImage, rez.NewBilinearFilter())
	return result, nil
}

func (s Screen) Close() {
	s.screen.close()
}

func (s Screen) Mouse() Mouse {
	return s.mouse
}

func (s Screen) Keyboard() Keyboard {
	return s.keyboard
}
