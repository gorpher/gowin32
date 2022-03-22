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

package gowin32

import (
	"context"
	"errors"
	"fmt"
	"github.com/gorpher/gowin32/wrappers"
	"net"
	"syscall"
	"time"
	"unsafe"
)

// WTSConnectState enum type - Go version of WTS_CONNECTSTATE_CLASS
type WTSConnectState uint32

const (
	WTSConnectStateActive       WTSConnectState = wrappers.WTSActive
	WTSConnectStateConnected    WTSConnectState = wrappers.WTSConnected
	WTSConnectStateConnectQuery WTSConnectState = wrappers.WTSConnectQuery
	WTSConnectStateShadow       WTSConnectState = wrappers.WTSShadow
	WTSConnectStateDisconnected WTSConnectState = wrappers.WTSDisconnected
	WTSConnectStateIdle         WTSConnectState = wrappers.WTSIdle
	WTSConnectStateListen       WTSConnectState = wrappers.WTSListen
	WTSConnectStateReset        WTSConnectState = wrappers.WTSReset
	WTSConnectStateDown         WTSConnectState = wrappers.WTSDown
	WTSConnectStateInit         WTSConnectState = wrappers.WTSInit
)

func (w WTSConnectState) String() string {
	switch w {
	case wrappers.WTSActive:
		return "WTSActive"
	case wrappers.WTSConnected:
		return "WTSConnected"
	case wrappers.WTSConnectQuery:
		return "WTSConnectQuery"
	case wrappers.WTSShadow:
		return "WTSShadow"
	case wrappers.WTSDisconnected:
		return "WTSDisconnected"
	case wrappers.WTSIdle:
		return "WTSIdle"
	case wrappers.WTSListen:
		return "WTSListen"
	case wrappers.WTSReset:
		return "WTSReset"
	case wrappers.WTSDown:
		return "WTSDown"
	case wrappers.WTSInit:
		return "WTSInit"
	}
	return ""
}

// WTSClientProtocolType enum type go version of WTSClientProtocolType
type WTSClientProtocolType uint32

const (
	WTSClientProtocolConsoleSession WTSClientProtocolType = 0
	WTSClientProtocolInternal       WTSClientProtocolType = 1
	WTSClientProtocolRDP            WTSClientProtocolType = 2
)

// WTSClientInfo - go version of WTSCLIENT structure
type WTSClientInfo struct {
	ClientName          string
	Domain              string
	UserName            string
	WorkDirectory       string
	InitialProgram      string
	EncryptionLevel     byte
	ClientAddressFamily AddressFamily
	clientAddress       [wrappers.CLIENTADDRESS_LENGTH + 1]uint16
	HRes                uint
	VRes                uint
	ColorDepth          uint
	ClientDirectory     string
	ClientBuildNumber   uint
	ClientHardwareId    uint
	ClientProductId     uint
	OutBufCountHost     uint
	OutBufCountClient   uint
	OutBufLength        uint
	DeviceID            string
}

func (ci *WTSClientInfo) ClientAddressToIP() (net.IP, error) {
	var buf [16]byte
	if ci.ClientAddressFamily == wrappers.AF_INET {
		for i := 0; i < 4; i++ {
			buf[i] = byte(ci.clientAddress[i])
		}
	} else {
		n := 0
		for i := 0; i < 8; i++ {
			buf[n] = byte(ci.clientAddress[i] >> 8 & 0xff)
			buf[n+1] = byte(ci.clientAddress[i])
			n = n + 2
		}
	}
	return clientAddressToIP(uint32(ci.ClientAddressFamily), buf[:])
}

// WTSClientDisplay - go version of WTS_CLIENT_DISPLAY structure
type WTSClientDisplay struct {
	HorizontalResolution uint
	VerticalResolution   uint
	ColorDepth           uint
}

// Info - go version of WTSINFO structure
type WTSInfo struct {
	State                   WTSConnectState
	SessionID               uint
	IncomingBytes           uint
	OutgoingBytes           uint
	IncomingFrames          uint
	OutgoingFrames          uint
	IncomingCompressedBytes uint
	OutgoingCompressedBytes uint
	WinStationName          string
	Domain                  string
	UserName                string
	ConnectTime             time.Time
	DisconnectTime          time.Time
	LastInputTime           time.Time
	LogonTime               time.Time
	CurrentTime             time.Time
}

