package wrappers

import (
	"syscall"
	"unsafe"
)

var (
	modnetapi32 = syscall.NewLazyDLL("netapi32.dll")

	procNetApiBufferFree           = modnetapi32.NewProc("NetApiBufferFree")
	procNetUserEnum                = modnetapi32.NewProc("NetUserEnum")
	procNetUserChangePassword      = modnetapi32.NewProc("NetUserChangePassword")
	procNetUserGetInfo             = modnetapi32.NewProc("NetUserGetInfo")
	procNetUserSetInfo             = modnetapi32.NewProc("NetUserSetInfo")
	procNetUserAdd                 = modnetapi32.NewProc("NetUserAdd")
	procNetUserDel                 = modnetapi32.NewProc("NetUserDel")
	procNetGroupEnum               = modnetapi32.NewProc("NetGroupEnum")
	procNetGroupAdd                = modnetapi32.NewProc("NetGroupAdd")
	procNetGroupAddUser            = modnetapi32.NewProc("NetGroupAddUser")
	procNetGroupDel                = modnetapi32.NewProc("NetGroupDel")
	procNetGroupDelUser            = modnetapi32.NewProc("NetGroupDelUser")
	procNetUserSetGroups           = modnetapi32.NewProc("NetUserSetGroups")
	procNetQueryDisplayInformation = modnetapi32.NewProc("NetQueryDisplayInformation")
	procNetUserGetGroups           = modnetapi32.NewProc("NetUserGetGroups")
	procNetGroupGetUsers           = modnetapi32.NewProc("NetGroupGetUsers")
	//sys  NetShareAdd(serverName *uint16, level uint32, buf *byte, parmErr *uint16) (neterr error) = modnetapi32.NetShareAdd
	//sys  NetShareDel(serverName *uint16, netName *uint16, reserved uint32) (neterr error) = modnetapi32.NetShareDel
	procNetShareAdd  = modnetapi32.NewProc("NetShareAdd")
	procNetShareDel  = modnetapi32.NewProc("NetShareDel")
	procNetShareEnum = modnetapi32.NewProc("NetShareEnum")
)

const (
	ACCESS_NONE   = 0
	ACCESS_READ   = 0x01
	ACCESS_WRITE  = 0x02
	ACCESS_CREATE = 0x04
	ACCESS_EXEC   = 0x08
	ACCESS_DELETE = 0x10
	ACCESS_ATRIB  = 0x20
	ACCESS_PERM   = 0x40

	ACCESS_GROUP = 0x8000

	STYPE_DISKTREE = 0x00
	ACCESS_ALL     = ACCESS_READ | ACCESS_WRITE | ACCESS_CREATE | ACCESS_EXEC | ACCESS_DELETE | ACCESS_ATRIB | ACCESS_PERM
)

type SHARE_INFO_2 struct {
	Netname     *uint16
	Type        uint32
	Remark      *uint16
	Permissions uint32
	MaxUses     uint32
	CurrentUses uint32
	Path        *uint16
	Passwd      *uint16
}

func (a SHARE_INFO_2) ShareInfo() ShareInfo {
	return ShareInfo{
		Netname:     LpstrToString(a.Netname),
		Remark:      LpstrToString(a.Remark),
		Path:        LpstrToString(a.Path),
		Passwd:      LpstrToString(a.Passwd),
		Type:        int(a.Type),
		Permissions: int(a.Permissions),
		MaxUses:     int(a.MaxUses),
		CurrentUses: int(a.CurrentUses),
	}
}

type ShareInfo struct {
	Netname     string
	Remark      string
	Path        string
	Passwd      string
	Type        int
	Permissions int
	MaxUses     int
	CurrentUses int
}
type SHARE_INFO_502 struct {
	Netname             *uint16
	Type                uint32
	Remark              *uint16
	Permissions         uint32
	Max_uses            uint32
	Current_uses        uint32
	Path                *uint16
	Passwd              *uint16
	Reserved            uint32
	Security_descriptor *byte
}

