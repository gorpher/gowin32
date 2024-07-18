package main

import (
	"fmt"
	"github.com/gorpher/gowin32"
	"os"
)

func main() {
	//users := gowin32.NetUserEnum()
	//json.NewEncoder(os.Stdout).Encode(users)
	//for _, user := range users {
	//	fmt.Println(user.UserId, user.Name, user.PrimaryGroupId)
	////}
	groups := gowin32.NetGroupEnum()
	for _, group := range groups {
		fmt.Println(group.GroupName)
	}
	//getUsers := gowin32.NetGroupGetUsers("Administrators")
	//json.NewEncoder(os.Stdout).Encode(getUsers)

	//information := gowin32.NetQueryDisplayUserInformation()
	//for _, info := range information {
	//	fmt.Println("user", info)
	//}
	//groupInformation := gowin32.NetQueryDisplayGroupInformation()
	//for _, info := range groupInformation {
	//	fmt.Println("group", info)
	//}
	//mInformation := gowin32.NetQueryDisplayMachineInformation()
	//for _, info := range mInformation {
	//	fmt.Println("machine", info)
	//}

	shareDir := "C:\\user_files\\test111"
	shareName := "test111"
	username := "administrator"

	shareList := gowin32.NetShareEnum()
	for _, info := range shareList {
		if info.Path == shareDir || info.Netname == shareName {
			fmt.Printf("共享 %s: %s 已存在\n", shareName, shareDir)
			return
		}
	}
	exists, err := gowin32.FileExists(shareDir)
	if err != nil {
		panic(err)
	}
	if !exists {
		err = os.MkdirAll(shareDir, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	err = gowin32.AddNetShare(username, shareDir, shareName)
	if err != nil {
		panic(err)
	}

	err = gowin32.DelNetShare(shareDir)
	if err != nil {
		panic(err)
	}
}
