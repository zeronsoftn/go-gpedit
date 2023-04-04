package go_gpedit

import (
	"golang.org/x/sys/windows"
	"unsafe"
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

var (
	CLSID_GroupPolicyObject = windows.GUID{
		Data1: 0xea502722,
		Data2: 0xbaa8,
		Data3: 0x11d1,
		Data4: [8]byte{0xbf, 0x3c, 0x00, 0xc0, 0x4f, 0xd8, 0xd5, 0x96},
	}
	IID_IGroupPolicyObject = windows.GUID{
		Data1: 0xea502723,
		Data2: 0xbaa8,
		Data3: 0x11d1,
		Data4: [8]byte{0xbf, 0x3c, 0x00, 0xc0, 0x4f, 0xd8, 0xd5, 0x96},
	}
)

type GroupPolicyObject struct {
	pvObj  uintptr
	vtable *GroupPolicyObjectVtable
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
	return nil, nil
}

func (gpo *GroupPolicyObject) AddRef() (uint, error) {
	return 0, nil
}

func (gpo *GroupPolicyObject) Release() (uint, error) {
	return 0, nil
}

func (gpo *GroupPolicyObject) New(pszDomainName string, pszDisplayName string, dwFlags uint32) error {
	return nil
}

func (gpo *GroupPolicyObject) OpenDSGPO(pszPath string, dwFlags uint32) error {
	return nil
}

func (gpo *GroupPolicyObject) OpenLocalMachineGPO(dwFlags GPO_OPEN_FLAG) error {
	return nil
}

func (gpo *GroupPolicyObject) OpenRemoteMachineGPO(pszComputerName string, dwFlags uint32) error {
	return nil
}

func (gpo *GroupPolicyObject) Save(bMachine bool, bAdd bool, pGuidExtension *windows.GUID, pGuid *windows.GUID) error {
	return nil
}

func (gpo *GroupPolicyObject) Delete() error {
	return nil
}

func (gpo *GroupPolicyObject) GetName(pszName string, cchMaxLength int) error {
	return nil
}

func (gpo *GroupPolicyObject) GetDisplayName(pszName string, cchMaxLength int) error {
	return nil
}

func (gpo *GroupPolicyObject) SetDisplayName(pszName string) error {
	return nil
}

func (gpo *GroupPolicyObject) GetPath(pszPath string, cchMaxPath int) error {
	return nil
}

func (gpo *GroupPolicyObject) GetDSPath(dwSection uint32, pszPath string, cchMaxPath int) error {
	return nil
}

func (gpo *GroupPolicyObject) GetFileSysPath(dwSection uint32, pszPath string, cchMaxPath int) error {
	return nil
}

func (gpo *GroupPolicyObject) GetRegistryKey(dwSection uint32, hKey *windows.Handle) error {
	return nil
}

func (gpo *GroupPolicyObject) GetOptions(dwOptions *uint32) error {
	return nil
}

func (gpo *GroupPolicyObject) SetOptions(dwOptions uint32, dwMask uint32) error {
	return nil
}

func (gpo *GroupPolicyObject) GetType(gpoType *GROUP_POLICY_OBJECT_TYPE) error {
	return nil
}

func (gpo *GroupPolicyObject) GetMachineName(pszName string, cchMaxLength int) error {
	return nil
}

//func (gpo *GroupPolicyObject) GetPropertySheetPages(hPages **HPROPSHEETPAGE, uPageCount *uint) error {
//	return 0
//}