type SHARE_INFO_503 struct {
	Netname             *uint16
	Type                uint32
	Remark              *uint16
	Permissions         uint32
	Max_uses            uint32
	Current_uses        uint32
	Path                *uint16
	Passwd              *uint16
	Servername          *uint16
	Reserved            uint32
	Security_descriptor *byte
}
type USER_INFO_1 struct {
	//	LPWSTR usri1_name;
	//  LPWSTR usri1_password;
	//  DWORD  usri1_password_age;
	//  DWORD  usri1_priv;
	//  LPWSTR usri1_home_dir;
	//  LPWSTR usri1_comment;
	//  DWORD  usri1_flags;
	//  LPWSTR usri1_script_path;
	Name         LPWSTR
	Password     LPWSTR
	Password_age DWORD
	Priv         DWORD
	Home_dir     LPWSTR
	Comment      LPWSTR
	Flags        DWORD
	Script_path  LPWSTR
}
type USER_INFO_3 struct {
	Name             LPWSTR
	Password         LPWSTR
	Password_age     DWORD
	Priv             DWORD
	Home_dir         LPWSTR
	Comment          LPWSTR
	Flags            DWORD
	Script_path      LPWSTR
	Auth_flags       DWORD
	Full_name        LPWSTR
	Usr_comment      LPWSTR
	Parms            LPWSTR
	Workstations     LPWSTR
	Last_logon       DWORD
	Last_logoff      DWORD
	Acct_expires     DWORD
	Max_storage      DWORD
	Units_per_week   DWORD
	Logon_hours      PBYTE
	Bad_pw_count     DWORD
	Num_logons       DWORD
	Logon_server     LPWSTR
	Country_code     DWORD
	Code_page        DWORD
	User_id          DWORD
	Primary_group_id DWORD
	Profile          LPWSTR
	Home_dir_drive   LPWSTR
	Password_expired DWORD
}

type UserRecord struct {
	Name            string
	Password        string
	PasswordAge     int
	Priv            int
	HomeDir         string
	Comment         string
	Flags           int
	ScriptPath      string
	AuthFlags       int
	FullName        string
	UserComment     string
	Params          string
	Workstations    string
	LastLogon       int
	LastLogoff      int
	AcctExpires     int
	MaxStorage      int
	UnitsPerWeek    int
	BadPwCount      int
	NumLogons       int
	LogonServer     string
	CountryCode     int
	CodePage        int
	UserId          int
	UserSid         string
	PrimaryGroupId  int
	Profile         string
	HomeDirDrive    string
	PasswordExpired int
}

func (a USER_INFO_3) UserRecord() UserRecord {
	name := LpstrToString(a.Name)
	sid, _, _, _ := syscall.LookupSID("", name)
	sidString, _ := sid.String()
	return UserRecord{
		Name:            name,
		Password:        LpstrToString(a.Password),
		PasswordAge:     int(a.Password_age),
		Priv:            int(a.Priv),
		HomeDir:         LpstrToString(a.Home_dir),
		Comment:         LpstrToString(a.Comment),
		Flags:           int(a.Flags),
		ScriptPath:      LpstrToString(a.Script_path),
		AuthFlags:       int(a.Auth_flags),
		FullName:        LpstrToString(a.Full_name),
		UserComment:     LpstrToString(a.Usr_comment),
		Params:          LpstrToString(a.Parms),
		Workstations:    LpstrToString(a.Workstations),
		LastLogon:       int(a.Last_logon),
		LastLogoff:      int(a.Last_logoff),
		AcctExpires:     int(a.Acct_expires),
		MaxStorage:      int(a.Max_storage),
		UnitsPerWeek:    int(a.Units_per_week),
		BadPwCount:      int(a.Bad_pw_count),
		NumLogons:       int(a.Num_logons),
		LogonServer:     LpstrToString(a.Logon_server),
		CountryCode:     int(a.Country_code),
		CodePage:        int(a.Code_page),
		UserId:          int(a.User_id),
		UserSid:         sidString,
		PrimaryGroupId:  int(a.Primary_group_id),
		Profile:         LpstrToString(a.Profile),
		HomeDirDrive:    LpstrToString(a.Home_dir_drive),
		PasswordExpired: int(a.Password_expired),
	}
}

