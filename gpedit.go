package go_gpedit

import (
	"log"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

type GPO_OPEN_FLAG uint32

const (
	GPO_OPEN_LOAD_REGISTRY = GPO_OPEN_FLAG(0x00000001)
	GPO_OPEN_READ_ONLY     = GPO_OPEN_FLAG(0x00000002)
)

type GPO_OPTION uint32

const (
	GPO_OPTION_DISABLE_USER    = GPO_OPTION(0x00000001)
	GPO_OPTION_DISABLE_MACHINE = GPO_OPTION(0x00000002)
)

type GPO_SECTION_TYPE uint32

const (
	GPO_SECTION_ROOT    = GPO_SECTION_TYPE(0)
	GPO_SECTION_USER    = GPO_SECTION_TYPE(1)
	GPO_SECTION_MACHINE = GPO_SECTION_TYPE(2)
)

type GROUP_POLICY_OBJECT_TYPE uint32

const (
	GPOTypeLocal      = GROUP_POLICY_OBJECT_TYPE(0)
	GPOTypeRemote     = GROUP_POLICY_OBJECT_TYPE(1)
	GPOTypeDS         = GROUP_POLICY_OBJECT_TYPE(2)
	GPOTypeLocalUser  = GROUP_POLICY_OBJECT_TYPE(3)
	GPOTypeLocalGroup = GROUP_POLICY_OBJECT_TYPE(4)
)

type COINIT uint32

const (
	COINIT_APARTMENTTHREADED = COINIT(0x2)
	COINIT_MULTITHREADED     = COINIT(0x0)
	COINIT_DISABLE_OLE1DDE   = COINIT(0x4)
	COINIT_SPEED_OVER_MEMORY = COINIT(0x8)
)

var (
	CLSID_GroupPolicyObject = windows.GUID{
		Data1: 0xea502722,
		Data2: 0xa23d,
		Data3: 0x11d1,
		Data4: [8]byte{0xa7, 0xd3, 0x0, 0x0, 0xf8, 0x75, 0x71, 0xe3},
	}
	IID_IGroupPolicyObject = windows.GUID{
		Data1: 0xea502723,
		Data2: 0xa23d,
		Data3: 0x11d1,
		Data4: [8]byte{0xa7, 0xd3, 0x0, 0x0, 0xf8, 0x75, 0x71, 0xe3},
	}
)

var (
	REGISTRY_EXTENSION_GUID = windows.GUID{
		Data1: 0x35378EAC,
		Data2: 0x683F,
		Data3: 0x11D2,
		Data4: [8]byte{0xA8, 0x9A, 0x00, 0xC0, 0x4F, 0xBB, 0xCF, 0xA2},
	}
	CLSID_GPESnapIn = windows.GUID{
		Data1: 0x8fc0b734,
		Data2: 0xa0e1,
		Data3: 0x11d1,
		Data4: [8]byte{0xa7, 0xd3, 0x0, 0x0, 0xf8, 0x75, 0x71, 0xe3},
	}
)

// Compatibilty for WindowAPI defined types in golang
type WINBOOL int

// Constants for readability
const (
	TRUE  = 1
	FALSE = 0
)

type GroupPolicyObject struct {
	pvObj  uintptr
	vtable *GroupPolicyObjectVtable
}

func InitializeCOM() error {
	err := coInitialize(uintptr(0), uint32(COINIT_APARTMENTTHREADED))
	if err != nil {
		return err
	}
	return nil
}

func NewGroupPolicyObject() (*GroupPolicyObject, error) {
	pvObj, err := coCreateInstance(&CLSID_GroupPolicyObject, nil, windows.CLSCTX_INPROC_SERVER, &IID_IGroupPolicyObject)
	if err != nil {
		return nil, err
	}
	vtable := (*GroupPolicyObjectVtable)(unsafe.Pointer(pvObj))
	return &GroupPolicyObject{
		pvObj:  pvObj,
		vtable: vtable,
	}, nil
}

func (gpo *GroupPolicyObject) QueryInterface(riid *windows.GUID) (any /* ppvObj */, error) {
	ret, _, err := syscall.SyscallN(
		gpo.vtable.QueryInterface,
		gpo.pvObj,
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(&gpo.pvObj)),
	)
	if err != windows.Errno(0) {
		log.Printf("%v", err)
	}
	return &gpo.pvObj, WinErrHandler(err, ret)
}

