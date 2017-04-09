/*
Package gotomation is cross-platform system automation library.
*/
package gotomation

import (
	"github.com/bamiaux/rez"
	"image"
)

type Screen struct {
	screen *screen
}

func (s Screen) ID() int {
	return s.screen.id
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
	return s.screen.capture(0, 0, s.screen.w, s.screen.h)
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

func (s Screen) fixRegion(x, y, w, h int) (int, int, int, int) {
	if x < 0 {
		x = 0
	} else if s.screen.w < x {
		x = s.screen.w
	}
	if y < 0 {
		y = 0
	} else if s.screen.h < h {
		y = s.screen.h
	}
	if w < 0 {
		w = 1
	}
	if s.screen.w < x+w {
		w = s.screen.w - x + 1
	}
	if h < 0 {
		h = 1
	}
	if s.screen.h < y+h {
		h = s.screen.h - y + 1
	}
	return x, y, w, h
}

func (s Screen) RawCaptureRegion(x, y, w, h int) (image.Image, error) {
	return s.screen.capture(s.fixRegion(x, y, w, h))
}

func (s Screen) CaptureRegion(x, y, w, h int) (image.Image, error) {
	x, y, w, h = s.fixRegion(x, y, w, h)
	rawImage, err := s.RawCaptureRegion(x, y, w, h)
	if err != nil {
		return nil, err
	}
	size := rawImage.Bounds().Size()
	if size.X == w && size.Y == h {
		return rawImage, err
	}
	result := image.NewRGBA(image.Rect(0, 0, w, h))
	rez.Convert(result, rawImage, rez.NewBilinearFilter())
	return result, nil
}
