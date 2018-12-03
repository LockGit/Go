/**
 * Created by Vim
 * User: lock
 * Date: 2018/12/3
 * Time: 14:59
 */
package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func startPrintln(content string) {
	fmt.Println("\033[1;31m" + content + "\033[0m")
}

func hightRecord(content string) {
	fmt.Println("\033[1;32m" + content + "\033[0m")
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("please input the domain !")
		fmt.Println("example: fdns baidu.com")
		os.Exit(0)
	}

	domain := args[0]
	ipRecords, _ := net.LookupIP(domain)
	startPrintln("Find A OR AAAA Record:")
	for _, ip := range ipRecords {
		hightRecord(ip.String())
	}
	startPrintln("----------")
	startPrintln("Cname:")
	cname, _ := net.LookupCNAME(domain)
	hightRecord(cname)
	startPrintln("----------")
	startPrintln("Ns Record:")
	nameServer, _ := net.LookupNS(domain)
	for _, ns := range nameServer {
		hightRecord(ns.Host)
	}
	startPrintln("----------")
	startPrintln("Mx Record:")
	mxRecords, _ := net.LookupMX(domain)
	for _, mx := range mxRecords {
		hightRecord(mx.Host + fmt.Sprintf("%v", mx.Pref))

	}
	startPrintln("----------")
	startPrintln("Txt Record:")
	txtRecords, _ := net.LookupTXT(domain)
	for _, txt := range txtRecords {
		hightRecord(txt)
	}
	startPrintln("----------")
	startPrintln("Srv Record:")
	cname, sRvs, err := net.LookupSRV("xmpp-server", "tcp", domain)
	if err == nil {
		for _, srv := range sRvs {
			s := fmt.Sprintf("%v:%v:%d:%d\n", srv.Target, srv.Port, srv.Priority, srv.Weight)
			hightRecord(strings.TrimSuffix(s,"\n"))
		}
	}
}