// WTSSessionInfo - go version of WTS_SESSION_INFO structure
type WTSSessionInfo struct {
	SessionID      uint
	WinStationName string
	State          WTSConnectState
}
type WTSPROCESSInfo struct {
	SessionId         uint
	ProcessId         uint
	ProcessName       string
	NumberOfThreads   int64
	HandleCount       int64
	PagefileUsage     int64 // 虚拟内存
	PeakPagefileUsage int64 // 峰值虚拟内存
	Memory            int64
	PeakMemory        int64 // 峰值内存
	UserTime          int64
	KernelTime        int64
}

type WTSServer struct {
	handle syscall.Handle
}

func OpenWTSServer(serverName string) *WTSServer {
	result := WTSServer{}
	if serverName != "" {
		result.handle = wrappers.WTSOpenServer(syscall.StringToUTF16Ptr(serverName))
	}
	return &result
}

func (wts *WTSServer) Close() {
	if wts.handle != 0 {
		wrappers.WTSCloseServer(wts.handle)
		wts.handle = 0
	}
}

func (wts *WTSServer) EnumerateSessions() ([]WTSSessionInfo, error) {
	var sessionInfo *wrappers.WTS_SESSION_INFO
	var count uint32

	if err := wrappers.WTSEnumerateSessions(wts.handle, 0, 1, &sessionInfo, &count); err != nil {
		return nil, err
	}
	defer wrappers.WTSFreeMemory((*byte)(unsafe.Pointer(sessionInfo)))

	si := sessionInfo
	result := make([]WTSSessionInfo, count)
	for i := uint32(0); i < count; i++ {
		result[i] = WTSSessionInfo{SessionID: uint(si.SessionId),
			WinStationName: LpstrToString(si.WinStationName),
			State:          WTSConnectState(si.State)}
		si = (*wrappers.WTS_SESSION_INFO)(unsafe.Pointer(uintptr(unsafe.Pointer(si)) + unsafe.Sizeof(*si)))
	}
	return result, nil
}

func (wts *WTSServer) LogoffSession(sessionID uint, wait bool) error {
	return wrappers.WTSLogoffSession(wts.handle, uint32(sessionID), wait)
}

func (wts *WTSServer) QuerySessionInitialProgram(sessionID uint) (string, error) {
	return wts.querySessionInformationAsString(sessionID, wrappers.WTSInitialProgram)
}

func (wts *WTSServer) QuerySessionApplicationName(sessionID uint) (string, error) {
	return wts.querySessionInformationAsString(sessionID, wrappers.WTSApplicationName)
}

func (wts *WTSServer) QuerySessionWorkingDirectory(sessionID uint) (string, error) {
	return wts.querySessionInformationAsString(sessionID, wrappers.WTSWorkingDirectory)
}

func (wts *WTSServer) QuerySessionID(sessionID uint) (uint, error) {
	r1, err := wts.querySessionInformationAsUint32(sessionID, wrappers.WTSSessionId)
	return uint(r1), err
}

func (wts *WTSServer) QuerySessionUserName(sessionID uint) (string, error) {
	return wts.querySessionInformationAsString(sessionID, wrappers.WTSUserName)
}

func (wts *WTSServer) QuerySessionWinStationName(sessionID uint) (string, error) {
	return wts.querySessionInformationAsString(sessionID, wrappers.WTSWinStationName)
}

func (wts *WTSServer) QuerySessionDomainName(sessionID uint) (string, error) {
	return wts.querySessionInformationAsString(sessionID, wrappers.WTSDomainName)
}

func (wts *WTSServer) QuerySessionConnectState(sessionID uint) (WTSConnectState, error) {
	r1, err := wts.querySessionInformationAsUint32(sessionID, wrappers.WTSConnectState)
	return WTSConnectState(r1), err
}

func (wts *WTSServer) QuerySessionClientBuildNumber(sessionID uint) (uint32, error) {
	return wts.querySessionInformationAsUint32(sessionID, wrappers.WTSClientBuildNumber)
}

