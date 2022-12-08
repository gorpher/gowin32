package gowin32

import (
	"github.com/gorpher/gowin32/wrappers"
	"syscall"
	"unsafe"
)

func NetUserEnum() []wrappers.UserRecord {
	level := uint32(3)
	entriesread := uint32(0)
	totalentries := uint32(0)
	resumeHandle := uint32(0)
	var buffer uintptr
	var result []wrappers.UserRecord
	defer wrappers.NetApiBufferFree(buffer)
	for {
		res := wrappers.NetUserEnum(wrappers.Lpcwstr(""), level,
			wrappers.FILTER_NORMAL_ACCOUNT|wrappers.FILTER_WORKSTATION_TRUST_ACCOUNT,
			&buffer,
			uint32(0xFFFFFFFF),
			&entriesread, &totalentries, &resumeHandle)
		if res == 0 {
			pos := buffer
			for i := uint32(0); i < entriesread; i++ {
				encodedUserRecord := (*wrappers.USER_INFO_3)(unsafe.Pointer(pos))
				if encodedUserRecord == nil {
					break
				}
				result = append(result, encodedUserRecord.UserRecord())
				pos = pos + unsafe.Sizeof(*encodedUserRecord)
			}
		}
		if res != wrappers.NET_API_STATUS(wrappers.ERROR_MORE_DATA) {
			break
		}
	}
	return result
}

func NetGroupGetUsers(groupName string) []wrappers.GroupUserInfo {
	level := uint32(0)

	entriesread := uint32(0)
	totalentries := uint32(0)
	resumeHandle := uint32(0)
	var buffer uintptr
	var result []wrappers.GroupUserInfo
	defer wrappers.NetApiBufferFree(buffer)
	for {
		res := wrappers.NetGroupGetUsers(wrappers.Lpcwstr(""), wrappers.Lpcwstr(groupName), level, &buffer, uint32(0xFFFFFFFF),
			&entriesread, &totalentries, &resumeHandle)
		if res != wrappers.NET_API_STATUS(wrappers.ERROR_MORE_DATA) {
			break
		}
		if res == 0 {
			pos := buffer
			for i := uint32(0); i < entriesread; i++ {
				encodedUserRecord := (*wrappers.GROUP_USERS_INFO_0)(unsafe.Pointer(pos))
				if encodedUserRecord == nil {
					break
				}
				result = append(result, encodedUserRecord.GroupUserInfo())
				pos = pos + unsafe.Sizeof(*encodedUserRecord)
			}
		}

	}
	return result
}

func NetGroupEnum() []wrappers.GroupInfo1 {
	level := uint32(0)
	entriesread := uint32(0)
	totalentries := uint32(0)
	resumeHandle := uint64(0)
	var buffer uintptr
	var result []wrappers.GroupInfo1
	defer wrappers.NetApiBufferFree(buffer)
	for {
		res := wrappers.NetGroupEnum(syscall.StringToUTF16Ptr(""), level,
			&buffer,
			uint32(65536),
			&entriesread, &totalentries, &resumeHandle)
		if syscall.Errno(res) == wrappers.ERROR_SUCCESS || syscall.Errno(res) == wrappers.ERROR_MORE_DATA {
			pos := buffer
			for i := uint32(0); i < entriesread; i++ {
				encodedUserRecord := (*wrappers.GROUP_INFO_1)(unsafe.Pointer(pos))
				if encodedUserRecord == nil {
					break
				}
				result = append(result, encodedUserRecord.GroupInfo1())
				pos = pos + unsafe.Sizeof(*encodedUserRecord)
			}
		}
		if res != wrappers.NET_API_STATUS(wrappers.ERROR_MORE_DATA) {
			break
		}
	}
	return result
}

func NetQueryDisplayUserInformation() []wrappers.NetDisplayUser {
	level := uint32(1)
	index := uint32(0)
	entriesread := uint32(65536)
	totalentries := uint32(0)
	var buffer uintptr
	var result []wrappers.NetDisplayUser
	defer wrappers.NetApiBufferFree(buffer)
	for {
		res := wrappers.NetQueryDisplayInformation(syscall.StringToUTF16Ptr(""), level, index,
			entriesread,
			uint32(0xFFFFFFFF),
			&totalentries, &buffer)
		if syscall.Errno(res) == wrappers.ERROR_SUCCESS || syscall.Errno(res) == wrappers.ERROR_MORE_DATA {
			pos := buffer
			for i := uint32(0); i < entriesread; i++ {
				encodedUserRecord := (*wrappers.NET_DISPLAY_USER)(unsafe.Pointer(pos))
				if encodedUserRecord == nil {
					break
				}
				index = encodedUserRecord.NextIndex()
				if index == 0 {
					break
				}
				result = append(result, encodedUserRecord.NetDisplayUser())
				pos = pos + unsafe.Sizeof(*encodedUserRecord)
			}
		}
		if syscall.Errno(res) != wrappers.ERROR_MORE_DATA {
			break
		}
	}
	return result
}

func NetQueryDisplayMachineInformation() []wrappers.NetDisplayMachine {
	level := uint32(2)
	index := uint32(0)
	entriesread := uint32(65536)
	totalentries := uint32(0)
	var buffer uintptr
	var result []wrappers.NetDisplayMachine
	defer wrappers.NetApiBufferFree(buffer)
	for {
		res := wrappers.NetQueryDisplayInformation(syscall.StringToUTF16Ptr(""), level, index,
			entriesread,
			uint32(0xFFFFFFFF),
			&totalentries, &buffer)
		if syscall.Errno(res) == wrappers.ERROR_SUCCESS || syscall.Errno(res) == wrappers.ERROR_MORE_DATA {
			pos := buffer
			for i := uint32(0); i < entriesread; i++ {
				encodedUserRecord := (*wrappers.NET_DISPLAY_MACHINE)(unsafe.Pointer(pos))
				if encodedUserRecord == nil {
					break
				}
				index = encodedUserRecord.NextIndex()
				if index == 0 {
					break
				}
				result = append(result, encodedUserRecord.NetDisplayMachine())
				pos = pos + unsafe.Sizeof(*encodedUserRecord)
			}
		}
		if syscall.Errno(res) != wrappers.ERROR_MORE_DATA {
			break
		}
	}
	return result
}

func NetQueryDisplayGroupInformation() []wrappers.NetDisplayGroup {
	level := uint32(3)
	index := uint32(0)
	entriesread := uint32(65536)
	totalentries := uint32(0)
	var buffer uintptr
	var result []wrappers.NetDisplayGroup
	defer wrappers.NetApiBufferFree(buffer)
	for {
		res := wrappers.NetQueryDisplayInformation(syscall.StringToUTF16Ptr(""), level, index,
			entriesread,
			uint32(0xFFFFFFFF),
			&totalentries, &buffer)
		if syscall.Errno(res) == wrappers.ERROR_SUCCESS || syscall.Errno(res) == wrappers.ERROR_MORE_DATA {
			pos := buffer
			for i := uint32(0); i < entriesread; i++ {
				encodedUserRecord := (*wrappers.NET_DISPLAY_GROUP)(unsafe.Pointer(pos))
				if encodedUserRecord == nil {
					break
				}
				index = encodedUserRecord.NextIndex()
				if index == 0 {
					break
				}
				result = append(result, encodedUserRecord.NetDisplayGroup())
				pos = pos + unsafe.Sizeof(*encodedUserRecord)
			}
		}
		if syscall.Errno(res) != wrappers.ERROR_MORE_DATA {
			break
		}
	}
	return result
}
