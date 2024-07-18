package gowin32

import (
	"github.com/gorpher/gowin32/wrappers"
	"strings"
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

func NetGroupAdd(server, name string) error {
	servername := wrappers.Lpcwstr(server)
	info1 := wrappers.GROUP_INFO_0{
		Grpi0Name: wrappers.Lpcwstr(name),
	}
	status := wrappers.NetGroupAdd(servername, 1, (*byte)(unsafe.Pointer(&info1)), nil)
	if status == 0 {
		return nil
	}
	return NewWindowsError("NetGroupAdd ", syscall.Errno(status))
}

func NetGroupDel(server, name string) error {
	status := wrappers.NetGroupDel(wrappers.Lpcwstr(server), wrappers.Lpcwstr(name))
	if status == 0 {
		return nil
	}
	return NewWindowsError("NetGroupDel ", syscall.Errno(status))
}

func NetGroupAddUser(server, groupName, username string) error {
	servername := wrappers.Lpcwstr(server)
	g := wrappers.Lpcwstr(groupName)
	u := wrappers.Lpcwstr(username)

	status := wrappers.NetGroupAddUser(servername, g, u)
	if status == 0 {
		return nil
	}
	return NewWindowsError("NetGroupAddUser ", syscall.Errno(status))
}

func NetGroupDelUser(server, groupName, username string) error {
	servername := wrappers.Lpcwstr(server)
	g := wrappers.Lpcwstr(groupName)
	u := wrappers.Lpcwstr(username)

	status := wrappers.NetGroupDelUser(servername, g, u)
	if status == 0 {
		return nil
	}
	return NewWindowsError("NetGroupDelUser ", syscall.Errno(status))
}

func NetUserAdd(server, username, comment, password string) error {
	servername := wrappers.Lpcwstr(server)
	info1 := wrappers.USER_INFO_1{
		Name:         wrappers.Lpcwstr(username),
		Password:     wrappers.Lpcwstr(password),
		Password_age: 0,
		Priv:         wrappers.USER_PRIV_USER,
		Flags:        wrappers.UF_PASSWD_NOTREQD | wrappers.UF_DONT_EXPIRE_PASSWD | wrappers.UF_NORMAL_ACCOUNT,
		Comment:      wrappers.Lpcwstr(comment),
	}
	status := wrappers.NetUserAdd(servername, 1, (*byte)(unsafe.Pointer(&info1)), nil)
	if status == 0 {
		return nil
	}
	return NewWindowsError("NetUserAdd ", syscall.Errno(status))
}

func SetBit(value uint32, nBitsIdx int, subOrAdd bool) uint32 {
	if subOrAdd {
		value |= 1 << nBitsIdx
	} else {
		value = value & (^(1 << nBitsIdx))
	}
	return value
}

func NetUserActive(server, username string, active bool) error {
	servername := wrappers.Lpcwstr(server)
	uname := wrappers.Lpcwstr(username)
	info1 := wrappers.USER_INFO_1{}
	status := wrappers.NetUserGetInfo(servername, uname, 1, (*byte)(unsafe.Pointer(&info1)))
	if status != 0 {
		return NewWindowsError("NetUserGetInfo ", syscall.Errno(status))
	}
	info1.Flags = wrappers.DWORD(SetBit(uint32(info1.Flags), wrappers.UF_ACCOUNTDISABLE-1, !active))
	status = wrappers.NetUserSetInfo(servername, uname, 1, (*byte)(unsafe.Pointer(&info1)), nil)
	if status == 0 {
		return nil
	}
	return NewWindowsError("NetUserSetInfo ", syscall.Errno(status))
}

func NetUserChangePassword(domainname, username, oldpassword, newpassword string) error {
	status := wrappers.NetUserChangePassword(wrappers.Lpcwstr(domainname), wrappers.Lpcwstr(username), wrappers.Lpcwstr(oldpassword), wrappers.Lpcwstr(newpassword))
	if status != 0 {
		return NewWindowsError("NetUserChangePassword ", syscall.Errno(status))
	}
	return nil
}

func NetUserSetPassword(server, username, password string) error {
	servername := wrappers.Lpcwstr(server)
	uname := wrappers.Lpcwstr(username)
	info1 := wrappers.USER_INFO_1{}
	status := wrappers.NetUserGetInfo(servername, uname, 1, (*byte)(unsafe.Pointer(&info1)))
	if status != 0 {
		return NewWindowsError("NetUserGetInfo ", syscall.Errno(status))
	}
	info1.Password = wrappers.Lpcwstr(password)
	status = wrappers.NetUserSetInfo(servername, uname, 1, (*byte)(unsafe.Pointer(&info1)), nil)
	if status == 0 {
		return nil
	}
	return NewWindowsError("NetUserSetInfo ", syscall.Errno(status))
}

func NetUserDel(server, username string) error {
	servername := wrappers.Lpcwstr(server)
	status := wrappers.NetUserDel(servername, wrappers.Lpcwstr(username))
	if status == 0 {
		return nil
	}
	return NewWindowsError("NetUserDel ", syscall.Errno(status))
}

func NetUserGroups(server, username string) []wrappers.GroupUserInfo {
	level := uint32(0)
	entriesread := uint32(0)
	totalentries := uint32(0)
	var buffer uintptr
	var result []wrappers.GroupUserInfo
	defer wrappers.NetApiBufferFree(buffer)
	for {
		res := wrappers.NetUserGetGroups(wrappers.Lpcwstr(""), wrappers.Lpcwstr(username), level,
			&buffer,
			uint32(0xFFFFFFFF),
			&entriesread, &totalentries)
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
		if res != wrappers.NET_API_STATUS(wrappers.ERROR_MORE_DATA) {
			break
		}
	}
	return result
}

func NetUserSetGroups(server, username string, groups ...string) error {
	entries := uint32(len(groups))
	if entries == 0 {
		return nil
	}
	servername := wrappers.Lpcwstr(server)
	uname := wrappers.Lpcwstr(username)
	gps := make([]wrappers.GROUP_INFO_0, entries)
	for i, group := range groups {
		gps[i] = wrappers.GROUP_INFO_0{
			Grpi0Name: wrappers.Lpcwstr(group),
		}
	}
	status := wrappers.NetUserSetGroups(servername, uname, 0, (*byte)(unsafe.Pointer(&gps)), entries)
	if status == 0 {
		return nil
	}
	return NewWindowsError("NetUserAdd ", syscall.Errno(status))
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

func NetGroupEnum() []wrappers.GroupInfo {
	level := uint32(0)
	entriesread := uint32(0)
	totalentries := uint32(0)
	var buffer uintptr
	var resumeHandle uintptr
	var result []wrappers.GroupInfo
	var serverName string
	defer wrappers.NetApiBufferFree(buffer)
	for {
		res := wrappers.NetGroupEnum(syscall.StringToUTF16Ptr(serverName), level,
			&buffer,
			uint32(65536),
			&entriesread, &totalentries, &resumeHandle)
		if syscall.Errno(res) == wrappers.ERROR_SUCCESS || syscall.Errno(res) == wrappers.ERROR_MORE_DATA {
			pos := buffer
			for i := uint32(0); i < entriesread; i++ {
				encodedUserRecord := (*wrappers.GROUP_INFO_0)(unsafe.Pointer(pos))
				if encodedUserRecord == nil {
					break
				}
				result = append(result, encodedUserRecord.GroupInfo())
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

func NetShareEnum() []wrappers.ShareInfo {
	level := uint32(2)
	entriesread := uint32(0)
	totalentries := uint32(0)
	resumeHandle := uint32(0)
	var buffer uintptr
	var result []wrappers.ShareInfo
	defer wrappers.NetApiBufferFree(buffer)
	for {
		res := wrappers.NetShareEnum("", level,
			&buffer,
			uint32(0xFFFFFFFF),
			&entriesread, &totalentries, &resumeHandle)
		if res == 0 {
			pos := buffer
			for i := uint32(0); i < entriesread; i++ {
				encodedUserRecord := (*wrappers.SHARE_INFO_2)(unsafe.Pointer(pos))
				if encodedUserRecord == nil {
					break
				}
				result = append(result, encodedUserRecord.ShareInfo())
				pos = pos + unsafe.Sizeof(*encodedUserRecord)
			}
		}
		if res != wrappers.NET_API_STATUS(wrappers.ERROR_MORE_DATA) {
			break
		}
	}
	return result
}

func AddNetShare(username, shareDir, shareName string) error {
	var (
		err      error
		adminSid SecurityID
		userSid  SecurityID
	)
	si502 := wrappers.SHARE_INFO_502{
		Netname:      Lpcwstr(shareName),      //共享名
		Type:         wrappers.STYPE_DISKTREE, //资源类型 这里是磁盘
		Path:         Lpcwstr(shareDir),       //文件夹路径
		Permissions:  wrappers.ACCESS_ALL,     //访问权限
		Passwd:       nil,                     //访问密码
		Max_uses:     65536,                   //最大用户连接
		Current_uses: 0,                       //当前连接用户
		Reserved:     0,                       //保留字段

	}

	userSid, _, _, err = GetLocalAccountByName(username)
	if err != nil {
		return err
	}
	var permissions = []PermissionEntry{
		{
			TrusteeType: TrusteeUser,
			Trustee:     userSid,
			Permissions: FileAllAccess,
			AccessMode:  AccessGrant,
		},
	}

	if strings.ToLower(username) != "administrator" {
		adminSid, _, _, err = GetLocalAccountByName("administrator")
		if err != nil {
			return err
		}
		permissions = append(permissions, PermissionEntry{
			TrusteeType: TrusteeUser,
			Trustee:     adminSid,
			Permissions: FileAllAccess,
			AccessMode:  AccessGrant,
		})
	}

	var explicitAccess []wrappers.EXPLICIT_ACCESS
	for _, entry := range permissions {
		explicitAccess = append(explicitAccess, wrappers.EXPLICIT_ACCESS{
			AccessPermissions: uint32(entry.Permissions),
			AccessMode:        int32(entry.AccessMode),
			Inheritance:       wrappers.NO_INHERITANCE,
			Trustee: wrappers.TRUSTEE{
				MultipleTrustee:          nil,
				MultipleTrusteeOperation: wrappers.NO_MULTIPLE_TRUSTEE,
				TrusteeForm:              wrappers.TRUSTEE_IS_SID,
				TrusteeType:              int32(entry.TrusteeType),
				Name:                     (*uint16)(unsafe.Pointer(entry.Trustee.sid)),
			},
		})
	}

	var acl *wrappers.ACL
	if err := wrappers.SetEntriesInAcl(uint32(len(explicitAccess)), &explicitAccess[0], nil, &acl); err != nil {
		return NewWindowsError("SetEntriesInAcl", err)
	}
	defer wrappers.LocalFree(syscall.Handle(unsafe.Pointer(acl)))

	sd := make([]byte, wrappers.SECURITY_DESCRIPTOR_MIN_LENGTH)
	if err := wrappers.InitializeSecurityDescriptor(&sd[0], wrappers.SECURITY_DESCRIPTOR_REVISION); err != nil {
		return NewWindowsError("InitializeSecurityDescriptor", err)
	}
	if err := wrappers.SetSecurityDescriptorDacl(&sd[0], true, acl, false); err != nil {
		return NewWindowsError("SetSecurityDescriptorDacl", err)
	}
	si502.Security_descriptor = &sd[0]
	var paraerr uint16
	return wrappers.NetShareAdd502("", si502, &paraerr)
}

// DelNetShare shareName 共享名称
func DelNetShare(shareName string) error {
	return wrappers.NetShareDel("", shareName, 0)
}