func (wts *WTSServer) QuerySessionClientName(sessionID uint) (string, error) {
	return wts.querySessionInformationAsString(sessionID, wrappers.WTSClientName)
}

func (wts *WTSServer) QuerySessionClientDirectory(sessionID uint) (string, error) {
	return wts.querySessionInformationAsString(sessionID, wrappers.WTSClientDirectory)
}

func (wts *WTSServer) QuerySessionClientProductId(sessionID uint) (uint16, error) {
	return wts.querySessionInformationAsUint16(sessionID, wrappers.WTSClientProductId)
}

func (wts *WTSServer) QuerySessionClientHardwareId(sessionID uint) (uint32, error) {
	return wts.querySessionInformationAsUint32(sessionID, wrappers.WTSClientHardwareId)
}

func (wts *WTSServer) QuerySessionClientAddress(sessionID uint) (net.IP, error) {
	var buffer *uint16
	var bytesReturned uint32

	if err := wrappers.WTSQuerySessionInformation(wts.handle, uint32(sessionID), wrappers.WTSClientAddress, &buffer, &bytesReturned); err != nil {
		return net.IP{}, err
	}
	defer wrappers.WTSFreeMemory((*byte)(unsafe.Pointer(buffer)))

	// MS doc: The SmbIP address is offset by two bytes from the start of the Address member of the WTS_CLIENT_ADDRESS structure.
	// https://msdn.microsoft.com/en-us/library/aa383861%28v=vs.85%29.aspx
	a := *(*wrappers.WTS_CLIENT_ADDRESS)(unsafe.Pointer(buffer))
	return clientAddressToIP(a.AddressFamily, a.Address[2:])
}

func (wts *WTSServer) QuerySessionClientDisplay(sessionID uint) (WTSClientDisplay, error) {
	var buffer *uint16
	var bytesReturned uint32

	if err := wrappers.WTSQuerySessionInformation(wts.handle, uint32(sessionID), wrappers.WTSClientDisplay, &buffer, &bytesReturned); err != nil {
		return WTSClientDisplay{}, err
	}
	defer wrappers.WTSFreeMemory((*byte)(unsafe.Pointer(buffer)))

	cd := *(*wrappers.WTS_CLIENT_DISPLAY)(unsafe.Pointer(buffer))
	return WTSClientDisplay{
		HorizontalResolution: uint(cd.HorizontalResolution),
		VerticalResolution:   uint(cd.HorizontalResolution),
		ColorDepth:           uint(cd.ColorDepth)}, nil
}

func (wts *WTSServer) QuerySessionClientProtocolType(sessionID uint) (WTSClientProtocolType, error) {
	r1, err := wts.querySessionInformationAsUint16(sessionID, wrappers.WTSClientProtocolType)
	return WTSClientProtocolType(r1), err
}

func (wts *WTSServer) QuerySessionClientInfo(sessionID uint) (WTSClientInfo, error) {
	var buffer *uint16
	var bytesReturned uint32

	if err := wrappers.WTSQuerySessionInformation(wts.handle, uint32(sessionID), wrappers.WTSClientInfo, &buffer, &bytesReturned); err != nil {
		return WTSClientInfo{}, err
	}
	defer wrappers.WTSFreeMemory((*byte)(unsafe.Pointer(buffer)))

	c := *(*wrappers.WTSCLIENT)(unsafe.Pointer(buffer))
	return WTSClientInfo{
		ClientName:          syscall.UTF16ToString(c.ClientName[:]),
		Domain:              syscall.UTF16ToString(c.Domain[:]),
		UserName:            syscall.UTF16ToString(c.UserName[:]),
		WorkDirectory:       syscall.UTF16ToString(c.WorkDirectory[:]),
		InitialProgram:      syscall.UTF16ToString(c.InitialProgram[:]),
		EncryptionLevel:     c.EncryptionLevel,
		ClientAddressFamily: AddressFamily(c.ClientAddressFamily),
		clientAddress:       c.ClientAddress,
		HRes:                uint(c.HRes),
		VRes:                uint(c.VRes),
		ColorDepth:          uint(c.ColorDepth),
		ClientDirectory:     syscall.UTF16ToString(c.ClientDirectory[:]),
		ClientBuildNumber:   uint(c.ClientBuildNumber),
		ClientHardwareId:    uint(c.ClientHardwareId),
		ClientProductId:     uint(c.ClientProductId),
		OutBufCountHost:     uint(c.OutBufCountHost),
		OutBufCountClient:   uint(c.OutBufCountClient),
		OutBufLength:        uint(c.OutBufLength),
		DeviceID:            syscall.UTF16ToString(c.DeviceId[:]),
	}, nil
}

