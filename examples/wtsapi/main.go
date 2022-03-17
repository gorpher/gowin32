package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gorpher/gowin32"
	"time"
)

func main() {
	server := gowin32.OpenWTSServer("")
	sessions, err := server.EnumerateSessions()
	if err != nil {
		panic(err)
	}
	fmt.Println(len(sessions))

	w := make(chan []byte, 2)
	r := make(chan []byte, 2)
	cname := flag.String("n", "c1", "")
	flag.Parse()
	ctx, cancelFunc := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancelFunc()
	query, err := server.WTSVirtualChannelQuery(*cname, 1)
	if err != nil {
		panic(err)
	}
	fmt.Println("get virtual channenl info", query)
	go func() {
		i := 0
		for {
			time.Sleep(1 * time.Second)
			i += 1
			w <- []byte(fmt.Sprintf("hello world :%d", i))
		}
	}()

	go func() {
		for buf := range r {
			fmt.Println("get data from client：", string(buf))
		}
	}()

	err = server.OpenWTSVirtualChannel(ctx, *cname, w, r)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("done")
}
