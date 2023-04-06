package go_gpedit

import (
	"unsafe"

	"golang.org/x/sys/windows"
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

var (
	modOle32             = windows.NewLazyDLL("ole32.dll")
	procCoInitialize     = modOle32.NewProc("CoInitialize")
	procCoUninitialize   = modOle32.NewProc("CoUninitialize")
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

func coUninitialize() error {
	hr, _, _ := procCoUninitialize.Call()
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
