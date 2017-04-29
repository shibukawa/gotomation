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

// build! darwin

package gotomation

/*
#cgo LDFLAGS: -framework Carbon
#include <ApplicationServices/ApplicationServices.h>
*/
import "C"

import (
	"errors"
	"fmt"
	"image"
	"time"
	"unsafe"
)

type screen struct {
	id   int
	w, h int
}

func (s screen) getID() int {
	return s.id
}

func GetMainScreen() (*Screen, error) {
	return GetScreen(int(C.CGMainDisplayID()))
}

func GetScreen(id int) (*Screen, error) {
	var w, h C.size_t
	w = C.CGDisplayPixelsWide(C.CGDirectDisplayID(id))
	h = C.CGDisplayPixelsHigh(C.CGDirectDisplayID(id))
	return &Screen{
		screen: &screen{
			id: id,
			w:  int(w),
			h:  int(h),
		},
		mouse: &mouse{},
		keyboard: &keyboard{
			waitBetweenChars: 50 * time.Millisecond,
		},
	}, nil
}

func (s screen) capture(rect image.Rectangle) (image.Image, error) {
	osImage := C.CGDisplayCreateImageForRect(C.CGDirectDisplayID(s.id),
		C.CGRectMake(C.CGFloat(rect.Min.X), C.CGFloat(rect.Min.Y), C.CGFloat(rect.Dx()), C.CGFloat(rect.Dy())))

	if C.uintptr_t(uintptr(unsafe.Pointer(osImage))) == 0 {
		return nil, errors.New("@1")
	}
	defer C.CGImageRelease(osImage)

	imageData := C.CGDataProviderCopyData(C.CGImageGetDataProvider(osImage))
	defer C.CFRelease(C.CFTypeRef(imageData))

	if C.uintptr_t(uintptr(unsafe.Pointer(imageData))) == 0 {
		return nil, errors.New("@2")
	}

	bufferSize := C.CFDataGetLength(imageData)
	buffer := make([]uint8, int(bufferSize))
	bufferPtr := uintptr(unsafe.Pointer(&buffer[0]))
	C.CFDataGetBytes(imageData, C.CFRangeMake(0, bufferSize), (*C.UInt8)(unsafe.Pointer(bufferPtr)))

	imageW := int(C.CGImageGetWidth(osImage))
	imageH := int(C.CGImageGetHeight(osImage))
	bytesPerRow := int(C.CGImageGetBytesPerRow(osImage))
	bitsPerPixel := int(C.CGImageGetBitsPerPixel(osImage))
	if bitsPerPixel == 32 {
		// macOS returns image with BGRA pixels
		for y := 0; y < imageH; y++ {
			for x := 0; x < imageW; x++ {
				offset := bytesPerRow*y + 4*x
				buffer[offset], buffer[offset+2] = buffer[offset+2], buffer[offset]
			}
		}
		return &image.RGBA{
			Pix:    buffer,
			Stride: bytesPerRow,
			Rect:   image.Rect(0, 0, imageW, imageH),
		}, nil
	}
	return nil, fmt.Errorf("Capture doesn't support color mode with %d bits per pixel", bitsPerPixel)
}

func (s screen) close() {
}
