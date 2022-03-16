package main

import (
	"encoding/json"
	"fmt"
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
		if strings.HasPrefix(session.WinStationName, "RDP-") {
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

			sv, _ := json.MarshalIndent(s, "", " ")
			fmt.Println(string(sv))
		}

	}
}