const (
	// from LMaccess.h

	USER_PRIV_GUEST = 0
	USER_PRIV_USER  = 1
	USER_PRIV_ADMIN = 2

	UF_SCRIPT                          = 0x0001
	UF_ACCOUNTDISABLE                  = 0x0002
	UF_HOMEDIR_REQUIRED                = 0x0008
	UF_LOCKOUT                         = 0x0010
	UF_PASSWD_NOTREQD                  = 0x0020
	UF_PASSWD_CANT_CHANGE              = 0x0040
	UF_ENCRYPTED_TEXT_PASSWORD_ALLOWED = 0x0080

	UF_TEMP_DUPLICATE_ACCOUNT    = 0x0100
	UF_NORMAL_ACCOUNT            = 0x0200
	UF_INTERDOMAIN_TRUST_ACCOUNT = 0x0800
	UF_WORKSTATION_TRUST_ACCOUNT = 0x1000
	UF_SERVER_TRUST_ACCOUNT      = 0x2000

	UF_ACCOUNT_TYPE_MASK = UF_TEMP_DUPLICATE_ACCOUNT |
		UF_NORMAL_ACCOUNT |
		UF_INTERDOMAIN_TRUST_ACCOUNT |
		UF_WORKSTATION_TRUST_ACCOUNT |
		UF_SERVER_TRUST_ACCOUNT

	UF_DONT_EXPIRE_PASSWD                     = 0x10000
	UF_MNS_LOGON_ACCOUNT                      = 0x20000
	UF_SMARTCARD_REQUIRED                     = 0x40000
	UF_TRUSTED_FOR_DELEGATION                 = 0x80000
	UF_NOT_DELEGATED                          = 0x100000
	UF_USE_DES_KEY_ONLY                       = 0x200000
	UF_DONT_REQUIRE_PREAUTH                   = 0x400000
	UF_PASSWORD_EXPIRED                       = 0x800000
	UF_TRUSTED_TO_AUTHENTICATE_FOR_DELEGATION = 0x1000000
	UF_NO_AUTH_DATA_REQUIRED                  = 0x2000000
	UF_PARTIAL_SECRETS_ACCOUNT                = 0x4000000
	UF_USE_AES_KEYS                           = 0x8000000

	UF_SETTABLE_BITS = UF_SCRIPT |
		UF_ACCOUNTDISABLE |
		UF_LOCKOUT |
		UF_HOMEDIR_REQUIRED |
		UF_PASSWD_NOTREQD |
		UF_PASSWD_CANT_CHANGE |
		UF_ACCOUNT_TYPE_MASK |
		UF_DONT_EXPIRE_PASSWD |
		UF_MNS_LOGON_ACCOUNT |
		UF_ENCRYPTED_TEXT_PASSWORD_ALLOWED |
		UF_SMARTCARD_REQUIRED |
		UF_TRUSTED_FOR_DELEGATION |
		UF_NOT_DELEGATED |
		UF_USE_DES_KEY_ONLY |
		UF_DONT_REQUIRE_PREAUTH |
		UF_PASSWORD_EXPIRED |
		UF_TRUSTED_TO_AUTHENTICATE_FOR_DELEGATION |
		UF_NO_AUTH_DATA_REQUIRED |
		UF_USE_AES_KEYS |
		UF_PARTIAL_SECRETS_ACCOUNT

	FILTER_TEMP_DUPLICATE_ACCOUNT    = uint32(0x0001)
	FILTER_NORMAL_ACCOUNT            = uint32(0x0002)
	FILTER_INTERDOMAIN_TRUST_ACCOUNT = uint32(0x0008)
	FILTER_WORKSTATION_TRUST_ACCOUNT = uint32(0x0010)
	FILTER_SERVER_TRUST_ACCOUNT      = uint32(0x0020)

	LG_INCLUDE_INDIRECT = 0x0001

	// Memory protection constants
	PAGE_EXECUTE           = 0x10
	PAGE_EXECUTE_READ      = 0x20
	PAGE_EXECUTE_READWRITE = 0x40
	PAGE_EXECUTE_WRITECOPY = 0x80
	PAGE_NOACCESS          = 0x1
	PAGE_READONLY          = 0x2
	PAGE_READWRITE         = 0x4
	PAGE_WRITECOPY         = 0x8

	// NtQuerySystemInformation
	SystemHandleInformation = 0x10
	SystemObjectInformation = 0x11

	// NtQueryObject
	ObjectBasicInformation = 0x0
	ObjectNameInformation  = 0x1
	ObjectTypeInformation  = 0x2

	// NtQueryInformationProcess
	ProcessBasicInformation       = 0x0
	ProcessImageFileName          = 27
	ProcessCommandLineInformation = 60

	// NtQueryInformationThread
	ThreadBasicInformation   = 0
	ThreadImpersonationToken = 5

	//PROCESS_QUERY_LIMITED_INFORMATION = 0x1000
	THREAD_QUERY_LIMITED_INFORMATION = 0x0800

	// NtOpenDirectoryObject
	DIRECTORY_QUERY    = 1
	DIRECTORY_TRAVERSE = 2

	SYMBOLIC_LINK_QUERY = 1
)

