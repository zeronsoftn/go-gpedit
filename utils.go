package go_gpedit

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	modOle32             = windows.NewLazyDLL("ole32.dll")
	procCoCreateInstance = modOle32.NewProc("CoCreateInstance")
)

func WinErrHandler(err error, rc uintptr) error {
	if err != nil {
		return err
	}
	if rc != 0 {
		return windows.Errno(rc)
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
