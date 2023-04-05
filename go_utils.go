package go_gpedit

import (
	"fmt"

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
