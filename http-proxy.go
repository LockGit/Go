/**
 * Created by Gogland.
 * User: lock
 * Date: 2017/10/29
 * Time: 00:31
 */
package main

import (
	"log"
	"flag"
	"net"
	"fmt"
	"bytes"
	"net/url"
	"strings"
	"io"
)

//set log format
func init()  {
	log.SetFlags(log.Ldate|log.Ltime|log.Lshortfile)
}

//handle network request
func handleRequest(client net.Conn,debug string)  {
	if client == nil {
		return
	}
	//the end execute client close
	defer client.Close()

	var bin [2048]byte
	n, err := client.Read(bin[:])
	if err != nil {
		log.Println(err)
		return
	}
	var method, host, address string
	if debug == "true" {
		log.Println("start get request stream data ...")
		log.Println(string(bin[:]))
	}
	fmt.Sscanf(string(bin[:bytes.IndexByte(bin[:], '\n')]), "%s%s", &method, &host)
	hostPortURL, err := url.Parse(host)
	if err != nil {
		log.Println(err)
		return
	}

	if hostPortURL.Opaque == "443" { //https访问
		address = hostPortURL.Scheme + ":443"
	} else { //http访问
		if strings.Index(hostPortURL.Host, ":") == -1 { //host不带端口， 默认80
			address = hostPortURL.Host + ":80"
		} else {
			address = hostPortURL.Host
		}
	}

	//向真正目标服务器请求
	server, err := net.Dial("tcp", address)
	if err != nil {
		log.Println(err)
		return
	}
	if method == "CONNECT" {
		fmt.Fprint(client, "HTTP/1.1 200 Connection established\r\n\r\n")
	} else {
		server.Write(bin[:n])
	}
	//转发
	go io.Copy(server, client)
	io.Copy(client, server)
}


func main()  {
	ip := flag.String("ip", "127.0.0.1", "--set listen ip address")
	port := flag.String( "port", "8080", "--set http listen port")
	debug := flag.String("debug","false","--set debug mode true or false")
	flag.Parse()

	log.Printf("start listen %s:%s", *ip, *port)
	listen, err := net.Listen("tcp", *ip + ":" + *port)
	if err != nil {
		log.Panic(err)
	}
	for {
		client, err := listen.Accept()
		if err != nil {
			log.Panic(err)
		}
		go handleRequest(client,*debug)
	}
}

