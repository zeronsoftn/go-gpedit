package go_gpedit

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

var (
	modOle32             = windows.NewLazyDLL("ole32.dll")
	procCoCreateInstance = modOle32.NewProc("CoCreateInstance")
)

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
