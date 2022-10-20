package wrappers

import (
	"syscall"
	"unsafe"
)

var (
	moduserenv                 = syscall.NewLazyDLL("userenv.dll")
	procCreateEnvironmentBlock = moduserenv.NewProc("CreateEnvironmentBlock")
)

func CreateEnvironmentBlock(envInfo *syscall.Handle, userToken syscall.Handle) error {
	r1, _, e1 := procCreateEnvironmentBlock.Call(uintptr(unsafe.Pointer(envInfo)), uintptr(userToken), 0)
	if r1 == 0 {
		if e1 != ERROR_SUCCESS {
			return e1
		} else {
			return syscall.EINVAL
		}
	}
	return nil
}
