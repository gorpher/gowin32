package main

import (
	"fmt"
	"github.com/gorpher/gowin32"
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
}
