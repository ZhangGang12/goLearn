package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

//注册命令行参数
var port = flag.String("port", "8000", "get clock port")
var tz = flag.String("tz", "UTC", "get tz")

func clock2() {
	flag.Parse()
	//从终端中获取端口号
	listener, err := net.Listen("tcp", "localhost:"+*port)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*port + *tz)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		//handleConn(conn) // handle one connection at a time
		go handleConn2(conn) // handle multiple connections at a time
	}
}

func handleConn2(c net.Conn) {
	defer c.Close()
	for {
		//init location
		loc, err := time.LoadLocation(*tz)
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.WriteString(c, time.Now().In(loc).Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}
