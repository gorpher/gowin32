package gowin32

import "github.com/gorpher/gowin32/wrappers"

func WNetAddConnection(szRemoteName, szLocalName, username, password string) error {
	ns := wrappers.NETRESOURCE{
		DwType:       wrappers.RESOURCETYPE_ANY,
		LpRemoteName: Lpcwstr(szRemoteName),
		LpLocalName:  Lpcwstr(szLocalName),
	}
	return wrappers.WNetAddConnection2W(&ns, Lpcwstr(username), Lpcwstr(password), wrappers.CONNECT_UPDATE_PROFILE)
}

func WNetCancelConnection2W(lpName string) error {
	return wrappers.WNetCancelConnection2W(lpName, wrappers.CONNECT_UPDATE_PROFILE, 0)
}
