package wrappers

import (
	"syscall"
	"unsafe"
)

// Group Policy Editor MMC SnapIn GUID
//
// {8FC0B734-A0E1-11d1-A7D3-0000F87571E3}
var (
	CLSID_GPESnapIn = GUID{0x8fc0b734, 0xa0e1, 0x11d1, [8]byte{0xa7, 0xd3, 0x0, 0x0, 0xf8, 0x75, 0x71, 0xe3}}

	//
	// Group Policy Editor node ids
	//

	//
	// Computer Configuration\Windows Settings
	// {8FC0B737-A0E1-11d1-A7D3-0000F87571E3}
	//

	NODEID_Machine = GUID{0x8fc0b737, 0xa0e1, 0x11d1, [8]byte{0xa7, 0xd3, 0x0, 0x0, 0xf8, 0x75, 0x71, 0xe3}}

	//
	// Computer Configuration\Software Settings
	// {8FC0B73A-A0E1-11d1-A7D3-0000F87571E3}
	//

	NODEID_MachineSWSettings = GUID{0x8fc0b73a, 0xa0e1, 0x11d1, [8]byte{0xa7, 0xd3, 0x0, 0x0, 0xf8, 0x75, 0x71, 0xe3}}

	//
	// User Configuration\Windows Settings
	// {8FC0B738-A0E1-11d1-A7D3-0000F87571E3}
	//

	NODEID_User = GUID{0x8fc0b738, 0xa0e1, 0x11d1, [8]byte{0xa7, 0xd3, 0x0, 0x0, 0xf8, 0x75, 0x71, 0xe3}}

	//
	// User Configuration\Software Settings
	// {8FC0B73C-A0E1-11d1-A7D3-0000F87571E3}
	//

	NODEID_UserSWSettings = GUID{0x8fc0b73c, 0xa0e1, 0x11d1, [8]byte{0xa7, 0xd3, 0x0, 0x0, 0xf8, 0x75, 0x71, 0xe3}}

	//
	// IGPEInformation interface id
	//
	// {8FC0B735-A0E1-11d1-A7D3-0000F87571E3}

	IID_IGPEInformation = GUID{0x8fc0b735, 0xa0e1, 0x11d1, [8]byte{0xa7, 0xd3, 0x0, 0x0, 0xf8, 0x75, 0x71, 0xe3}}

	//
	// Group Policy Object class id
	//
	// {EA502722-A23D-11d1-A7D3-0000F87571E3}

	CLSID_GroupPolicyObject = GUID{0xea502722, 0xa23d, 0x11d1, [8]byte{0xa7, 0xd3, 0x0, 0x0, 0xf8, 0x75, 0x71, 0xe3}}

	//
	// Group Policy Object interface id
	//
	// {EA502723-A23D-11d1-A7D3-0000F87571E3}

	IID_IGroupPolicyObject = GUID{0xea502723, 0xa23d, 0x11d1, [8]byte{0xa7, 0xd3, 0x0, 0x0, 0xf8, 0x75, 0x71, 0xe3}}

	//
	// GUID that identifies the registry extension
	//

	REGISTRY_EXTENSION_GUID = GUID{0x35378EAC, 0x683F, 0x11D2, [8]byte{0xA8, 0x9A, 0x00, 0xC0, 0x4F, 0xBB, 0xCF, 0xA2}}

	//
	// GUID that identifies the comments extension
	//

	ADMXCOMMENTS_EXTENSION_GUID = GUID{0x6C5A2A86, 0x9EB3, 0x42b9, [8]byte{0xAA, 0x83, 0xA7, 0x37, 0x1B, 0xA0, 0x11, 0xB9}}

	//========================================================================================
	//
	// Resultant Set of Policy node ids
	//
	//========================================================================================

	//
	// Resultant Set of Policy MMC SnapIn GUID
	//
	// {6DC3804B-7212-458D-ADB0-9A07E2AE1FA2}

	CLSID_RSOPSnapIn = GUID{0x6dc3804b, 0x7212, 0x458d, [8]byte{0xad, 0xb0, 0x9a, 0x07, 0xe2, 0xae, 0x1f, 0xa2}}

	//
	// Computer Configuration\Windows Settings
	// {BD4C1A2E-0B7A-4A62-A6B0-C0577539C97E}
	//

	NODEID_RSOPMachine = GUID{0xbd4c1a2e, 0x0b7a, 0x4a62, [8]byte{0xa6, 0xb0, 0xc0, 0x57, 0x75, 0x39, 0xc9, 0x7e}}

	//
	// Computer Configuration\Software Settings
	// {6A76273E-EB8E-45DB-94C5-25663A5f2C1A}
	//

	NODEID_RSOPMachineSWSettings = GUID{0x6a76273e, 0xeb8e, 0x45db, [8]byte{0x94, 0xc5, 0x25, 0x66, 0x3a, 0x5f, 0x2c, 0x1a}}

	//
	// User Configuration\Windows Settings
	// {AB87364F-0CEC-4CD8-9BF8-898F34628FB8}
	//

	NODEID_RSOPUser = GUID{0xab87364f, 0x0cec, 0x4cd8, [8]byte{0x9b, 0xf8, 0x89, 0x8f, 0x34, 0x62, 0x8f, 0xb8}}

	//
	// User Configuration\Software Settings
	// {E52C5CE3-FD27-4402-84DE-D9A5F2858910}
	//

	NODEID_RSOPUserSWSettings = GUID{0xe52c5ce3, 0xfd27, 0x4402, [8]byte{0x84, 0xde, 0xd9, 0xa5, 0xf2, 0x85, 0x89, 0x10}}

	//
	// IRSOPInformation interface id
	//
	// {9A5A81B5-D9C7-49EF-9D11-DDF50968C48D}

	IID_IRSOPInformation = GUID{0x9a5a81b5, 0xd9c7, 0x49ef, [8]byte{0x9d, 0x11, 0xdd, 0xf5, 0x09, 0x68, 0xc4, 0x8d}}
)

