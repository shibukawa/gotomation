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
	}, nil
}

func (s screen) capture(x, y, w, h int) (image.Image, error) {
	osImage := C.CGDisplayCreateImageForRect(C.CGDirectDisplayID(s.id),
		C.CGRectMake(C.CGFloat(x), C.CGFloat(y), C.CGFloat(w), C.CGFloat(h)))

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