type NET_API_STATUS DWORD

func NetApiBufferFree(Buffer uintptr) (status NET_API_STATUS) {
	r0, _, _ := syscall.Syscall(procNetApiBufferFree.Addr(), 1, uintptr(Buffer), 0, 0)
	status = NET_API_STATUS(r0)
	return
}

func NetUserEnum(servername *uint16, level uint32, filter uint32, bufptr *uintptr,
	prefmaxlen uint32, entriesread *uint32, totalentries *uint32, resume_handle *uint32) NET_API_STATUS {
	r0, _, _ := syscall.SyscallN(procNetUserEnum.Addr(), uintptr(unsafe.Pointer(servername)),
		uintptr(level), uintptr(filter), uintptr(unsafe.Pointer(bufptr)), uintptr(prefmaxlen),
		uintptr(unsafe.Pointer(entriesread)), uintptr(unsafe.Pointer(totalentries)),
		uintptr(unsafe.Pointer(resume_handle)))
	return NET_API_STATUS(r0)
}

func NetUserGetGroups(servername *uint16, username *uint16, level uint32, bufptr *uintptr, prefmaxlen uint32,
	entriesread *uint32, totalentries *uint32) (status NET_API_STATUS) {
	r0, _, _ := syscall.SyscallN(procNetUserGetGroups.Addr(),
		uintptr(unsafe.Pointer(servername)), uintptr(unsafe.Pointer(username)),
		uintptr(level), uintptr(unsafe.Pointer(bufptr)), uintptr(prefmaxlen),
		uintptr(unsafe.Pointer(entriesread)), uintptr(unsafe.Pointer(totalentries)))
	status = NET_API_STATUS(r0)
	return
}

//GROUP_INFO_0
//typedef struct _GROUP_INFO_0 {
//  LPWSTR grpi0_name;
//} GROUP_INFO_0, *PGROUP_INFO_0, *LPGROUP_INFO_0;

type GROUP_INFO_0 struct {
	Grpi0Name LPWSTR
}

type GroupInfo struct {
	GroupName string
}

func (g GROUP_INFO_0) GroupInfo() GroupInfo {
	return GroupInfo{
		GroupName: LpstrToString(g.Grpi0Name),
	}
}

