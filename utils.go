package go_gpedit

import (
	"fmt"
	"unsafe"

	"golang.org/x/sys/windows"
)

// Return nil only if all values are 0 or nil
func WinErrHandler(i ...interface{}) error {
	for _, el := range i {
		switch v := el.(type) {
		case windows.Errno:
			if v == windows.Errno(0) {
				continue
			} else {
				return v
			}
		case uintptr:
			if v == 0 {
				continue
			} else {
				return fmt.Errorf("%d", v)
			}
		default:
			fmt.Printf("WinErrHandler: Unhandled type %T of value %v", v, v)
		}
	}
	return nil
}

var (
	modOle32             = windows.NewLazyDLL("ole32.dll")
	procCoInitialize     = modOle32.NewProc("CoInitialize")
	procCoCreateInstance = modOle32.NewProc("CoCreateInstance")
)

func coInitialize(pvReserved uintptr, dwCoInit uint32) error {
	hr, _, _ := procCoInitialize.Call(
		uintptr(pvReserved),
		uintptr(dwCoInit),
	)
	if hr != 0 {
		return windows.Errno(hr)
	}
	return nil
}

func coCreateInstance(rclsid *windows.GUID, pUnkOuter *windows.GUID, dwClsContext uint32, riid *windows.GUID) (uintptr, error) {
	var pvObj uintptr = 0
	hr, _, _ := procCoCreateInstance.Call(
		uintptr(unsafe.Pointer(rclsid)),
		uintptr(unsafe.Pointer(pUnkOuter)),
		uintptr(dwClsContext),
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(&pvObj)),
	)
	if hr != 0 {
		return pvObj, windows.Errno(hr)
	}
	return pvObj, nil
}
