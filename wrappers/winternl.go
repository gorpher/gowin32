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

type UNICODE_STRING struct {
	Length        uint16
	MaximumLength uint16
	Buffer        uintptr
}

type RTL_USER_PROCESS_PARAMETERS struct {
	Reserved1     [16]byte
	Reserved2     [10]uintptr
	ImagePathName UNICODE_STRING
	CommandLine   UNICODE_STRING
}

type PEB struct {
	Reserved1              [2]byte
	BeingDebugged          byte
	Reserved2              [1]byte
	Reserved3              [2]uintptr
	Ldr                    uintptr
	ProcessParameters      uintptr
	Reserved4              [104]byte
	Reserved5              [52]uintptr
	PostProcessInitRoutine uintptr
	Reserved6              [128]byte
	Reserved7              [1]uintptr
	SessionId              uint32
}

type OBJECT_ATTRIBUTES struct {
	Length                   uint32
	RootDirectory            syscall.Handle
	ObjectName               *UNICODE_STRING
	Attributes               uint32
	SecurityDescriptor       uintptr
	SecurityQualityOfService uintptr
}

type PROCESS_BASIC_INFORMATION struct {
	Reserved1       uintptr
	PebBaseAddress  uintptr
	Reserved2       [2]uintptr
	UniqueProcessId uintptr
	Reserved3       uintptr
}
type (
	FILE_DIRECTORY_INFORMATION struct {
		NextEntryOffset uint32
		FileIndex       uint32
		CreationTime    uint64
		LastAccessTime  uint64
		LastWriteTime   uint64
		ChangeTime      uint64
		EndOfFile       uint64
		AllocationSize  uint64
		FileAttributes  uint32
		FileNameLength  uint32
		FileName        [1]uint16
	}

	FILE_NAMES_INFORMATION struct {
		NextEntryOffset uint32
		FileIndex       uint32
		FileNameLength  uint32
		FileName        [1]uint16
	}

	IO_STATUS_BLOCK struct {
		StatusPointer uintptr
		Information   uintptr
	}
)

const (
	//ProcessBasicInformation = 0
	ProcessWow64Information = 26
)

func NT_SUCCESS(status uint32) bool {
	return int32(uintptr(status)) >= 0
}

func NT_INFORMATION(status uint32) bool {
	return (status >> 30) == 1
}

func NT_WARNING(status uint32) bool {
	return (status >> 30) == 2
}

func NT_ERROR(status uint32) bool {
	return (status >> 30) == 3
}

var (
	modntdll = syscall.NewLazyDLL("ntdll.dll")

	procNtQueryInformationProcess   = modntdll.NewProc("NtQueryInformationProcess")
	procRtlFreeUnicodeString        = modntdll.NewProc("RtlFreeUnicodeString")
	procRtlInitUnicodeString        = modntdll.NewProc("RtlInitUnicodeString")
	procNtSetQuotaInformationFile   = modntdll.NewProc("NtSetQuotaInformationFile")
	procNtQueryQuotaInformationFile = modntdll.NewProc("NtQueryQuotaInformationFile")
)

// NtQueryInformationProcess
// __kernel_entry NTSTATUS NtQueryInformationProcess(
//
//	[in]            HANDLE           ProcessHandle,
//	[in]            PROCESSINFOCLASS ProcessInformationClass,
//	[out]           PVOID            ProcessInformation,
//	[in]            ULONG            ProcessInformationLength,
//	[out, optional] PULONG           ReturnLength
//
// );
func NtQueryInformationProcess(processHandle syscall.Handle, processInformationClass int32, processInformation *byte, processInformationLength uint32,
	returnLength *uint32) uint32 {
	r1, _, _ := syscall.SyscallN(
		procNtQueryInformationProcess.Addr(),
		uintptr(processHandle),
		uintptr(processInformationClass),
		uintptr(unsafe.Pointer(processInformation)),
		uintptr(processInformationLength),
		uintptr(unsafe.Pointer(returnLength)))
	return uint32(r1)
}

func RtlFreeUnicodeString(unicodeString *UNICODE_STRING) {
	syscall.SyscallN(
		procRtlFreeUnicodeString.Addr(),
		uintptr(unsafe.Pointer(unicodeString)))
}

func RtlInitUnicodeString(destinationString *UNICODE_STRING, sourceString *uint16) {
	syscall.Syscall(
		procRtlInitUnicodeString.Addr(),
		2,
		uintptr(unsafe.Pointer(destinationString)),
		uintptr(unsafe.Pointer(sourceString)),
		0)
}

// NtQueryQuotaInformationFile
// __kernel_entry NTSYSCALLAPI NTSTATUS NtQueryQuotaInformationFile(
//
//	[in]           HANDLE           FileHandle,
//	[out]          PIO_STATUS_BLOCK IoStatusBlock,
//	[out]          PVOID            Buffer,
//	[in]           ULONG            Length,
//	[in]           BOOLEAN          ReturnSingleEntry,
//	[in, optional] PVOID            SidList,
//	[in]           ULONG            SidListLength,
//	[in, optional] PSID             StartSid,
//	[in]           BOOLEAN          RestartScan
//
// );
func NtQueryQuotaInformationFile(fileHandle syscall.Handle, ioStatusBlock *IO_STATUS_BLOCK, buffer *byte, length uint32,
	returnSingleEntry uint32, sidList *byte, sidListLength uint32, startSid uint32, restartScan uint32) uint32 {
	r1, _, _ := syscall.SyscallN(
		procRtlInitUnicodeString.Addr(),
		uintptr(fileHandle),
		uintptr(unsafe.Pointer(ioStatusBlock)),
		uintptr(unsafe.Pointer(buffer)),
		uintptr(length),
		uintptr(returnSingleEntry),
		uintptr(unsafe.Pointer(sidList)),
		uintptr(sidListLength),
		uintptr(startSid),
		uintptr(restartScan))
	return uint32(r1)
}
