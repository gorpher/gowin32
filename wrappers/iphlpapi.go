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

var (
	modiphlpapi = syscall.NewLazyDLL("iphlpapi.dll")

	procGetTcpTable          = modiphlpapi.NewProc("GetTcpTable")
	procSendARP              = modiphlpapi.NewProc("SendARP")
	procGetTcpStatistics     = modiphlpapi.NewProc("GetTcpStatistics")
	procGetExtendedTcpTable  = modiphlpapi.NewProc("GetExtendedTcpTable")
	procGetExtendedUdpTable  = modiphlpapi.NewProc("GetExtendedUdpTable")
	procGetBestRoute         = modiphlpapi.NewProc("GetBestRoute")
	procGetIpForwardTable    = modiphlpapi.NewProc("GetIpForwardTable")
	procGetInterfaceInfo     = modiphlpapi.NewProc("GetInterfaceInfo")
	procGetIfTable           = modiphlpapi.NewProc("GetIfTable")
	procDeleteIpForwardEntry = modiphlpapi.NewProc("DeleteIpForwardEntry")
	procCreateIpForwardEntry = modiphlpapi.NewProc("CreateIpForwardEntry")
)

func GetTcpTable(tcpTable *MIB_TCPTABLE, size *uint32, order bool) error {
	r1, _, _ := syscall.Syscall(
		procGetTcpTable.Addr(),
		3,
		uintptr(unsafe.Pointer(tcpTable)),
		uintptr(unsafe.Pointer(size)),
		boolToUintptr(order))
	if err := syscall.Errno(r1); err != ERROR_SUCCESS {
		return err
	}
	return nil
}

func SendARP(destIP, srcIP uint32, macAddr, macAddrLen *uint32) error {
	r1, _, _ := syscall.Syscall6(
		procSendARP.Addr(),
		4,
		uintptr(destIP),
		uintptr(srcIP),
		uintptr(unsafe.Pointer(macAddr)),
		uintptr(unsafe.Pointer(macAddrLen)),
		0,
		0)
	if err := syscall.Errno(r1); err != ERROR_SUCCESS {
		return err
	}
	return nil
}

func GetTcpStatistics(statistics *MIB_TCPSTATS) int {
	ret, _, _ := procGetTcpStatistics.Call(
		uintptr(unsafe.Pointer(statistics)),
	)
	return int(ret)
}

func GetExtendedTcpTable(tcpTable uintptr, size *uint32, order int32, af uint32, tableClass TCP_TABLE_CLASS) int {
	ret, _, _ := procGetExtendedTcpTable.Call(
		tcpTable,
		uintptr(unsafe.Pointer(size)),
		uintptr(order),
		uintptr(af),
		uintptr(tableClass),
		0,
	)
	return int(ret)
}

func GetExtendedUdpTable(udpTable uintptr, size *uint32, order int32, af uint32, tableClass UDP_TABLE_CLASS) int {
	ret, _, _ := procGetExtendedUdpTable.Call(
		udpTable,
		uintptr(unsafe.Pointer(size)),
		uintptr(order),
		uintptr(af),
		uintptr(tableClass),
		0,
	)
	return int(ret)
}

func GetBestRoute(destAddr, sourceAddr uint32, bestRoute *MIB_IPFORWARDROW) int {
	ret, _, _ := procGetBestRoute.Call(
		uintptr(destAddr),
		uintptr(sourceAddr),
		uintptr(unsafe.Pointer(bestRoute)),
	)
	return int(ret)
}

func GetIpForwardTable(table *MIB_IPFORWARDTABLE, size *uint32, order int32) int {
	ret, _, _ := procGetIpForwardTable.Call(
		uintptr(unsafe.Pointer(table)),
		uintptr(unsafe.Pointer(size)),
		uintptr(order),
	)
	return int(ret)
}

func GetInterfaceInfo(ifTable *IP_INTERFACE_INFO, outBufLen *uint32) int {
	ret, _, _ := procGetInterfaceInfo.Call(
		uintptr(unsafe.Pointer(ifTable)),
		uintptr(unsafe.Pointer(outBufLen)),
	)
	return int(ret)
}

func GetIfTable(table *MIB_IFTABLE, size *uint32, order int32) int {
	ret, _, _ := procGetIfTable.Call(
		uintptr(unsafe.Pointer(table)),
		uintptr(unsafe.Pointer(size)),
		uintptr(order),
	)
	return int(ret)
}

func DeleteIpForwardEntry(route *MIB_IPFORWARDROW) uint32 {
	ret, _, _ := procDeleteIpForwardEntry.Call(
		uintptr(unsafe.Pointer(route)),
	)
	return uint32(ret)
}

func CreateIpForwardEntry(route *MIB_IPFORWARDROW) uint32 {
	ret, _, _ := procCreateIpForwardEntry.Call(
		uintptr(unsafe.Pointer(route)),
	)
	return uint32(ret)
}