func (wts *WTSServer) QuerySessionSesionInfo(sessionID uint) (WTSInfo, error) {
	var buffer *uint16
	var bytesReturned uint32

	if err := wrappers.WTSQuerySessionInformation(wts.handle, uint32(sessionID), wrappers.WTSSessionInfo, &buffer, &bytesReturned); err != nil {
		return WTSInfo{}, err
	}
	defer wrappers.WTSFreeMemory((*byte)(unsafe.Pointer(buffer)))

	i := *(*wrappers.WTSINFO)(unsafe.Pointer(buffer))
	return WTSInfo{
		State:                   WTSConnectState(i.State),
		SessionID:               uint(i.SessionId),
		IncomingBytes:           uint(i.IncomingBytes),
		OutgoingBytes:           uint(i.OutgoingBytes),
		IncomingFrames:          uint(i.IncomingFrames),
		OutgoingFrames:          uint(i.OutgoingFrames),
		IncomingCompressedBytes: uint(i.IncomingCompressedBytes),
		OutgoingCompressedBytes: uint(i.OutgoingCompressedBytes),
		WinStationName:          syscall.UTF16ToString(i.WinStationName[:]),
		Domain:                  syscall.UTF16ToString(i.Domain[:]),
		UserName:                syscall.UTF16ToString(i.UserName[:]),
		ConnectTime:             windowsFileTimeToTime(i.ConnectTime),
		DisconnectTime:          windowsFileTimeToTime(i.DisconnectTime),
		LastInputTime:           windowsFileTimeToTime(i.LastInputTime),
		LogonTime:               windowsFileTimeToTime(i.LogonTime),
		CurrentTime:             windowsFileTimeToTime(i.CurrentTime)}, nil
}

func (wts *WTSServer) QuerySessionAddressV4(sessionID uint) (wrappers.WTS_CLIENT_ADDRESS, error) {
	var buffer *uint16
	var bytesReturned uint32

	if err := wrappers.WTSQuerySessionInformation(wts.handle, uint32(sessionID), wrappers.WTSSessionAddressV4, &buffer, &bytesReturned); err != nil {
		return wrappers.WTS_CLIENT_ADDRESS{}, err
	}
	defer wrappers.WTSFreeMemory((*byte)(unsafe.Pointer(buffer)))

	return *(*wrappers.WTS_CLIENT_ADDRESS)(unsafe.Pointer(buffer)), nil
}

func (wts *WTSServer) QuerySessionIsRemoteSession(sessionID uint) (bool, error) {
	return wts.querySessionInformationAsBool(sessionID, wrappers.WTSIsRemoteSession)
}

func (wts *WTSServer) QuerySessionProcessEx(sessionID uint) ([]WTSPROCESSInfo, error) {
	var processInfo *wrappers.WTS_PROCESS_INFO_EX = nil
	var count uint32 = 0
	var level uint32 = 1
	fmt.Printf("sessionID: %d\n", sessionID)
	if err := wrappers.WTSEnumerateProcessesEX(wts.handle, &level, uint32(sessionID), &processInfo, &count); err != nil {
		return nil, err
	}
	defer wrappers.WTSFreeMemory((*byte)(unsafe.Pointer(processInfo)))

	si := processInfo
	result := make([]WTSPROCESSInfo, count)

	for i := uint32(0); i < count; i++ {
		result[i] = WTSPROCESSInfo{
			SessionId:         uint(si.SessionId),
			ProcessId:         uint(si.ProcessId),
			ProcessName:       LpstrToString(si.ProcessName),
			NumberOfThreads:   int64(si.NumberOfThreads),
			HandleCount:       int64(si.HandleCount),
			PagefileUsage:     int64(si.PagefileUsage),
			PeakPagefileUsage: int64(si.PeakPagefileUsage),
			Memory:            int64(si.WorkingSetSize),
			PeakMemory:        int64(si.PeakWorkingSetSize),
			UserTime:          int64(si.UserTime),
			KernelTime:        int64(si.UserTime),
		}
		si = (*wrappers.WTS_PROCESS_INFO_EX)(unsafe.Pointer(uintptr(unsafe.Pointer(si)) + unsafe.Sizeof(*si)))
	}
	return result, nil
}

