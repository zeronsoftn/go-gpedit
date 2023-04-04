package main

import (
	go_gpedit "github.com/zeronsoftn/go-gpedit"
	"log"
)

func main() {
	gpedit, err := go_gpedit.NewGroupPolicyObject()
	if err != nil {
		log.Fatalln(err)
	}

	defer gpedit.Release()

	if err := gpedit.OpenLocalMachineGPO(go_gpedit.GPO_OPEN_LOAD_REGISTRY); err != nil {
		log.Fatalln(err)
	}

	// See https://bitbandit.org/20190622/configure-wsus-gpo-programmatically/
}
