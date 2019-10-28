/**
 * Created by lock
 * Date: 2019-10-28
 * Time: 11:30
 */
package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
)

var ip = "0.0.0.0"
var port = 16888
var s5port = 26888

type ProxyServer struct {
	Ip     string
	Port   int
	S5Port int
}

func (ps *ProxyServer) handler() {
	ip := ps.Ip
	port := ps.Port
	s5port := ps.S5Port
	listen, err := net.ListenTCP("tcp", &net.TCPAddr{net.ParseIP(ip), port, ""})
	if err != nil {
		panic(err)
	}
	defer listen.Close()
	fmt.Println("listen port:", port)
	s5listen, err := net.ListenTCP("tcp", &net.TCPAddr{net.ParseIP(ip), s5port, ""})
	if err != nil {
		panic(err)
	}
	defer s5listen.Close()
	fmt.Println("socks5 port:", s5port)
	ps.server(listen, s5listen)
}

func (ps *ProxyServer) server(listen *net.TCPListener, s5listen *net.TCPListener) {
	for {
		s5conn, err := s5listen.Accept()
		if err != nil {
			fmt.Println("accept client exception:", err.Error())
			continue
		}
		fmt.Println("client connect from:", s5conn.RemoteAddr().String())
		//defer s5conn.Close()

		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept client connect exception:", err.Error())
			continue
		}
		fmt.Println("[^_^]client from:", conn.RemoteAddr().String())
		go ps.handleConn(conn, s5conn)
	}
}

func (ps *ProxyServer) handleConn(srcConn net.Conn, dstConn net.Conn) {
	defer srcConn.Close()
	defer dstConn.Close()
	exitChan := make(chan bool, 1)
	go func(src net.Conn, dst net.Conn, Exit chan bool) {
		io.Copy(dst, src)
		exitChan <- true
	}(srcConn, dstConn, exitChan)

	go func(src net.Conn, dst net.Conn, Exit chan bool) {
		io.Copy(src, dst)
		exitChan <- true
	}(srcConn, dstConn, exitChan)
	<-exitChan
}

func main() {
	args := os.Args
	port, _ = strconv.Atoi(args[1])
	s5port, _ = strconv.Atoi(args[2])
	ps := &ProxyServer{
		Ip:     ip,
		Port:   port,
		S5Port: s5port,
	}
	ps.handler()
}