// NetGroupEnum
// NET_API_STATUS NET_API_FUNCTION NetGroupEnum(
//
//	[in]      LPCWSTR    servername,
//	[in]      DWORD      level,
//	[out]     LPBYTE     *bufptr,
//	[in]      DWORD      prefmaxlen,
//	[out]     LPDWORD    entriesread,
//	[out]     LPDWORD    totalentries,
//	[in, out] PDWORD_PTR resume_handle
//
// );
func NetGroupEnum(servername *uint16, level uint32, bufptr *uintptr, prefmaxlen uint32,
	entriesread LPDWORD, totalentries LPDWORD, resume_handle *uintptr) NET_API_STATUS {
	r1, _, _ := syscall.SyscallN(procNetGroupEnum.Addr(),
		uintptr(unsafe.Pointer(servername)), uintptr(level), uintptr(unsafe.Pointer(bufptr)),
		uintptr(prefmaxlen), uintptr(unsafe.Pointer(entriesread)), uintptr(unsafe.Pointer(totalentries)),
		uintptr(unsafe.Pointer(resume_handle)))
	return NET_API_STATUS(r1)
}

// NetGroupAdd
// NET_API_STATUS NET_API_FUNCTION NetGroupAdd(
//
//	[in]  LPCWSTR servername,
//	[in]  DWORD   level,
//	[in]  LPBYTE  buf,
//	[out] LPDWORD parm_err
//
// );
func NetGroupAdd(servername *uint16, level uint32, bufptr *byte, parmErr *LPDWORD) NET_API_STATUS {
	r1, _, _ := syscall.SyscallN(procNetGroupAdd.Addr(),
		uintptr(unsafe.Pointer(servername)), uintptr(level), uintptr(unsafe.Pointer(bufptr)), uintptr(unsafe.Pointer(parmErr)))
	return NET_API_STATUS(r1)
}

// NetUserChangePassword
// NET_API_STATUS NET_API_FUNCTION NetUserChangePassword(
//
//	[in] LPCWSTR domainname,
//	[in] LPCWSTR username,
//	[in] LPCWSTR oldpassword,
//	[in] LPCWSTR newpassword
//
// );
func NetUserChangePassword(domainname, username, oldpassword, newpassword *uint16) NET_API_STATUS {
	r1, _, _ := syscall.SyscallN(procNetGroupAdd.Addr(),
		uintptr(unsafe.Pointer(domainname)), uintptr(unsafe.Pointer(username)), uintptr(unsafe.Pointer(oldpassword)), uintptr(unsafe.Pointer(newpassword)))
	return NET_API_STATUS(r1)
}

// NetUserGetInfo
// NET_API_STATUS NET_API_FUNCTION NetUserGetInfo(
//
//	[in]  LPCWSTR servername,
//	[in]  LPCWSTR username,
//	[in]  DWORD   level,
//	[out] LPBYTE  *bufptr
//
// );
func NetUserGetInfo(servername *uint16, username *uint16, level uint32, bufptr *byte) NET_API_STATUS {
	r1, _, _ := syscall.SyscallN(procNetUserGetInfo.Addr(),
		uintptr(unsafe.Pointer(servername)), uintptr(unsafe.Pointer(username)), uintptr(level), uintptr(unsafe.Pointer(bufptr)))
	return NET_API_STATUS(r1)
}

// NetUserSetInfo
// NET_API_STATUS NET_API_FUNCTION NetUserSetInfo(
//
//	[in]  LPCWSTR servername,
//	[in]  LPCWSTR username,
//	[in]  DWORD   level,
//	[in]  LPBYTE  buf,
//	[out] LPDWORD parm_err
//
// );
func NetUserSetInfo(servername *uint16, username *uint16, level uint32, bufptr *byte, parmErr *LPDWORD) NET_API_STATUS {
	r1, _, _ := syscall.SyscallN(procNetUserSetInfo.Addr(),
		uintptr(unsafe.Pointer(servername)), uintptr(unsafe.Pointer(username)), uintptr(level), uintptr(unsafe.Pointer(bufptr)), uintptr(unsafe.Pointer(parmErr)))
	return NET_API_STATUS(r1)
}

