package wrappers

import (
	"syscall"
	"unsafe"
)

var (
	modmpr                     = syscall.NewLazyDLL("Mpr.dll")
	procWNetAddConnection2W    = modmpr.NewProc("WNetAddConnection2W")
	procWNetCancelConnection2W = modmpr.NewProc("WNetCancelConnection2W")
)

const (
	RESOURCE_CONNECTED  = 0x00000001
	RESOURCE_GLOBALNET  = 0x00000002
	RESOURCE_REMEMBERED = 0x00000003

	RESOURCE_RECENT  = 0x00000004
	RESOURCE_CONTEXT = 0x00000005

	RESOURCETYPE_ANY   = 0x00000000
	RESOURCETYPE_DISK  = 0x00000001
	RESOURCETYPE_PRINT = 0x00000002

	RESOURCETYPE_RESERVED = 0x00000008
	RESOURCETYPE_UNKNOWN  = 0xFFFFFFFF
)

const (
	CONNECT_UPDATE_PROFILE = 0x00000001
	CONNECT_UPDATE_RECENT  = 0x00000002
	CONNECT_TEMPORARY      = 0x00000004
	CONNECT_INTERACTIVE    = 0x00000008
	CONNECT_PROMPT         = 0x00000010
	CONNECT_NEED_DRIVE     = 0x00000020

	CONNECT_REFCOUNT      = 0x00000040
	CONNECT_REDIRECT      = 0x00000080
	CONNECT_LOCALDRIVE    = 0x00000100
	CONNECT_CURRENT_MEDIA = 0x00000200
	CONNECT_DEFERRED      = 0x00000400
	CONNECT_RESERVED      = 0xFF000000

	CONNECT_COMMANDLINE  = 0x00000800
	CONNECT_CMD_SAVECRED = 0x00001000

	CONNECT_CRED_RESET = 0x00002000
)

type NETRESOURCE struct {
	DwScope       uint32
	DwType        uint32
	DwDisplayType uint32
	DwUsage       uint32
	LpLocalName   *uint16
	LpRemoteName  *uint16
	LpComment     *uint16
	LpProvider    *uint16
}

func WNetAddConnection2W(ns *NETRESOURCE, username, password *uint16, flag uint32) error {
	r1, _, e1 := syscall.Syscall6(procWNetAddConnection2W.Addr(), 4,
		uintptr(unsafe.Pointer(ns)),
		uintptr(unsafe.Pointer(password)),
		uintptr(unsafe.Pointer(username)),
		uintptr(flag),
		0,
		0,
	)
	if r1 != 0 {
		if e1 != ERROR_SUCCESS {
			return e1
		} else {
			return syscall.EINVAL
		}
	}
	return nil
}

func WNetCancelConnection2W(lpName string, flag, force uint32) error {
	r1, _, e1 := syscall.Syscall6(procWNetCancelConnection2W.Addr(), 3,
		uintptr(unsafe.Pointer(Lpcwstr(lpName))),
		uintptr(flag),
		uintptr(force),
		0,
		0,
		0,
	)
	if r1 != 0 {
		if e1 != ERROR_SUCCESS {
			return e1
		} else {
			return syscall.EINVAL
		}
	}
	return nil
}
