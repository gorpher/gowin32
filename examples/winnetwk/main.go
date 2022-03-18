package main

import (
	"errors"
	"fmt"
	"github.com/gorpher/gowin32"
	"github.com/gorpher/gowin32/wrappers"
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
	err = gowin32.WNetAddConnection("\\\\192.168.51.110\\",
		"Z:",
		"",
		"")
	if err != nil {
		if !errors.Is(err, wrappers.ERROR_ALREADY_ASSIGNED) {
			fmt.Printf("%#v\n", err)
			panic(err)
		}
	}

}