// NetGroupAddUser
// NET_API_STATUS NET_API_FUNCTION NetGroupAddUser(
//
//	[in] LPCWSTR servername,
//	[in] LPCWSTR GroupName,
//	[in] LPCWSTR username
//
// );
func NetGroupAddUser(servername, groupName, username *uint16) NET_API_STATUS {
	r1, _, _ := syscall.SyscallN(procNetGroupAddUser.Addr(),
		uintptr(unsafe.Pointer(servername)), uintptr(unsafe.Pointer(groupName)), uintptr(unsafe.Pointer(username)))
	return NET_API_STATUS(r1)
}

// NetGroupDelUser
// NET_API_STATUS NET_API_FUNCTION NetGroupDelUser(
//
//	[in] LPCWSTR servername,
//	[in] LPCWSTR GroupName,
//	[in] LPCWSTR Username
//
// );
func NetGroupDelUser(servername, groupName, username *uint16) NET_API_STATUS {
	r1, _, _ := syscall.SyscallN(procNetGroupDelUser.Addr(),
		uintptr(unsafe.Pointer(servername)), uintptr(unsafe.Pointer(groupName)), uintptr(unsafe.Pointer(username)))
	return NET_API_STATUS(r1)
}

// NetGroupDel
// NET_API_STATUS NET_API_FUNCTION NetGroupDel(
//
//	[in] LPCWSTR servername,
//	[in] LPCWSTR groupname
//
// );
func NetGroupDel(servername *uint16, groupname *uint16) NET_API_STATUS {
	r1, _, _ := syscall.SyscallN(procNetGroupDel.Addr(),
		uintptr(unsafe.Pointer(servername)), uintptr(unsafe.Pointer(groupname)))
	return NET_API_STATUS(r1)
}

// NetUserAdd
// NET_API_STATUS NET_API_FUNCTION NetUserAdd(
//
//	[in]  LPCWSTR servername,
//	[in]  DWORD   level,
//	[in]  LPBYTE  buf,
//	[out] LPDWORD parm_err
//
// );
func NetUserAdd(servername *uint16, level uint32, bufptr *byte, parmErr *LPDWORD) NET_API_STATUS {
	r1, _, _ := syscall.SyscallN(procNetUserAdd.Addr(),
		uintptr(unsafe.Pointer(servername)), uintptr(level), uintptr(unsafe.Pointer(bufptr)), uintptr(unsafe.Pointer(parmErr)))
	return NET_API_STATUS(r1)
}

func NetUserDel(servername *uint16, username *uint16) NET_API_STATUS {
	r1, _, _ := syscall.SyscallN(procNetUserDel.Addr(),
		uintptr(unsafe.Pointer(servername)), uintptr(unsafe.Pointer(username)))
	return NET_API_STATUS(r1)
}

// NetUserSetGroups
// NET_API_STATUS NET_API_FUNCTION NetUserSetGroups(
//
//	[in] LPCWSTR servername,
//	[in] LPCWSTR username,
//	[in] DWORD   level,
//	[in] LPBYTE  buf,
//	[in] DWORD   num_entries
//
// );
func NetUserSetGroups(servername *uint16, username *uint16, level uint32, bufptr *byte, entries uint32) NET_API_STATUS {
	r1, _, _ := syscall.SyscallN(procNetUserSetGroups.Addr(),
		uintptr(unsafe.Pointer(servername)), uintptr(unsafe.Pointer(username)), uintptr(level), uintptr(unsafe.Pointer(bufptr)), uintptr(entries))
	return NET_API_STATUS(r1)
}

type GROUP_INFO_1 struct {
	grpi1_name    LPWSTR
	grpi1_comment LPWSTR
}
type GroupInfo1 struct {
	GroupName    string
	GroupComment string
}

func (g GROUP_INFO_1) GroupInfo1() GroupInfo1 {
	return GroupInfo1{
		GroupName:    LpstrToString(g.grpi1_name),
		GroupComment: LpstrToString(g.grpi1_comment),
	}

}

