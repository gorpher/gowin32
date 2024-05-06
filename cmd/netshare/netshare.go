package main

import (
	"flag"
	"fmt"
	"github.com/gorpher/gowin32"
)

func main() {
	var (
		username string
		shareDir string
	)
	flag.StringVar(&username, "username", "", "用户名")
	flag.StringVar(&shareDir, "shareDir", "", "共享目录")
	flag.Parse()
	if username == "" {
		fmt.Println("请输入用户名")
		return
	}
	if shareDir == "" {
		fmt.Println("请输入共享目录")
		return
	}
	err := gowin32.AddNetShare(username, shareDir, username, gowin32.FileAllAccess)
	if err != nil {
		panic(err)
	}
}
