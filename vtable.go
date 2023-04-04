package go_gpedit

type GroupPolicyObjectVtable struct {
	QueryInterface        uintptr
	AddRef                uintptr
	Release               uintptr
	New                   uintptr
	OpenDSGPO             uintptr
	OpenLocalMachineGPO   uintptr
	OpenRemoteMachineGPO  uintptr
	Save                  uintptr
	Delete                uintptr
	GetName               uintptr
	GetDisplayName        uintptr
	SetDisplayName        uintptr
	GetPath               uintptr
	GetDSPath             uintptr
	GetFileSysPath        uintptr
	GetRegistryKey        uintptr
	GetOptions            uintptr
	SetOptions            uintptr
	GetType               uintptr
	GetMachineName        uintptr
	GetPropertySheetPages uintptr
}
