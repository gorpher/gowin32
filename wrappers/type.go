//go:build windows
// +build windows

package wrappers

import (
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

// http://msdn.microsoft.com/en-us/library/s3f49ktz.aspx
// http://msdn.microsoft.com/en-us/library/windows/desktop/aa383751.aspx
// ATOM                  WORD
// BOOL                  int32
// BOOLEAN               byte
// BYTE                  byte
// CCHAR                 int8
// CHAR                  int8
// COLORREF              DWORD
// DWORD                 uint32
// DWORDLONG             ULONGLONG
// DWORD_PTR             ULONG_PTR
// DWORD32               uint32
// DWORD64               uint64
// FLOAT                 float32
// HACCEL                HANDLE
// HALF_PTR              struct{} // ???
// HANDLE                PVOID
// HBITMAP               HANDLE
// HBRUSH                HANDLE
// HCOLORSPACE           HANDLE
// HCONV                 HANDLE
// HCONVLIST             HANDLE
// HCURSOR               HANDLE
// HDC                   HANDLE
// HDDEDATA              HANDLE
// HDESK                 HANDLE
// HDROP                 HANDLE
// HDWP                  HANDLE
// HENHMETAFILE          HANDLE
// HFILE                 HANDLE
// HFONT                 HANDLE
// HGDIOBJ               HANDLE
// HGLOBAL               HANDLE
// HHOOK                 HANDLE
// HICON                 HANDLE
// HINSTANCE             HANDLE
// HKEY                  HANDLE
// HKL                   HANDLE
// HLOCAL                HANDLE
// HMENU                 HANDLE
// HMETAFILE             HANDLE
// HMODULE               HANDLE
// HPALETTE              HANDLE
// HPEN                  HANDLE
// HRESULT               int32
// HRGN                  HANDLE
// HSZ                   HANDLE
// HWINSTA               HANDLE
// HWND                  HANDLE
// INT                   int32
// INT_PTR               uintptr
// INT8                  int8
// INT16                 int16
// INT32                 int32
// INT64                 int64
// LANGID                WORD
// LCID                  DWORD
// LCTYPE                DWORD
// LGRPID                DWORD
// LONG                  int32
// LONGLONG              int64
// LONG_PTR              uintptr
// LONG32                int32
// LONG64                int64
// LPARAM                LONG_PTR
// LPBOOL                *BOOL
// LPBYTE                *BYTE
// LPCOLORREF            *COLORREF
// LPCSTR                *int8
// LPCTSTR               LPCWSTR
// LPCVOID               unsafe.Pointer
// LPCWSTR               *WCHAR
// LPDWORD               *DWORD
// LPHANDLE              *HANDLE
// LPINT                 *INT
// LPLONG                *LONG
// LPSTR                 *CHAR
// LPTSTR                LPWSTR
// LPVOID                unsafe.Pointer
// LPWORD                *WORD
// LPWSTR                *WCHAR
// LRESULT               LONG_PTR
// PBOOL                 *BOOL
// PBOOLEAN              *BOOLEAN
// PBYTE                 *BYTE
// PCHAR                 *CHAR
// PCSTR                 *CHAR
// PCTSTR                PCWSTR
// PCWSTR                *WCHAR
// PDWORD                *DWORD
// PDWORDLONG            *DWORDLONG
// PDWORD_PTR            *DWORD_PTR
// PDWORD32              *DWORD32
// PDWORD64              *DWORD64
// PFLOAT                *FLOAT
// PHALF_PTR             *HALF_PTR
// PHANDLE               *HANDLE
// PHKEY                 *HKEY
// PINT_PTR              *INT_PTR
// PINT8                 *INT8
// PINT16                *INT16
// PINT32                *INT32
// PINT64                *INT64
// PLCID                 *LCID
// PLONG                 *LONG
// PLONGLONG             *LONGLONG
// PLONG_PTR             *LONG_PTR
// PLONG32               *LONG32
// PLONG64               *LONG64
// POINTER_32            struct{} // ???
// POINTER_64            struct{} // ???
// POINTER_SIGNED        uintptr
// POINTER_UNSIGNED      uintptr
// PSHORT                *SHORT
// PSIZE_T               *SIZE_T
// PSSIZE_T              *SSIZE_T
// PSTR                  *CHAR
// PTBYTE                *TBYTE
// PTCHAR                *TCHAR
// PTSTR                 PWSTR
// PUCHAR                *UCHAR
// PUHALF_PTR            *UHALF_PTR
// PUINT                 *UINT
// PUINT_PTR             *UINT_PTR
// PUINT8                *UINT8
// PUINT16               *UINT16
// PUINT32               *UINT32
// PUINT64               *UINT64
// PULONG                *ULONG
// PULONGLONG            *ULONGLONG
// PULONG_PTR            *ULONG_PTR
// PULONG32              *ULONG32
// PULONG64              *ULONG64
// PUSHORT               *USHORT
// PVOID                 unsafe.Pointer
// PWCHAR                *WCHAR
// PWORD                 *WORD
// PWSTR                 *WCHAR
// QWORD                 uint64
// SC_HANDLE             HANDLE
// SC_LOCK               LPVOID
// SERVICE_STATUS_HANDLE HANDLE
// SHORT                 int16
// SIZE_T                ULONG_PTR
// SSIZE_T               LONG_PTR
// TBYTE                 WCHAR
// TCHAR                 WCHAR
// UCHAR                 uint8
// UHALF_PTR             struct{} // ???
// UINT                  uint32
// UINT_PTR              uintptr
// UINT8                 uint8
// UINT16                uint16
// UINT32                uint32
// UINT64                uint64
// ULONG                 uint32
// ULONGLONG             uint64
// ULONG_PTR             uintptr
// ULONG32               uint32
// ULONG64               uint64
// USHORT                uint16
// USN                   LONGLONG
// WCHAR                 uint16
// WORD                  uint16
// WPARAM                UINT_PTR

type (
	HANDLE    uintptr
	HMODULE   HANDLE
	PDWORD    uintptr
	ULONG_PTR uintptr
	LPVOID    uintptr
	LPDWORD   *uint32
	LPCWSTR   *uint16
	LPWSTR    *uint16
	LPBYTE    *byte
	PBYTE     *byte
	DWORD     uint32
	ULONG     uint32
	BOOL      int32
)

type MIB_TCPSTATS struct {
	RtoAlgorithm DWORD
	RtoMin       DWORD
	RtoMax       DWORD
	MaxConn      DWORD
	ActiveOpens  DWORD
	PassiveOpens DWORD
	AttemptFails DWORD
	EstabResets  DWORD
	CurrEstab    DWORD
	InSegs       DWORD
	OutSegs      DWORD
	RetransSegs  DWORD
	InErrs       DWORD
	OutRsts      DWORD
	NumConns     DWORD
}

type MIB_TCPROW_OWNER_PID struct {
	State      DWORD
	LocalAddr  DWORD
	LocalPort  DWORD
	RemoteAddr DWORD
	RemotePort DWORD
	OwningPid  DWORD
}

//type MIB_TCPTABLE_OWNER_PID struct {
//	NumEntries DWORD
//	Table      [1 << 30]MIB_TCPROW_OWNER_PID
//}

type TCP_TABLE_CLASS DWORD

const (
	TCP_TABLE_BASIC_LISTENER TCP_TABLE_CLASS = iota
	TCP_TABLE_BASIC_CONNECTIONS
	TCP_TABLE_BASIC_ALL
	TCP_TABLE_OWNER_PID_LISTENER
	TCP_TABLE_OWNER_PID_CONNECTIONS
	TCP_TABLE_OWNER_PID_ALL
	TCP_TABLE_OWNER_MODULE_LISTENER
	TCP_TABLE_OWNER_MODULE_CONNECTIONS
	TCP_TABLE_OWNER_MODULE_ALL
)

type MIB_UDPROW_OWNER_PID struct {
	LocalAddr DWORD
	LocalPort DWORD
	OwningPid DWORD
}

type MIB_UDP6ROW_OWNER_PID struct {
	LocalAddr    [16]uint8
	LocalScopeId DWORD
	LocalPort    DWORD
	OwningPid    DWORD
}

//type MIB_UDPTABLE_OWNER_PID struct {
//	NumEntries DWORD
//	Table      [1 << 30]MIB_UDPROW_OWNER_PID
//}
//
//type MIB_UDP6TABLE_OWNER_PID struct {
//	NumEntries DWORD
//	Table      [1 << 30]MIB_UDP6ROW_OWNER_PID
//}

type UDP_TABLE_CLASS DWORD

const (
	UDP_TABLE_BASIC UDP_TABLE_CLASS = iota
	UDP_TABLE_OWNER_PID
	UDP_TABLE_OWNER_MODULE
)

type ModuleEntry32 struct {
	Size         uint32
	ModuleID     uint32
	ProcessID    uint32
	GlblcntUsage uint32
	ProccntUsage uint32
	ModBaseAddr  *uint8
	ModBaseSize  uint32
	HModule      HMODULE
	Module       [MAX_MODULE_NAME32 + 1]uint16
	ExePath      [MAX_PATH]uint16
}

type MIB_IPFORWARDROW struct {
	ForwardDest      uint32
	ForwardMask      uint32
	ForwardPolicy    uint32
	ForwardNextHop   uint32
	ForwardIfIndex   uint32
	ForwardType      uint32
	ForwardProto     uint32
	ForwardAge       uint32
	ForwardNextHopAS uint32
	ForwardMetric1   uint32
	ForwardMetric2   uint32
	ForwardMetric3   uint32
	ForwardMetric4   uint32
	ForwardMetric5   uint32
}

//type MIB_IPFORWARDTABLE struct {
//	NumEntries DWORD
//	Table      [1 << 30]MIB_IPFORWARDROW
//}
//
//type IP_INTERFACE_INFO struct {
//	NumAdapters int32
//	Adapter     [1 << 30]IP_ADAPTER_INDEX_MAP
//}

//type IP_ADAPTER_INDEX_MAP struct {
//	Index uint32
//	Name  [MAX_ADAPTER_NAME]uint16
//}

//type MIB_IFTABLE struct {
//	NumEntries uint32
//	Table      [1 << 30]MIB_IFROW
//}

type MIB_IFROW struct {
	Name            [MAX_INTERFACE_NAME_LEN]uint16
	Index           uint32
	Type            uint32
	Mtu             uint32
	Speed           uint32
	PhysAddrLen     uint32
	PhysAddr        [MAXLEN_PHYSADDR]uint8
	AdminStatus     uint32
	OperStatus      uint32
	LastChange      uint32
	InOctets        uint32
	InUcastPkts     uint32
	InNUcastPkts    uint32
	InDiscards      uint32
	InErrors        uint32
	InUnknownProtos uint32
	OutOctets       uint32
	OutUcastPkts    uint32
	OutNUcastPkts   uint32
	OutDiscards     uint32
	OutErrors       uint32
	OutQLen         uint32
	DescrLen        uint32
	Descr           [MAXLEN_IFDESCR]uint8
}

type FWPM_DISPLAY_DATA0 struct {
	Name        *uint16
	Description *uint16
}

type FWPM_SESSION0 struct {
	SessionKey           windows.GUID
	DisplayData          FWPM_DISPLAY_DATA0
	Flags                uint32
	TxnWaitTimeoutInMSec uint32
	ProcessId            uint32
	Sid                  *windows.SID
	Username             *uint16
	KernelMode           int32
}

type FWP_BYTE_BLOB struct {
	size uint32
	data *uint8
}

type FWPM_SUBLAYER0 struct {
	SubLayerKey  windows.GUID // Windows type: GUID
	DisplayData  FWPM_DISPLAY_DATA0
	Flags        uint32
	ProviderKey  *windows.GUID // Windows type: *GUID
	ProviderData FWP_BYTE_BLOB
	Weight       uint16
}

type FWP_VALUE0 struct {
	Type  uint32
	Value uintptr
}

type FWP_CONDITION_VALUE0 FWP_VALUE0

type FWPM_FILTER_CONDITION0 struct {
	FieldKey       windows.GUID // Windows type: GUID
	MatchType      uint32
	ConditionValue FWP_CONDITION_VALUE0
}

type FWPM_ACTION0 struct {
	Type  uint32
	Value windows.GUID
}

type FWPM_FILTER0 struct {
	FilterKey           windows.GUID
	DisplayData         FWPM_DISPLAY_DATA0
	Flags               uint32
	ProviderKey         *windows.GUID
	ProviderData        FWP_BYTE_BLOB
	LayerKey            windows.GUID
	SubLayerKey         windows.GUID
	Weight              FWP_VALUE0
	NumFilterConditions uint32
	FilterCondition     *FWPM_FILTER_CONDITION0
	Action              FWPM_ACTION0
	Offset1             [4]byte
	Context             windows.GUID
	Reserved            *windows.GUID
	FilterId            uint64
	EffectiveWeight     FWP_VALUE0
}

func BstrToString(bstr *uint16) string {
	if bstr == nil {
		return ""
	}
	len := SysStringLen(bstr)
	buf := make([]uint16, len)
	RtlMoveMemory(
		(*byte)(unsafe.Pointer(&buf[0])),
		(*byte)(unsafe.Pointer(bstr)),
		uintptr(2*len))
	return syscall.UTF16ToString(buf)
}

func LpstrToString(lpstr *uint16) string {
	if lpstr == nil {
		return ""
	}
	len := Lstrlen(lpstr)
	if len == 0 {
		return ""
	}
	buf := make([]uint16, len)
	RtlMoveMemory(
		(*byte)(unsafe.Pointer(&buf[0])),
		(*byte)(unsafe.Pointer(lpstr)),
		uintptr(2*len))
	return syscall.UTF16ToString(buf)
}

// Lpcwstr  golang string to  c  lpcwstr Type
func Lpcwstr(items ...string) *uint16 {
	var chars []uint16
	for _, s := range items {
		chars = append(chars, syscall.StringToUTF16(s)...)
	}
	chars = append(chars, 0)
	return &chars[0]
}

// Lpstr  golang string to  c  lpstr Type
func Lpstr(str string) uintptr {
	var buf = make([]byte, len(str)+1)
	copy(buf, str)
	return uintptr(unsafe.Pointer(&buf[0])) //nolint
}
