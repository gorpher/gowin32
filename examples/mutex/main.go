package main

import (
	"errors"
	"fmt"
	"github.com/gorpher/gowin32/wrappers"
	"time"
)

func main() {
	if IsRunning() {
		fmt.Println("已经运行")
		return
	}
	ticker := time.NewTicker(time.Second)
	for t := range ticker.C {
		fmt.Println(t.String())
	}
}

func IsRunning() bool {
	_, err := wrappers.CreateMutex(nil, true, wrappers.Lpcwstr("luancher"))
	if err != nil && errors.Is(err, wrappers.ERROR_ALREADY_EXISTS) {
		return true
	}
	return false
}