func (wts *WTSServer) QuerySessionProcess(sessionID uint) ([]WTSPROCESSInfo, error) {
	var processInfo *wrappers.WTS_PROCESS_INFO
	var count uint32
	var level uint32 = 0
	if err := wrappers.WTSEnumerateProcesses(wts.handle, &level, uint32(sessionID), &processInfo, &count); err != nil {
		return nil, err
	}
	defer wrappers.WTSFreeMemory((*byte)(unsafe.Pointer(processInfo)))
	si := processInfo
	result := make([]WTSPROCESSInfo, count)
	for i := uint32(0); i < count; i++ {
		result[i] = WTSPROCESSInfo{SessionId: uint(si.SessionId),
			ProcessId:   uint(si.ProcessId),
			ProcessName: LpstrToString(si.ProcessName),
		}
		si = (*wrappers.WTS_PROCESS_INFO)(unsafe.Pointer(uintptr(unsafe.Pointer(si)) + unsafe.Sizeof(*si)))
	}
	return result, nil
}

func (wts *WTSServer) QueryUserToken(sessionID uint) (*Token, error) {
	var handle syscall.Handle
	if err := wrappers.WTSQueryUserToken(uint32(sessionID), &handle); err != nil {
		return nil, err
	}
	return &Token{handle: handle}, nil
}

func (wts *WTSServer) querySessionInformationAsBool(sessionID uint, infoClass uint32) (bool, error) {
	var buffer *uint16
	var bytesReturned uint32

	if err := wrappers.WTSQuerySessionInformation(wts.handle, uint32(sessionID), infoClass, &buffer, &bytesReturned); err != nil {
		return false, err
	}
	defer wrappers.WTSFreeMemory((*byte)(unsafe.Pointer(buffer)))

	if bytesReturned != 1 {
		return false, bufferSizeError(1, bytesReturned)
	}

	return *(*byte)(unsafe.Pointer(buffer)) != 0, nil
}

func (wts *WTSServer) querySessionInformationAsString(sessionID uint, infoClass uint32) (string, error) {
	var buffer *uint16
	var bytesReturned uint32

	if err := wrappers.WTSQuerySessionInformation(wts.handle, uint32(sessionID), infoClass, &buffer, &bytesReturned); err != nil {
		return "", err
	}
	defer wrappers.WTSFreeMemory((*byte)(unsafe.Pointer(buffer)))

	return LpstrToString(buffer), nil
}

func (wts *WTSServer) querySessionInformationAsUint16(sessionID uint, infoClass uint32) (uint16, error) {
	var buffer *uint16
	var bytesReturned uint32

	if err := wrappers.WTSQuerySessionInformation(wts.handle, uint32(sessionID), infoClass, &buffer, &bytesReturned); err != nil {
		return 0, err
	}
	defer wrappers.WTSFreeMemory((*byte)(unsafe.Pointer(buffer)))

	if bytesReturned != 2 {
		return 0, bufferSizeError(2, bytesReturned)
	}
	return *(*uint16)(unsafe.Pointer(buffer)), nil
}

func (wts *WTSServer) querySessionInformationAsUint32(sessionID uint, infoClass uint32) (uint32, error) {
	var buffer *uint16
	var bytesReturned uint32

	if err := wrappers.WTSQuerySessionInformation(wts.handle, uint32(sessionID), infoClass, &buffer, &bytesReturned); err != nil {
		return 0, err
	}
	defer wrappers.WTSFreeMemory((*byte)(unsafe.Pointer(buffer)))

	if bytesReturned != 4 {
		return 0, bufferSizeError(4, bytesReturned)
	}
	return *(*uint32)(unsafe.Pointer(buffer)), nil
}