const (
	//
	// Group Policy Object Section flags
	//

	GPO_SECTION_ROOT    = 0 // Root
	GPO_SECTION_USER    = 1 // User
	GPO_SECTION_MACHINE = 2 // Machine

)
const (
	GPO_OPEN_LOAD_REGISTRY = 0x00000001 // Load the registry files
	GPO_OPEN_READ_ONLY     = 0x00000002 // Open the GPO as read only

	//
	// Group Policy Object option flags
	//

	GPO_OPTION_DISABLE_USER    = 0x00000001 // The user portion of this GPO is disabled
	GPO_OPTION_DISABLE_MACHINE = 0x00000002 // The machine portion of this GPO is disabled

)

type IGroupPolicyObjectVtbl struct {
	IDispatchVtbl
	Delete                uintptr
	GetDisplayName        uintptr
	GetDSPath             uintptr
	GetFileSysPath        uintptr
	GetMachineName        uintptr
	GetName               uintptr
	GetOptions            uintptr
	GetPath               uintptr
	GetPropertySheetPages uintptr
	GetRegistryKey        uintptr
	GetType               uintptr
	New                   uintptr
	OpenDSGPO             uintptr
	OpenLocalMachineGPO   uintptr
	OpenRemoteMachineGPO  uintptr
	/*method*/
	Save           uintptr
	SetDisplayName uintptr
	SetOptions     uintptr
}

type IGroupPolicyObject struct {
	IDispatch
}

func (self *IGroupPolicyObject) Delete() uint32 {
	vtbl := (*IGroupPolicyObjectVtbl)(unsafe.Pointer(self.Vtbl))
	r1, _, _ := syscall.SyscallN(
		vtbl.Delete,
		uintptr(unsafe.Pointer(self)))
	return uint32(r1)
}

// GetDisplayName
// HRESULT GetDisplayName(
//
//	[out] LPOLESTR pszName,
//	[in]  int      cchMaxLength
//
// );
func (self *IGroupPolicyObject) GetDisplayName(pszName *byte, cchMaxLength uint32) uint32 {
	vtbl := (*IGroupPolicyObjectVtbl)(unsafe.Pointer(self.Vtbl))
	r1, _, _ := syscall.SyscallN(
		vtbl.GetDisplayName,
		uintptr(unsafe.Pointer(self)),
		uintptr(unsafe.Pointer(pszName)),
		uintptr(cchMaxLength),
	)
	return uint32(r1)
}

func (self *IGroupPolicyObject) OpenLocalMachineGPO(flag uint32) uint32 {
	vtbl := (*IGroupPolicyObjectVtbl)(unsafe.Pointer(self.Vtbl))
	r1, _, _ := syscall.SyscallN(
		vtbl.OpenLocalMachineGPO,
		uintptr(unsafe.Pointer(self)),
		uintptr(flag))
	return uint32(r1)
}

// GetRegistryKey
// HRESULT GetRegistryKey(
//
//	[in]  DWORD dwSection,
//	[out] HKEY  *hKey
//
// );
func (self *IGroupPolicyObject) GetRegistryKey(dwSection uint32, hKey *syscall.Handle) uint32 {
	vtbl := (*IGroupPolicyObjectVtbl)(unsafe.Pointer(self.Vtbl))
	r1, _, _ := syscall.SyscallN(
		vtbl.GetRegistryKey,
		uintptr(unsafe.Pointer(self)),
		uintptr(unsafe.Pointer(hKey)),
	)
	return uint32(r1)
}

// Save
// HRESULT Save(
//
//	[in] BOOL bMachine,
//	[in] BOOL bAdd,
//	[in] GUID *pGuidExtension,
//	[in] GUID *pGuid
//
// );
func (self *IGroupPolicyObject) Save(bMachine, bAdd uint32, pGuidExtension, pGuid *GUID) uint32 {
	vtbl := (*IGroupPolicyObjectVtbl)(unsafe.Pointer(self.Vtbl))
	r1, _, _ := syscall.SyscallN(
		vtbl.Save,
		uintptr(unsafe.Pointer(self)),
		uintptr(bMachine),
		uintptr(bMachine),
		uintptr(unsafe.Pointer(pGuidExtension)),
		uintptr(unsafe.Pointer(pGuid)),
	)
	return uint32(r1)
}