// NetGroupGetUsers
// NET_API_STATUS NET_API_FUNCTION NetGroupGetUsers(
//
//	[in]      LPCWSTR    servername,
//	[in]      LPCWSTR    groupname,
//	[in]      DWORD      level,
//	[out]     LPBYTE     *bufptr,
//	[in]      DWORD      prefmaxlen,
//	[out]     LPDWORD    entriesread,
//	[out]     LPDWORD    totalentries,
//	[in, out] PDWORD_PTR ResumeHandle
//
// );
func NetGroupGetUsers(servername *uint16, groupname *uint16, level uint32, bufptr *uintptr, prefmaxlen uint32,
	entriesread *uint32, totalentries *uint32, resume_handle *uint32) NET_API_STATUS {
	r1, _, _ := procNetGroupGetUsers.Call(uintptr(unsafe.Pointer(servername)),
		uintptr(unsafe.Pointer(groupname)), uintptr(level), uintptr(unsafe.Pointer(bufptr)), uintptr(prefmaxlen),
		uintptr(unsafe.Pointer(entriesread)), uintptr(unsafe.Pointer(totalentries)), uintptr(unsafe.Pointer(resume_handle)))
	return NET_API_STATUS(r1)
}

//typedef struct _GROUP_USERS_INFO_1 {
//LPWSTR grui1_name;
//DWORD  grui1_attributes;
//} GROUP_USERS_INFO_1, *PGROUP_USERS_INFO_1, *LPGROUP_USERS_INFO_1;

type GroupUserInfo struct {
	GroupName string
}

type GROUP_USERS_INFO_0 struct {
	grui0_name LPWSTR
}

func (g GROUP_USERS_INFO_0) GroupUserInfo() GroupUserInfo {
	return GroupUserInfo{
		GroupName: LpstrToString(g.grui0_name),
	}
}

// NetQueryDisplayInformation
// NET_API_STATUS NET_API_FUNCTION NetQueryDisplayInformation(
//
//	[in]  LPCWSTR ServerName,
//	[in]  DWORD   Level,
//	[in]  DWORD   Index,
//	[in]  DWORD   EntriesRequested,
//	[in]  DWORD   PreferredMaximumLength,
//	[out] LPDWORD ReturnedEntryCount,
//	[out] PVOID   *SortedBuffer
//
// );
func NetQueryDisplayInformation(servername *uint16, level uint32, index uint32,
	entriesRequested uint32, preferredMaximumLength uint32, returnedEntryCount *uint32, sortedBuffer *uintptr) NET_API_STATUS {
	r0, _, _ := syscall.SyscallN(procNetQueryDisplayInformation.Addr(),
		uintptr(unsafe.Pointer(servername)),
		uintptr(level), uintptr(index), uintptr(entriesRequested), uintptr(preferredMaximumLength),
		uintptr(unsafe.Pointer(returnedEntryCount)), uintptr(unsafe.Pointer(sortedBuffer)),
	)
	return NET_API_STATUS(r0)
}

type NET_DISPLAY_USER struct {
	usri1_name       LPWSTR
	usri1_comment    LPWSTR
	usri1_flags      DWORD
	usri1_full_name  LPWSTR
	usri1_user_id    DWORD
	usri1_next_index DWORD
}

func (d *NET_DISPLAY_USER) NextIndex() uint32 {
	return uint32(d.usri1_next_index)
}

func (d NET_DISPLAY_USER) NetDisplayUser() NetDisplayUser {
	return NetDisplayUser{
		UserName:     LpstrToString(d.usri1_name),
		UserFullName: LpstrToString(d.usri1_full_name),
		UserComment:  LpstrToString(d.usri1_comment),
		UserFlag:     uint32(d.usri1_flags),
		UserId:       uint32(d.usri1_user_id),
	}
}

type NetDisplayUser struct {
	UserName     string
	UserFullName string
	UserComment  string
	UserFlag     uint32
	UserId       uint32
}

type NET_DISPLAY_GROUP struct {
	grpi3_name       LPWSTR
	grpi3_comment    LPWSTR
	grpi3_group_id   DWORD
	grpi3_attributes DWORD
	grpi3_next_index DWORD
}

