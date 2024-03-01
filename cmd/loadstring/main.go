package main

import (
	"fmt"
	"github.com/gorpher/gowin32"
	"github.com/gorpher/gowin32/win"
	"syscall"
)

func main() {
	hd, err2 := syscall.LoadLibrary(`shell32.dll`)
	if err2 != nil {
		panic(err2)
	}
	defer syscall.FreeLibrary(hd)
	var length int32 = 1024

	for i := 0; i < 25536; i++ {
		chars := make([]uint16, length)
		k := win.LoadString(win.HINSTANCE(hd), uint32(i), &chars[0], length)
		if k == 0 {
			continue
		}
		chars = chars[:k]
		fmt.Println(i, gowin32.LpstrToString(&chars[0]))
	}

}
