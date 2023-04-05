package main

import (
	"log"

	go_gpedit "github.com/zeronsoftn/go-gpedit"
	"golang.org/x/sys/windows"
	"golang.org/x/sys/windows/registry"
)

func main() {
	gpedit, err := go_gpedit.NewGroupPolicyObject()
	if err != nil {
		log.Fatalln(err)
	}

	defer gpedit.Release()

	if err = gpedit.OpenLocalMachineGPO(go_gpedit.GPO_OPEN_LOAD_REGISTRY); err != nil {
		log.Fatalln(err)
	}

	var keyMachine windows.Handle
	if err = gpedit.GetRegistryKey(uint32(go_gpedit.GPO_SECTION_MACHINE), &keyMachine); err != nil {
		log.Fatalln(err)
	}

	k, old, err := registry.CreateKey(registry.Key(keyMachine), "SOFTWARE\\Policies\\Microsoft\\FVE", registry.SET_VALUE)
	if err != nil {
		log.Fatalln(err)
	}
	if old {
		log.Println("Overwriting existing value...")
	}
	defer k.Close()

	err = k.SetDWordValue("UseEnhancedPin", 1)
	if err != nil {
		log.Fatalln(err)
	}

	if err = gpedit.Save(go_gpedit.TRUE,
		go_gpedit.TRUE,
		&go_gpedit.REGISTRY_EXTENSION_GUID,
		&go_gpedit.CLSID_GPESnapIn,
	); err != nil {
		log.Fatalln(err)
	}

	// See https://bitbandit.org/20190622/configure-wsus-gpo-programmatically/
}