func (g *NET_DISPLAY_GROUP) NextIndex() uint32 {
	return uint32(g.grpi3_next_index)
}

func (g NET_DISPLAY_GROUP) NetDisplayGroup() NetDisplayGroup {
	return NetDisplayGroup{
		GroupName:       LpstrToString(g.grpi3_name),
		GroupComment:    LpstrToString(g.grpi3_comment),
		GroupId:         uint32(g.grpi3_group_id),
		GroupAttributes: uint32(g.grpi3_attributes),
	}
}

type NetDisplayGroup struct {
	GroupName       string
	GroupComment    string
	GroupId         uint32
	GroupAttributes uint32
}

type NET_DISPLAY_MACHINE struct {
	usri2_name       LPWSTR
	usri2_comment    LPWSTR
	usri2_flags      DWORD
	usri2_user_id    DWORD
	usri2_next_index DWORD
}

func (g *NET_DISPLAY_MACHINE) NextIndex() uint32 {
	return uint32(g.usri2_next_index)
}

func (g NET_DISPLAY_MACHINE) NetDisplayMachine() NetDisplayMachine {
	return NetDisplayMachine{
		MachineName:    LpstrToString(g.usri2_name),
		MachineComment: LpstrToString(g.usri2_name),
		MachineFlags:   uint32(g.usri2_flags),
		MachineUserId:  uint32(g.usri2_user_id),
	}
}

type NetDisplayMachine struct {
	MachineName    string
	MachineComment string
	MachineFlags   uint32
	MachineUserId  uint32
}

// NetShareAdd https://learn.microsoft.com/zh-cn/windows/win32/api/lmshare/nf-lmshare-netshareadd
func NetShareAdd(serverName string, level uint32, buf *byte, parmErr *uint16) (neterr error) {
	sname := Lpcwstr(serverName)
	r0, _, _ := procNetShareAdd.Call(
		uintptr(unsafe.Pointer(sname)),
		uintptr(level),
		uintptr(unsafe.Pointer(buf)),
		uintptr(unsafe.Pointer(parmErr)))
	if r0 != 0 {
		neterr = syscall.Errno(r0)
	}
	return
}

func NetShareAdd2(serverName string, share SHARE_INFO_2, parmErr *uint16) error {
	return NetShareAdd(serverName, 2, (*byte)(unsafe.Pointer(&share)), parmErr)
}

func NetShareAdd502(serverName string, share SHARE_INFO_502, parmErr *uint16) error {
	return NetShareAdd(serverName, 502, (*byte)(unsafe.Pointer(&share)), parmErr)
}

func NetShareAdd503(serverName string, share SHARE_INFO_503, parmErr *uint16) (neterr error) {
	return NetShareAdd(serverName, 503, (*byte)(unsafe.Pointer(&share)), parmErr)
}

// NetShareDel https://learn.microsoft.com/zh-cn/windows/win32/api/lmshare/nf-lmshare-netsharedel
func NetShareDel(serverName string, netName string, reserved uint32) (neterr error) {
	sname := Lpcwstr(serverName)
	nname := Lpcwstr(netName)
	r0, _, _ := procNetShareDel.Call(uintptr(unsafe.Pointer(sname)), uintptr(unsafe.Pointer(nname)), uintptr(reserved))
	if r0 != 0 {
		neterr = syscall.Errno(r0)
	}
	return
}

func NetShareEnum(servername string, level uint32, buf *uintptr,
	prefmaxlen uint32, entriesread *uint32, totalentries *uint32, resume_handle *uint32) NET_API_STATUS {
	sname := Lpcwstr(servername)
	r1, _, _ := procNetShareEnum.Call(
		uintptr(unsafe.Pointer(sname)),
		uintptr(level),
		uintptr(unsafe.Pointer(buf)),
		uintptr(prefmaxlen),
		uintptr(unsafe.Pointer(entriesread)),
		uintptr(unsafe.Pointer(totalentries)),
		uintptr(unsafe.Pointer(resume_handle)))
	return NET_API_STATUS(r1)
}
