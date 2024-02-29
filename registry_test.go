package gowin32

import (
	"testing"
)

func TestGetSubKeys(t *testing.T) {
	const UserAPP = "Software\\Microsoft\\Windows\\CurrentVersion\\Uninstall"
	regKey, err := OpenRegKey(RegRootHKCU, UserAPP, true)
	if err != nil {
		panic(err)
	}
	defer regKey.Close()

	keys, err := regKey.GetSubKeys()
	if err != nil {
		panic(err)
	}
	for _, key := range keys {
		t.Log(key)
	}
}