func (gpo *GroupPolicyObject) AddRef() (uint, error) {
	ret, _, err := syscall.SyscallN(
		gpo.vtable.AddRef,
		gpo.pvObj,
	)
	return uint(ret), WinErrHandler(err)
}

func (gpo *GroupPolicyObject) Release() (uint, error) {
	ret, _, err := syscall.SyscallN(
		gpo.vtable.Release,
		gpo.pvObj,
	)
	return uint(ret), WinErrHandler(err)
}

func (gpo *GroupPolicyObject) New(pszDomainName string, pszDisplayName string, dwFlags uint32) error {
	s1, err := windows.UTF16PtrFromString(pszDomainName)
	if err != nil {
		return err
	}
	s2, err := windows.UTF16PtrFromString(pszDisplayName)
	if err != nil {
		return err
	}

	ret, _, err := syscall.SyscallN(
		gpo.vtable.New,
		gpo.pvObj,
		uintptr(unsafe.Pointer(s1)),
		uintptr(unsafe.Pointer(s2)),
		uintptr(dwFlags),
	)
	if err != windows.Errno(0) {
		log.Printf("Error code: %v", err)
	}
	return WinErrHandler(ret)
}

func (gpo *GroupPolicyObject) OpenDSGPO(pszPath string, dwFlags uint32) error {
	s1, err := windows.UTF16PtrFromString(pszPath)
	if err != nil {
		return err
	}

	ret, _, err := syscall.SyscallN(
		gpo.vtable.OpenDSGPO,
		gpo.pvObj,
		uintptr(unsafe.Pointer(s1)),
		uintptr(dwFlags),
	)
	return WinErrHandler(err, ret)
}

func (gpo *GroupPolicyObject) OpenLocalMachineGPO(dwFlags GPO_OPEN_FLAG) error {
	ret, _, err := syscall.SyscallN(
		gpo.vtable.OpenLocalMachineGPO,
		gpo.pvObj,
		uintptr(dwFlags),
	)
	return WinErrHandler(err, ret)
}

func (gpo *GroupPolicyObject) OpenRemoteMachineGPO(pszComputerName string, dwFlags uint32) error {
	s1, err := windows.UTF16PtrFromString(pszComputerName)
	if err != nil {
		return err
	}

	ret, _, err := syscall.SyscallN(
		gpo.vtable.OpenRemoteMachineGPO,
		gpo.pvObj,
		uintptr(unsafe.Pointer(s1)),
		uintptr(dwFlags),
	)
	return WinErrHandler(err, ret)
}

func (gpo *GroupPolicyObject) Save(bMachine WINBOOL, bAdd WINBOOL, pGuidExtension *windows.GUID, pGuid *windows.GUID) error {
	ret, _, err := syscall.SyscallN(
		gpo.vtable.Save,
		gpo.pvObj,
		uintptr(bMachine),
		uintptr(bAdd),
		uintptr(unsafe.Pointer(pGuidExtension)),
		uintptr(unsafe.Pointer(pGuid)),
	)
	return WinErrHandler(err, ret)
}

func (gpo *GroupPolicyObject) Delete() error {
	ret, _, err := syscall.SyscallN(
		gpo.vtable.Delete,
		gpo.pvObj,
	)
	return WinErrHandler(err, ret)
}

func (gpo *GroupPolicyObject) GetName(pszName string, cchMaxLength int) error {
	s1, err := windows.UTF16PtrFromString(pszName)
	if err != nil {
		return nil
	}

	ret, _, err := syscall.SyscallN(
		gpo.vtable.GetName,
		gpo.pvObj,
		uintptr(unsafe.Pointer(s1)),
		uintptr(cchMaxLength),
	)
	return WinErrHandler(err, ret)
}

func (gpo *GroupPolicyObject) GetDisplayName(pszName string, cchMaxLength int) error {
	s1, err := windows.UTF16PtrFromString(pszName)
	if err != nil {
		return err
	}

	ret, _, err := syscall.SyscallN(
		gpo.vtable.GetDisplayName,
		gpo.pvObj,
		uintptr(unsafe.Pointer(s1)),
		uintptr(cchMaxLength),
	)
	return WinErrHandler(err, ret)
}

