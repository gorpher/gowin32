/*
 * Copyright (c) 2014-2017 MongoDB, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the license is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package wrappers

import (
	"syscall"
	"unsafe"
)

const (
	FO_MOVE   = 0x0001
	FO_COPY   = 0x0002
	FO_DELETE = 0x0003
	FO_RENAME = 0x0004
)

const (
	FOF_MULTIDESTFILES        = 0x0001
	FOF_CONFIRMMOUSE          = 0x0002
	FOF_SILENT                = 0x0004
	FOF_RENAMEONCOLLISION     = 0x0008
	FOF_NOCONFIRMATION        = 0x0010
	FOF_WANTMAPPINGHANDLE     = 0x0020
	FOF_ALLOWUNDO             = 0x0040
	FOF_FILESONLY             = 0x0080
	FOF_SIMPLEPROGRESS        = 0x0100
	FOF_NOCONFIRMMKDIR        = 0x0200
	FOF_NOERRORUI             = 0x0400
	FOF_NOCOPYSECURITYATTRIBS = 0x0800
	FOF_NORECURSION           = 0x1000
	FOF_NO_CONNECTED_ELEMENTS = 0x2000
	FOF_WANTNUKEWARNING       = 0x4000
	FOF_NORECURSEREPARSE      = 0x8000
	FOF_NO_UI                 = FOF_SILENT | FOF_NOCONFIRMATION | FOF_NOERRORUI | FOF_NOCONFIRMMKDIR
)

const (
	SHGFI_SYSICONINDEX      = 0x4000
	SHGFI_ICON              = 0x000000100
	SHGFI_LARGEICON         = 0x000000000
	SHGFI_USEFILEATTRIBUTES = 0x10
	SHIL_JUMBO              = 0x4
	SHIL_EXTRALARGE         = 0x2
)

type SHFILEOPSTRUCT struct {
	Hwnd                 syscall.Handle
	Func                 uint32
	From                 *uint16
	To                   *uint16
	Flags                uint16
	AnyOperationsAborted int32
	NameMappings         *byte
	ProgressTitle        *uint16
}

type SHFILEINFO struct {
	HIcon         syscall.Handle
	IIcon         int32
	DwAttributes  uint32
	SzDisplayName [MAX_PATH]uint16
	SzTypeName    [80]uint16
}

var (
	modshell32 = syscall.NewLazyDLL("shell32.dll")

	procCommandLineToArgvW = modshell32.NewProc("CommandLineToArgvW")
	procSHFileOperationW   = modshell32.NewProc("SHFileOperationW")
	procExtractIconExW     = modshell32.NewProc("ExtractIconExW")
	procSHGetFileInfoW     = modshell32.NewProc("SHGetFileInfoW")
)

func CommandLineToArgvW(cmdLine *uint16, numArgs *int32) (**uint16, error) {
	r1, _, e1 := syscall.Syscall(
		procCommandLineToArgvW.Addr(),
		2,
		uintptr(unsafe.Pointer(cmdLine)),
		uintptr(unsafe.Pointer(numArgs)),
		0)
	if r1 == 0 {
		if e1 != ERROR_SUCCESS {
			return nil, e1
		} else {
			return nil, syscall.EINVAL
		}
	}
	return (**uint16)(unsafe.Pointer(r1)), nil
}

func SHFileOperation(fileOp *SHFILEOPSTRUCT) error {
	r1, _, _ := syscall.Syscall(procSHFileOperationW.Addr(), 1, uintptr(unsafe.Pointer(fileOp)), 0, 0)
	if err := syscall.Errno(r1); err != ERROR_SUCCESS {
		return err
	}
	return nil
}

func ExtractIconExW(lpszFile *uint16, nIconIndex int, phiconLarge *syscall.Handle, phiconSmall *syscall.Handle, nIcons int) error {
	r1, _, e1 := syscall.Syscall6(procExtractIconExW.Addr(),
		5, uintptr(unsafe.Pointer(lpszFile)),
		uintptr(nIconIndex),
		uintptr(unsafe.Pointer(phiconLarge)),
		uintptr(unsafe.Pointer(phiconSmall)),
		uintptr(nIcons), 0)
	if r1 == 0 {
		if e1 != ERROR_SUCCESS {
			return e1
		} else {
			return syscall.EINVAL
		}
	}
	return nil
}

func SHGetFileInfoW(pszPath *uint16, dwFileAttributes uint32, psfi uintptr, cbFileInfo uint32, uFlags uint32) error {
	r1, _, e1 := syscall.Syscall6(procSHGetFileInfoW.Addr(),
		5, uintptr(unsafe.Pointer(pszPath)), uintptr(dwFileAttributes), uintptr(psfi), uintptr(cbFileInfo), uintptr(uFlags), 0)
	if r1 == 0 {
		if e1 != ERROR_SUCCESS {
			return e1
		} else {
			return syscall.EINVAL
		}
	}
	return nil
}
