package main

import (
	"fmt"
	"github.com/gorpher/gowin32"
	"image/png"
	"os"
)

func main() {
	//filename := "D:\\download\\TrafficMonitor_V1.83_x86\\TrafficMonitor\\TrafficMonitor.exe"
	filename := "C:\\Program Files (x86)\\Tencent\\WeChat\\WeChat.exe"
	img, err := gowin32.ExtractPrivateExtractIcons(filename, 128, 128)
	if err != nil {
		panic(err)
	}
	fp, _ := os.Create("output0.png")
	err = png.Encode(fp, img)
	if err != nil {
		fmt.Println(err)
	}
	fp.Close()

	//img, err = gowin32.ExtractIconToImageByExt(filename)
	//if err != nil {
	//	panic(err)
	//}
	//
	//fp, _ = os.Create("output1.png")
	//defer fp.Close()
	//err = png.Encode(fp, img)
	//if err != nil {
	//	fmt.Println(err)
	//}

}