func (wts *WTSServer) WTSVirtualChannelWrite(sessionID uint32, channelName string, buf []byte) error {
	virtualChannelHandle := wrappers.WTSVirtualChannelOpenEx(sessionID, channelName, 0)
	defer wrappers.WTSVirtualChannelClose(virtualChannelHandle)
	var byteWrote uint64
	var length = uint64(len(buf))
	return wrappers.WTSVirtualChannelWrite(virtualChannelHandle, &buf[0], length, &byteWrote)
}

func (wts *WTSServer) WTSVirtualChannelRead(sessionID uint32, channelName string, buf []byte, pBytesRead *uint32) error {
	virtualChannelHandle := wrappers.WTSVirtualChannelOpenEx(sessionID, channelName, 0)
	defer wrappers.WTSVirtualChannelClose(virtualChannelHandle)
	return wrappers.WTSVirtualChannelRead(virtualChannelHandle, 6000, &buf[0], uint32(len(buf)), pBytesRead)
}
func (wts *WTSServer) OpenWTSVirtualChannel(ctx context.Context, channelName string, writeChan <-chan []byte, readChan chan<- []byte) error {
	var sessionID uint32
	err := wrappers.ProcessIdToSessionId(wrappers.GetCurrentProcessId(), &sessionID)
	if err != nil {
		return err
	}
	virtualChannelHandle := wrappers.WTSVirtualChannelOpenEx(sessionID, channelName, 0)
	defer wrappers.WTSVirtualChannelClose(virtualChannelHandle)
	buf := make([]byte, 65536)
	var byteReade uint32
	for {
		select {
		case <-ctx.Done():
			return nil
		case writeBody := <-writeChan:
			var byteWrote uint64
			if len(writeBody) > 0 {
				errv := wrappers.WTSVirtualChannelWrite(virtualChannelHandle, &writeBody[0], uint64(len(writeBody)), &byteWrote)
				if errv != nil {
					return errv
				}
			}
		default:
			errv := wrappers.WTSVirtualChannelRead(virtualChannelHandle, 1000, &buf[0], uint32(len(buf)), &byteReade)
			if errv != nil {
				if !errors.Is(errv, wrappers.ERROR_IO_INCOMPLETE) {
					return errv
				}
			}
			if byteReade > 0 {
				respBody := make([]byte, byteReade)
				copy(respBody, buf[:byteReade])
				readChan <- respBody
			}
		}
	}
}

func (wts *WTSServer) WTSVirtualChannelQuery(channelName string, vClass uint32) (string, error) {
	var sessionID uint32
	err := wrappers.ProcessIdToSessionId(wrappers.GetCurrentProcessId(), &sessionID)
	if err != nil {
		return "", err
	}
	virtualChannelHandle := wrappers.WTSVirtualChannelOpenEx(sessionID, channelName, 0)
	var buffer *uint16
	var bytesReturned uint32
	if err := wrappers.WTSVirtualChannelQuery(virtualChannelHandle, vClass, &buffer, &bytesReturned); err != nil {
		return "", err
	}
	defer wrappers.WTSFreeMemory((*byte)(unsafe.Pointer(buffer)))
	defer wrappers.WTSVirtualChannelClose(virtualChannelHandle)
	return LpstrToString(buffer), nil
}

func bufferSizeError(excpected, returned uint32) error {
	return fmt.Errorf("Invalid buffer size. Expected: %d returned: %d", excpected, returned)
}

func clientAddressToIP(addressFamily uint32, address []byte) (net.IP, error) {
	switch addressFamily {
	case wrappers.AF_INET:
		if len(address) >= 4 {
			return net.IPv4(address[0], address[1], address[2], address[3]), nil
		}
	case wrappers.AF_INET6:
		if len(address) >= 16 {
			return net.IP(address[:16]), nil
		}
	}
	return nil, fmt.Errorf("Unknown addressFamily: %v", addressFamily)
}

func windowsFileTimeToTime(fileTime int64) time.Time {
	const TicksPerSecond = 10000000
	const EpochDifference = 11644473600
	// we also can use win32 api FileTimeToSystemTime
	return time.Unix((fileTime/TicksPerSecond)-EpochDifference, 0)
}
