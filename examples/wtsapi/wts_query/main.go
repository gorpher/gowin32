package main

import (
	"fmt"

	"github.com/gorpher/gone"
	"github.com/gorpher/gowin32"
	"strings"
)

func main() {
	server := gowin32.OpenWTSServer("")
	sessions, err := server.EnumerateSessions()

	if err != nil {
		panic(err)
	}
	fmt.Println(len(sessions))
	for _, session := range sessions {
		fmt.Printf("SessionID: %d , WinStationName: %s ,State  %s", session.SessionID, session.WinStationName, session.State)
		name, err := server.QuerySessionUserName(session.SessionID)
		if err != nil {
			panic(err)
		}

		fmt.Printf(",UserName: %s\n", name)
		if strings.HasPrefix(session.WinStationName, "Console") {
			//if session.SessionID == 2 {
			appName, err := server.QuerySessionApplicationName(session.SessionID)
			if err != nil {
				panic(err)
			}
			fmt.Printf(",AppName: %s\n", appName)
			s, err := server.QuerySessionProcessEx(session.SessionID)
			if err != nil {
				panic(err)
			}
			var sum int64 = 0
			for _, v := range s {
				sum += v.Memory

				if strings.HasPrefix(v.ProcessName, "navicat.exe") ||
					strings.HasPrefix(v.ProcessName, "clion") ||
					strings.HasPrefix(v.ProcessName, "WeChat") {
					fmt.Printf("%s Memory==> %s    PeakMemory==> %s    PagefileUsage==> %s    PeakPagefileUsage==> %s   HandleCount==> %d   NumberOfThreads==> %d ProcessId==> %d \n", v.ProcessName,
						gone.FormatBytesStringOhMyGod(v.Memory),
						gone.FormatBytesStringOhMyGod(v.PeakMemory),
						gone.FormatBytesStringOhMyGod(v.PagefileUsage),
						gone.FormatBytesStringOhMyGod(v.PeakPagefileUsage),
						v.HandleCount,
						v.NumberOfThreads,
						v.ProcessId,
					)

				}
				//sv, _ := json.MarshalIndent(s, "", " ")
				//fmt.Println(string(sv))
			}
			fmt.Printf("Sumary %s\n", gone.FormatBytesStringOhMyGod(sum))

		}
	}
}
