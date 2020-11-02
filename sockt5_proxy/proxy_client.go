/**
 * Created by lock
 * Date: 2019-10-28
 * Time: 11:30
 */
package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"os"
)

type ProxyClient struct {
	A, B, C, D   byte
	Port         uint16
	ListenTarget string
}

func (pc *ProxyClient) toAddr() string {
	return fmt.Sprintf("%d.%d.%d.%d:%d", pc.A, pc.B, pc.C, pc.D, pc.Port)
}

func (pc *ProxyClient) handleClientRequest(client net.Conn) {
	if client == nil {
		return
	}
	defer client.Close()
	var b [1024]byte
	n, err := client.Read(b[:])
	if err != nil {
		return
	}
	//deal socks5 protocol
	if b[0] == 0x05 {
		client.Write([]byte{0x05, 0x00})
		n, err = client.Read(b[:])
		var addr string
		switch b[3] {
		case 0x01:
			if err := binary.Read(bytes.NewReader(b[4:n]), binary.BigEndian, &pc); err != nil {
				fmt.Println("request parse error")
				return
			}
			addr = pc.toAddr()
		case 0x03:
			host := string(b[5 : n-2])
			var port uint16
			err = binary.Read(bytes.NewReader(b[n-2:n]), binary.BigEndian, &port)
			if err != nil {
				fmt.Println(err)
				return
			}
			addr = fmt.Sprintf("%s:%d", host, port)
		}

		server, err := net.Dial("tcp", addr)
		if err != nil {
			return
		}
		defer server.Close()
		//response client success
		client.Write([]byte{0x05, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})
		//stream forward
		go io.Copy(server, client)
		io.Copy(client, server)
	}
}

func (pc *ProxyClient) handle() string {
	var RemoteConn net.Conn
	var err error
	fmt.Println(fmt.Sprintf("start run at --> %s", pc.ListenTarget))
	for {
		for {
			RemoteConn, err = net.Dial("tcp", pc.ListenTarget)
			if err == nil {
				break
			}
		}
		go pc.handleClientRequest(RemoteConn)
	}
}

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("args error! please input ip:port")
		os.Exit(0)
	}
	pc := &ProxyClient{
		ListenTarget: args[1],
	}
	pc.handle()
}
