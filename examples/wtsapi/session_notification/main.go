package main

import (
	"context"
	"github.com/gorpher/gowin32/session_notifications"
	"log"
)

func main() {

	chanMessages := make(chan session_notifications.Message, 100)
	ctx, cl := context.WithCancel(context.Background())
	defer cl()
	go func() {
		for {
			select {
			case m := <-chanMessages:
				log.Println(m)
				switch m.UMsg {
				case session_notifications.WM_WTSSESSION_CHANGE:
					switch m.WParam {
					case session_notifications.WTS_CONSOLE_CONNECT:
						log.Println("WTS_CONSOLE_CONNECT: 由 lParam 标识的会话已连接到控制台终端或RemoteFX会话。")
					case session_notifications.WTS_CONSOLE_DISCONNECT:
						log.Println("WTS_CONSOLE_DISCONNECT: lParam 标识的会话与控制台终端或RemoteFX会话断开连接。")
					case session_notifications.WTS_REMOTE_CONNECT:
						log.Println("WTS_REMOTE_CONNECT : lParam 标识的会话已连接到远程终端。")
					case session_notifications.WTS_REMOTE_DISCONNECT:
						log.Println("WTS_REMOTE_DISCONNECT : lParam 标识的会话已与远程终端断开连接。")
					case session_notifications.WTS_SESSION_LOGON:
						log.Println("WTS_SESSION_LOGON : 用户已登录到 lParam 标识的会话。")
					case session_notifications.WTS_SESSION_LOGOFF:
						log.Println("WTS_SESSION_LOGOFF : 用户已注销 lParam 标识的会话。")
					case session_notifications.WTS_SESSION_LOCK:
						log.Println("session ID", m.LParam, "WTS_SESSION_LOCK : lParam 标识的会话已被锁定。")
					case session_notifications.WTS_SESSION_UNLOCK:
						log.Println("session ID", m.LParam, "WTS_SESSION_UNLOCK  : lParam 标识的会话已解锁。")
					case session_notifications.WTS_SESSION_REMOTE_CONTROL:
						log.Println("session ID", m.LParam, "WTS_SESSION_REMOTE_CONTROL  : 由 lParam 标识的会话已更改其远程控制状态。 若要确定状态，请调用 GetSystemMetrics 并检查 SM_REMOTECONTROL 指标。")
					}
				case session_notifications.WM_QUERYENDSESSION:
					log.Println("log off or shutdown")
				}
				close(m.ChanOk)
			}
		}
	}()

	session_notifications.Subscribe(chanMessages, ctx)

	// ctrl+c to quit
	<-ctx.Done()
}
