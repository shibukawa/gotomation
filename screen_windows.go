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
	"image"
	"time"
	"unsafe"
)

type screen struct {
	hwnd uintptr
	w, h int
}

func (s screen) getID() int {
	return int(s.hwnd)
}

func GetMainScreen() (*Screen, error) {
	hwnd, _, err := procGetDesktopWindow.Call()
	if isError(err) {
		return nil, err
	}
	return GetScreen(hwnd)
}

func GetScreen(hwnd uintptr) (*Screen, error) {
	hdc, _, err := procGetDC.Call(hwnd)
	if isError(err) {
		return nil, err
	}
	defer procReleaseDC.Call(hwnd, hdc)

	displayW, _, err := procGetDeviceCaps.Call(hdc, wHORZRES)
	if isError(err) {
		return nil, err
	}
	displayH, _, err := procGetDeviceCaps.Call(hdc, wVERTRES)
	if isError(err) {
		return nil, err
	}
	return &Screen{
		screen: &screen{
			hwnd: hwnd,
			w:    int(displayW),
			h:    int(displayH),
		},
		mouse: &mouse{},
		keyboard: &keyboard{
			waitBetweenChars: 50 * time.Millisecond,
		},
	}, nil
}

func (s screen) capture(rect image.Rectangle) (image.Image, error) {
	w := rect.Dx()
	h := rect.Dy()
	bi := wBITMAPINFO{}
	bi.bmiHeader.biSize = uint32(unsafe.Sizeof(bi))
	bi.bmiHeader.biWidth = int32(w)
	bi.bmiHeader.biHeight = int32(h)
	bi.bmiHeader.biPlanes = 1
	bi.bmiHeader.biBitCount = 32
	bi.bmiHeader.biCompression = wBI_RGB
	bi.bmiHeader.biSizeImage = uint32(4 * w * h)

	screen, _, err := procGetDC.Call(0)
	if isError(err) {
		return nil, err
	}
	defer procReleaseDC.Call(0, screen)
	var data uintptr
	dib, _, err := procCreateDIBSection.Call(screen, uintptr(unsafe.Pointer(&bi)), wDIB_RGB_COLORS, uintptr(unsafe.Pointer(&data)), 0, 0)
	if isError(err) {
		return nil, err
	}
	defer procDeleteObject.Call(dib)
	screenMem, _, err := procCreateCompatibleDC.Call(screen)
	if isError(err) {
		return nil, err
	}
	defer procDeleteDC.Call(screenMem)
	_, _, err = procSelectObject.Call(screenMem, dib)
	if isError(err) {
		return nil, err
	}
	_, _, err = procBitBlt.Call(screenMem, 0, 0, uintptr(w), uintptr(h), screen, uintptr(rect.Min.X), uintptr(rect.Min.Y), wSRCCOPY)
	if isError(err) {
		return nil, err
	}
	buffer := make([]byte, 4*w*h)
	// R and B are swapped, upside down
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			sourceOffset := (y*w + x) * 4
			destOffset := ((h-y-1)*w + x) * 4
			buffer[destOffset+2] = *(*byte)(unsafe.Pointer(data + uintptr(sourceOffset)))
			buffer[destOffset+1] = *(*byte)(unsafe.Pointer(data + uintptr(sourceOffset+1)))
			buffer[destOffset] = *(*byte)(unsafe.Pointer(data + uintptr(sourceOffset+2)))
			buffer[destOffset+3] = *(*byte)(unsafe.Pointer(data + uintptr(sourceOffset+3)))
		}
	}
	return &image.RGBA{
		Pix:    buffer,
		Stride: w * 4,
		Rect:   image.Rect(0, 0, w, h),
	}, nil
}

func (s screen) close() {
}