func (gpo *GroupPolicyObject) SetDisplayName(pszName string) error {
	s1, err := windows.UTF16PtrFromString(pszName)
	if err != nil {
		return err
	}

	ret, _, err := syscall.SyscallN(
		gpo.vtable.SetDisplayName,
		gpo.pvObj,
		uintptr(unsafe.Pointer(s1)),
	)
	return WinErrHandler(err, ret)
}

func (gpo *GroupPolicyObject) GetPath(pszPath string, cchMaxPath int) error {
	s1, err := windows.UTF16PtrFromString(pszPath)
	if err != nil {
		return err
	}

	ret, _, err := syscall.SyscallN(
		gpo.vtable.GetPath,
		gpo.pvObj,
		uintptr(unsafe.Pointer(s1)),
		uintptr(cchMaxPath),
	)
	return WinErrHandler(err, ret)
}

func (gpo *GroupPolicyObject) GetDSPath(dwSection uint32, pszPath string, cchMaxPath int) error {
	s1, err := windows.UTF16PtrFromString(pszPath)
	if err != nil {
		return err
	}

	ret, _, err := syscall.SyscallN(
		gpo.vtable.GetDSPath,
		gpo.pvObj,
		uintptr(dwSection),
		uintptr(unsafe.Pointer(s1)),
		uintptr(cchMaxPath),
	)
	return WinErrHandler(err, ret)
}

func (gpo *GroupPolicyObject) GetFileSysPath(dwSection uint32, pszPath string, cchMaxPath int) error {
	s1, err := windows.UTF16PtrFromString(pszPath)
	if err != nil {
		return err
	}

	ret, _, err := syscall.SyscallN(
		gpo.vtable.GetFileSysPath,
		gpo.pvObj,
		uintptr(dwSection),
		uintptr(unsafe.Pointer(s1)),
		uintptr(cchMaxPath),
	)
	return WinErrHandler(err, ret)
}

func (gpo *GroupPolicyObject) GetRegistryKey(dwSection uint32, hKey *windows.Handle) error {
	ret, _, err := syscall.SyscallN(
		gpo.vtable.GetRegistryKey,
		gpo.pvObj,
		uintptr(dwSection),
		uintptr(unsafe.Pointer(hKey)),
	)
	return WinErrHandler(err, ret)
}

func (gpo *GroupPolicyObject) GetOptions(dwOptions *uint32) error {
	ret, _, err := syscall.SyscallN(
		gpo.vtable.GetOptions,
		gpo.pvObj,
		uintptr(unsafe.Pointer(dwOptions)),
	)
	return WinErrHandler(err, ret)
}

func (gpo *GroupPolicyObject) SetOptions(dwOptions uint32, dwMask uint32) error {
	ret, _, err := syscall.SyscallN(
		gpo.vtable.SetOptions,
		gpo.pvObj,
		uintptr(dwOptions),
		uintptr(dwMask),
	)
	return WinErrHandler(err, ret)
}

func (gpo *GroupPolicyObject) GetType(gpoType *GROUP_POLICY_OBJECT_TYPE) error {
	ret, _, err := syscall.SyscallN(
		gpo.vtable.GetType,
		gpo.pvObj,
		uintptr(unsafe.Pointer(gpoType)),
	)
	return WinErrHandler(err, ret)
}

func (gpo *GroupPolicyObject) GetMachineName(pszName string, cchMaxLength int) error {
	s1, err := windows.UTF16PtrFromString(pszName)
	if err != nil {
		return err
	}

	ret, _, err := syscall.SyscallN(
		gpo.vtable.GetMachineName,
		gpo.pvObj,
		uintptr(unsafe.Pointer(s1)),
		uintptr(cchMaxLength),
	)
	return WinErrHandler(err, ret)
}

/*func (gpo *GroupPolicyObject) GetPropertySheetPages(hPages **HPROPSHEETPAGE, uPageCount *uint) error {
	ret, _, err := syscall.SyscallN(
		gpo.vtable.GetDSPath,
		gpo.pvObj,
		uintptr(unsafe.Pointer(hPages)),
		uintptr(unsafe.Pointer(uPageCount)),
	)
	return WinErrHandler(err, ret)
}*/
