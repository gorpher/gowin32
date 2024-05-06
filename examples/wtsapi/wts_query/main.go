package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"

	"github.com/gorpher/gowin32"
)

func main() {
	server := gowin32.OpenWTSServer("")
	sessions, err := server.EnumerateSessions()

	if err != nil {
		panic(err)
	}
	fmt.Println("Session Length:", len(sessions))
	for _, session := range sessions {
		if session.SessionID == 0 {
			continue
		}
		if session.SessionID > math.MaxUint16 {
			continue
		}
		clientName, err := server.QuerySessionClientName(uint(session.SessionID))
		if err != nil {
			return
		}
		initialProgram, err := server.QuerySessionInitialProgram(uint(session.SessionID))
		if err != nil {
			return
		}
		applicationName, err := server.QuerySessionApplicationName(uint(session.SessionID))
		if err != nil {
			return
		}
		remoteOk, err := server.QuerySessionIsRemoteSession(uint(session.SessionID))
		if err != nil {
			log.Fatal(err)
			return
		}
		if !remoteOk {
			continue
		}
		fmt.Println(remoteOk)
		info, err := server.QuerySessionClientInfo(uint(session.SessionID))
		if err != nil {
			log.Fatal(err)
			return
		}
		infoBy, err := json.Marshal(info)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println(string(infoBy))
		address, err := server.QuerySessionClientAddress(uint(session.SessionID))
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println(address.String())

		fmt.Printf("ClientName: %s ,InitialProgram: %s ,ApplicationName: %s ,SessionID: %d , WinStationName: %s ,State  %s", clientName, applicationName, initialProgram, session.SessionID, session.WinStationName, session.State)
		name, err := server.QuerySessionUserName(uint(session.SessionID))
		if err != nil {
			panic(err)
		}

		fmt.Printf(",UserName: %s\n", name)
		//if strings.HasPrefix(session.WinStationName, "Console") {
		//	//if session.SessionID == 2 {
		//	appName, err := server.QuerySessionApplicationName(uint(session.SessionID))
		//	if err != nil {
		//		panic(err)
		//	}
		//	fmt.Printf(",AppName: %s\n", appName)
		//	s, err := server.QuerySessionProcessEx(uint(session.SessionID))
		//	if err != nil {
		//		panic(err)
		//	}
		//	var sum int64 = 0
		//	for _, v := range s {
		//		sum += v.Memory
		//
		//		if strings.HasPrefix(v.ProcessName, "navicat.exe") ||
		//			strings.HasPrefix(v.ProcessName, "clion") ||
		//			strings.HasPrefix(v.ProcessName, "WeChat") {
		//			fmt.Printf("%s Memory==> %s    PeakMemory==> %s    PagefileUsage==> %s    PeakPagefileUsage==> %s   HandleCount==> %d   NumberOfThreads==> %d ProcessId==> %d \n", v.ProcessName,
		//				gone.FormatBytesStringOhMyGod(v.Memory),
		//				gone.FormatBytesStringOhMyGod(v.PeakMemory),
		//				gone.FormatBytesStringOhMyGod(v.PagefileUsage),
		//				gone.FormatBytesStringOhMyGod(v.PeakPagefileUsage),
		//				v.HandleCount,
		//				v.NumberOfThreads,
		//				v.ProcessId,
		//			)
		//
		//		}
		//		//sv, _ := json.MarshalIndent(s, "", " ")
		//		//fmt.Println(string(sv))
		//	}
		//	fmt.Printf("Sumary %s\n", gone.FormatBytesStringOhMyGod(sum))
		//
		//}
	}
}
