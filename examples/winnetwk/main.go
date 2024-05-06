package main

import (
	"errors"
	"fmt"
	"github.com/gorpher/gowin32"
	"github.com/gorpher/gowin32/wrappers"
	"strings"
)

func main() {
	err := gowin32.WNetCancelConnection2W("Z:")
	if err != nil {
		if !(errors.Is(err, wrappers.ERROR_NOT_CONNECTED) ||
			errors.Is(err, wrappers.ERROR_OPEN_FILES)) {
			fmt.Printf("%#v\n", err)
			panic(err)
		}

	}
	remoteName := "\\\\mount1.com\\share"
	err = setRegMountLabel("Ok", remoteName)
	if err != nil {
		panic(err)
		return
	}

	err = gowin32.WNetAddConnection(
		remoteName,
		"Z:",
		"share",
		"123456",
		wrappers.CONNECT_TEMPORARY)
	if err != nil {
		if !errors.Is(err, wrappers.ERROR_ALREADY_ASSIGNED) {
			fmt.Printf("%#v\n", err)
			panic(err)
		}
	}

}

const RegMountLabelKey = "SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Explorer\\MountPoints2"

func setRegMountLabel(label, path string) error {
	rootKey, err := gowin32.OpenRegKey(gowin32.RegRootHKCU, RegMountLabelKey, true)
	if err != nil {
		return err
	}
	path = strings.ReplaceAll(path, "\\", "#")
	currentKey, err := rootKey.CreateSubKey(path)
	if err != nil {
		return err
	}
	return currentKey.SetValueString("_LabelFromReg", label)
}
