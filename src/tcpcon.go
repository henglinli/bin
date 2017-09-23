// -*-coding:utf-8-unix;-*-
package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func conn(remote string) {
	start := time.Now()
	conn, err := net.DialTimeout("tcp", remote, 2*time.Second)
	if err != nil {
		fmt.Println(err)
		return
	}
	elapsed := time.Since(start)
	fmt.Println(remote, elapsed)
	defer conn.Close()
}

func main() {
	args := len(os.Args)
	var addr, port string
	switch args {
	case 2:
		addr = os.Args[1]
		port = "443"
	case 3:
		addr = os.Args[1]
		port = os.Args[2]
	default:
		fmt.Println("Usage: ", os.Args[0], "ip port")
		return
	}
	remote := addr + ":" + port
	//
	conn(remote)
}
